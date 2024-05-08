// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "../openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title Pledge
 * @dev For pledge management
 */
contract Pledge {
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
}