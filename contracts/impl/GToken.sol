// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../openzeppelin/contracts/access/Ownable.sol";


contract GToken is ERC20, Ownable {
    constructor()
        ERC20("GToken", "GTK") //此处表示代币的名称和缩写
        Ownable()
    {}

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }
}