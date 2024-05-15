// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../openzeppelin/contracts/access/Ownable.sol";


contract Credit is ERC20, Ownable {
    constructor()
        ERC20("Credit", "CRE") //此处表示代币的名称和缩写
        Ownable()
    {}


    // credit can be mint by admin without limit
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    // everyone can burn its value; need it or add it in transfer
    // gas: 38435
    function burn(uint256 burnAmount) external {
        _burn(msg.sender, burnAmount);
    }
}