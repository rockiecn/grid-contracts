// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "../openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../openzeppelin/contracts/access/Ownable.sol";

/**
 * @title Pledge
 * @dev For pledge management
 */
contract Pledge is Ownable {
    // record the price per hour for each resource type
    struct PledgeInfo{
        uint256 pledge;
        uint256 interest;

        // the timestamp for the last settlement
        uint256 lastTime;
    }

    // the interest rate
    uint256 rate;

    // the pledge info for everybody
    mapping(address => PledgeInfo) pledges;

    // constructor, the creator as the contract owner
    constructor() 
        Ownable()
    {
        // the approximate interest/sec for 1 ETH at the yearly rate of 2%
        rate = 6 * 10e8; 
    }

    // sender pledge some token, approve to contract must be called first
    function pledge(address tokenAddr, uint256 amount) external {
        require(amount > 0,"amount of pledge should larger than 0");

        // settle the interest
        _settle(msg.sender);

        // transfer from sender
        IERC20(tokenAddr).transferFrom(msg.sender, address(this), amount);

        // update pledge amount
        pledges[msg.sender].pledge += amount;
    }

    // unpledge some token
    function unpledge(address tokenAddr, uint256 amount) external {
        require(amount > 0,"amount of pledge should larger than 0");
        require(pledges[msg.sender].pledge >= amount, "the remained pledge not enough for unpledge");

        // settle the interest
        _settle(msg.sender);

        // transfer to sender
        IERC20(tokenAddr).transfer(msg.sender, amount);

        // update pledge amount
        pledges[msg.sender].pledge -= amount;
    }

    // get the current pledge amount
    function getPledge() view external returns(uint256) {
        return pledges[msg.sender].pledge;
    }

    // settle the current pending interest for an address
    function _settle(address a) internal {
        uint256 nowTime = block.timestamp;
        uint256 lastTime = pledges[a].lastTime;
        require(nowTime > lastTime, "now time not later than last time");

        // the seconds to calculate interest
        uint256 since = nowTime - lastTime;
        // new profit to add
        uint256 profit = since * rate;

        // update the interest for the address
        pledges[a].interest += profit;
        // update last time
        pledges[a].lastTime = nowTime;
    }

    // withdraw some interest
    function withdrawInterest(address tokenAddr, uint256 amount) external {
        // settle the pending interest before hand
        _settle(msg.sender);

        require(pledges[msg.sender].interest >= amount, "not enough interest remained");

        // transfer to sender
        IERC20(tokenAddr).transfer(msg.sender, amount);

        // update interest amount
        pledges[msg.sender].interest -= amount;
    }

    // set the interest rate, only called by owner
    function setRate(uint256 _rate) external onlyOwner() {
        // the rate should between (0 ~ 20%)
        require(_rate > 0 && _rate < 6 * 10e10);

        rate = _rate;
    }
    
    // get the current interest rate
    function getRate() view external returns(uint256) {
        return rate;
    }

    // update the current interest amount with settle
    function updateInterest() external {
        _settle(msg.sender);
    }

    // get interest
    function getInterest() view external returns(uint256){
        return pledges[msg.sender].interest;
    }
}