// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../openzeppelin/contracts/access/Ownable.sol";


contract MyToken is ERC20, Ownable{

    // initial tokens for the owner
    uint256 public constant initialSupply = 3 * 10**26;
    
    // record all nodes' pledge info, assign profits for all nodes when a profit is received
    struct PledgePool{
        // record all nodes' pledge amount
        mapping(address => uint256) _pledges;

        // pool total pledge amount in pool, updated when node pledge or withdraw pledge
        uint256 poolPledge;

        // the current reward in pool, it shuold be given to all nodes with pledges, and this is done by settle when pledge or unpledge happens 
        // after each settle operation, this value must be reset
        uint256 reward;

        // the accumulated reward in pool
        uint256 historyReward;

        // the accumulated profit factor for the pool, it should be updated when a profit is received in pool
        uint256 accu;

        // the latest pledge time
        uint256 pledgeTime;
    }

    // record the reward info for each account
    struct RewardInfo{
        // current reward that can be withdrawed
        uint256 reward;

        // accumulative pledge amount
        uint256 totalPledge;
        // total reward, accumulated for all profits
        uint256 totalReward;

        // the reward factor for an account, it means the part of reward that has been withdrawed
        uint256 accu;
    }

    // node address to reward info
    mapping(address => RewardInfo) public rewards;

    // pledge pool data
    PledgePool public pp;

    // contract constructor, can assign an address to be the initial owner
    constructor(address )
        ERC20("GridToken", "GTK")//此处表示代币的名称和缩写
        Ownable()
    {
        // mint some tokens for the initial owner
        _mint(msg.sender, initialSupply);
    }

    // mint for token
    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    // additional reward for pledge pool, called by owner when new order created
    function addition(uint256 amount) public onlyOwner {
        require(amount>=0, "addition amount less than 0");

        // update pledge pool reward info
        pp.reward += amount;
        pp.historyReward += amount;
    }

    // pledge for a node
    function pledge(address addr, uint256 amount) public {
        // settle current reward in pool
        updateReward(addr);

        // update node pledge
        pp._pledges[addr] += amount;
        // update pool pledge
        pp.poolPledge += amount; 
    }

    // unpledged for a node
    function unPledge(address addr, uint256 amount) public {
        // settle current reward in pool
        updateReward(addr);

        // update node pledge
        require(pp._pledges[addr] >= amount, "too many amount to unpledge");
        pp._pledges[addr] -= amount;
        // update pool pledge
        require(pp.poolPledge >= amount, "too many amount to unpledge");
        pp.poolPledge -= amount; 
    }

    // query the pledge amount for a node
    // reward update is needed
    function getPledge(address addr) public view returns (uint256) {
        // need settle reward first
        return pp._pledges[addr];
    }

    // get total pledge amount of a node
    function getTotalPledge(address addr) public view returns (uint256) {
        return rewards[addr].totalPledge;
    }

    // get the current reward amount of a node
    function getReward(address addr) public view returns (uint256) {
        return rewards[addr].reward;
    }

    // get total reward amount of a node
    function getTotalReward(address addr) public view returns (uint256) {
        return rewards[addr].totalReward;
    }

    // get total pledge in pool
    function getPoolPledge() public view returns (uint256) {
            return pp.poolPledge;
        }

    // get current reward in pool
    function getPoolReward() public view returns (uint256) {
        return pp.reward;
    }
    
    // get historical reward in pool
    function getPoolHistoryReward() public view returns (uint256) {
        return pp.historyReward;
    }

    //?
    // node withdraw some reward into token balance, mint reward amount into balance
    function withdrawReward(address addr, uint256 amount) public {
        require(rewards[addr].reward >= amount, "not enough reward");
        rewards[addr].reward -= amount;

        // mint token for node
        _mint(addr, amount);
    }

    // update the reward for a node
    function updateReward(address addr) public {
        // get current pledge amount for node
        uint256 amount = pp._pledges[addr];

        // get the current reward amount of pool
        uint256 poolReward = pp.reward;
        // get the total plege amount in pool
        uint256 poolPledge = pp.poolPledge;
        // calculate the new uint reward for all pledge
        uint256 u = poolReward * 1e18 / poolPledge;

        // update the accu of pool
        pp.accu += u;
        // increase the total reward for pool
        pp.historyReward += poolReward;
        // after accu updated, the current pool reward must be reset
        pp.reward = 0;

        // calc the reward amount for the node
        uint256 delta = (pp.accu - rewards[addr].accu);
        uint256 re = delta * amount / 1e18;

        // update reward info for node
        rewards[addr].reward += re;
        rewards[addr].totalReward += re;
        rewards[addr].accu = pp.accu;
    }
}
