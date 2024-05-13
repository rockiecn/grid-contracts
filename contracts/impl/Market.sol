// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "../openzeppelin/contracts/token/ERC20/IERC20.sol";

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

    // constructor
    constructor() {
    }

    /**
     * @dev user create an order with a provider， user approve must be called previously
     */
    function createOrder(address tokenAddr, address provider, Order memory order) external {
        // transfer token to market contract
        IERC20(tokenAddr).transferFrom(msg.sender, address(this), order.totalValue);

        // store order
        orders[msg.sender][provider] = order;
    }

    /**
     * @dev get the order info with a provider
     */
    function getOrder(address provider) external view returns(Order memory) {
        return orders[msg.sender][provider];
    }

    // calc the fee of an order
    function feeOrder(Order memory order) external view returns(uint256) {

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
        orders[user][msg.sender].lastSettleTime=block.timestamp;
        orders[user][msg.sender].activateTime=block.timestamp;
        orders[user][msg.sender].status=1;
    }

    // the user cancel the order
    // refund credit in this order is needed
    function userCancel(address tokenAddr, address provider) external {
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
        IERC20(tokenAddr).transfer(msg.sender, _order.remain);
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


    // settle fee for an order, up to the settleTime
    function _settle(address user, address provider) internal {
        Order memory _order = orders[user][provider];

        // current timestamp
        //uint256 nowTime = block.timestamp;
        uint256 nowTime = _order.lastSettleTime+10;// for test

        require(_order.status==1,"order must be in activate status to settle");
        require(nowTime >= _order.lastSettleTime);

        // calc the elapsed time since last settlement.
        // note: the last time is set to activate time when order activated.
        // use second this version.
        uint256 elapsed = nowTime - _order.lastSettleTime;
        // the total service time
        uint256 total = nowTime - _order.activateTime;

        // unit fee per second for an order
        uint256 unitFee = _order.totalValue/_order.duration;
        // the service fee for this time
        uint256 settleFee;

        // check if order is time up
        bool timeup = false;
        if (total >= _order.duration) {
            settleFee = _order.remain;
            timeup = true;
        } else {
            settleFee = unitFee * elapsed;
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
    
    // provider withdraw some token in an order
    function proWithdraw(address tokenAddr, address user, uint256 amount) external {
        Order memory _order = orders[user][msg.sender];

        require(_order.provider == msg.sender,"caller must be the provider");
        require(amount <= _order.remuneration,"the amount should not larger than remuneration");

        // transfer token to the provider
        IERC20(tokenAddr).transfer(msg.sender, amount);
        orders[user][msg.sender].remuneration -= amount;
    }
}