// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// 
interface IRegistry {
    function update(address provider, uint64 nCPU, uint64 nGPU, uint64 nMEM, uint64 nDISK) external;
}