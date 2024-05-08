// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// for manage erc20 
interface IERC20 {
    // mint token
    function mint(address to, uint256 amount) external;
}