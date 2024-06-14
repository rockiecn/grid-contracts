// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "../interfaces/ICredit.sol";

/**
 * @title Market
 * @dev For order management
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Market {
    // record the price per hour for each resource type
    struct PricePerHour{
        uint64 pCPU;
        uint64 pGPU;
        uint64 pMEM;
        uint64 pDISK;
    }

    // record the number of each resource type
    struct Resources{
        uint64 nCPU;
        uint64 nGPU;
        uint64 nMEM;
        uint64 nDISK;
    }

    // record the order info
    struct Order{
        address user;
        address provider;

        PricePerHour p;
        Resources r;

        // the total value user paid for this order
        uint256 totalValue;

        // the remained value in this order
        uint256 remain;
        // the current profit for provider
        uint256 remuneration;

        // confirmation from user or provider
        bool userConfirm;
        bool providerConfirm;

        // activate timestamp
        uint256 activateTime;
        // the last time settle timestamp
        uint256 lastSettleTime;

        // the time for trying out
        uint256 probation;

        // the time for order's service
        uint256 duration;

        // order status: 0-unactive 1-active 2-cancelled 3-completed
        uint8 status;
    }

    // record all orders for every user/provider pair
    mapping(address => mapping(address => Order)) orders;
    // record all keys for every order
    address[] private keys;

    // constructor
    constructor() {
    }

    event OrderValue(uint256 totalValue);

    /**
     * @dev user create an order with a provider， user approve must be called previously
     */
    function createOrder(address creditAddr, address provider, Order memory order) external {
        // verify the total value of this order with all resources and price
        uint256 totalValue = _valueOrder(order);
        require(totalValue==order.totalValue,"the totalvalue verify failed for this order");

        emit OrderValue(totalValue);

        // transfer token to market contract
        ICredit(creditAddr).transferFrom(msg.sender, address(this), order.totalValue);

        // store order
        orders[msg.sender][provider] = order;
        // record the order key
        keys.push(msg.sender);
    }

    /**
     * @dev get the order info with a provider
     */
    function getOrder(address provider) external view returns(Order memory) {
        return orders[msg.sender][provider];
    }

     /**
     * @dev get the order keys
     */
    function getKeys() external view returns(address[] memory) {
        return keys;
    }

    // calc the totalvalue of an order
    function _valueOrder(Order memory order) internal pure returns(uint256) {
        // unit price * amount
        uint256 nCPU = order.r.nCPU;
        uint256 nGPU = order.r.nGPU;
        uint256 nMEM = order.r.nMEM;
        uint256 nDISK = order.r.nDISK;

        uint256 pCPU = order.p.pCPU;
        uint256 pGPU = order.p.pGPU;
        uint256 pMEM = order.p.pMEM;
        uint256 pDISK = order.p.pDISK;

        uint256 value = (nCPU*pCPU + nGPU*pGPU + nMEM*pMEM + nDISK*pDISK) * order.duration;

        return value;
    }

    // set the status of an order, only called by contract
    function _setStat(address user, address provider, uint8 status) internal {
        require((status==0 || status==1 || status==2 || status==3),"wrong status value");

        orders[user][provider].status = status;
    }

    // activate an order, called by the provider
    function activate(address user) external {
        Order memory _order = orders[user][msg.sender];

        // only be called by the order's provider
        require(msg.sender == _order.provider,"activate only called by provider");
        require(_order.status==0,"only unactive order can be activated");

        // activate
        orders[user][msg.sender].lastSettleTime = block.timestamp;
        orders[user][msg.sender].activateTime = block.timestamp;
        orders[user][msg.sender].status=1;
    }

    // the user cancel the order
    // refund credit in this order is needed
    function userCancel(address creditAddr, address provider) external {
        Order memory _order = orders[msg.sender][provider];

        require(_order.user == msg.sender,"the caller not order's user");
        require(_order.provider == provider,"the provider is error");

        // call settle before cancel
        _settle(msg.sender,provider);

        // get the new order info after settled
        _order = orders[msg.sender][provider];

        // set status to cancelled
        orders[msg.sender][provider].status = 2;

        // transfer remained token to the user
        ICredit(creditAddr).transfer(msg.sender, _order.remain);
        // no token remained in order
        orders[msg.sender][provider].remain = 0;
    }

    // the provider cancel the order
    function providerCancel(address user) external {
        Order memory _order = orders[user][msg.sender];

        require(_order.user == user,"the user is error");
        require(_order.provider == msg.sender,"the caller is not order's provider");

        // call settle before cancel
        _settle(user,msg.sender);

        // 2-cancelled
        orders[user][msg.sender].status = 2;
    }

    // output debug msg
    event Output(uint8 msg);
    event Paytime(uint256 pt);

    // settle fee for an order, up to the settleTime
    function _settle(address user, address provider) internal {
        Order memory _order = orders[user][provider];

        // current timestamp
        uint256 nowTime = block.timestamp;

        require(_order.status == 1,"order must be in activate status to settle");
        require(nowTime >= _order.lastSettleTime,"the settle time must later than last settlement");

        // the time that is not in probation and should be paid
        uint256 payTime;
        // the time when probation is over
        uint256 probationTime =_order.activateTime + _order.probation;

        // last settle in probation
        if (_order.lastSettleTime <= probationTime) {
            // if now time over probation, should pay
            if (nowTime > probationTime) {
                payTime = nowTime - probationTime;
            } else {
                // if now time is in probation, no pay, just change the settle time and return
                orders[user][provider].lastSettleTime = nowTime;
                return;
            }
        } else {
            // if last settle over probation already, just get the pay time
            payTime = nowTime - _order.lastSettleTime;
        }

        emit Paytime(payTime);

        // the total service time
        uint256 total = nowTime - _order.activateTime;

        // unit fee per second for an order
        uint256 unitFee = _order.totalValue/_order.duration;
        // the service fee for this time
        uint256 settleFee;

        // check if order is time up

        bool timeup = false;
        if (total >= _order.duration + _order.probation) {
            // all remaining fee should be paid to provider when order is timeup
            settleFee = _order.remain;
            timeup = true;
        } else {
            // calc settle fee with pay time
            settleFee = unitFee * payTime;
        }


        // update order info

        orders[user][provider].remain -= settleFee;
        orders[user][provider].remuneration += settleFee;
        orders[user][provider].lastSettleTime = nowTime;

        // when service time is over, the order status is set to completed

        if(timeup) {
            orders[user][provider].status = 3;
        }
    }

    // provider settle the fee for an order, called by provider
    function proSettle(address user) external {
        _settle(user, msg.sender);
    }

    // provider withdraw some token in an order
    function proWithdraw(address creditAddr, address user, uint256 amount) external {
        Order memory _order = orders[user][msg.sender];

        require(_order.provider == msg.sender,"caller must be the provider");
        require(amount <= _order.remuneration,"the amount should not larger than remuneration");

        // transfer token to the provider
        ICredit(creditAddr).transfer(msg.sender, amount);
        orders[user][msg.sender].remuneration -= amount;
    }
}