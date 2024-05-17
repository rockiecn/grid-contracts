// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/access/Ownable.sol";

/**
 * @author MemoLabs
 * @title manage erc20
 * @dev Gas: 2366834
 */
contract Access is Ownable {
    uint16 public version = 1;
    
    uint256 public nonce;
    
    // access management
    mapping(address => bool) public access;

    /// @dev Deployed by admin. Gas:545446
    constructor() 
        Ownable() 
    {}

    /// @dev Gas: 112394
    function set(address addr, bool isSet) external onlyOwner {
        // set access for this addr
        access[addr] = isSet;

        // increase nonce
        nonce+=1;
    }
}