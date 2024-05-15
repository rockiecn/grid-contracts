// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../openzeppelin/contracts/access/Ownable.sol";


contract GToken is ERC20, Ownable {
    // 上限6亿；初始发行3亿
    uint256 public constant initialSupply = 3 * 10**26;
    uint256 public constant maxSupply = 6 * 10**26;
   
    constructor()
        ERC20("GToken", "GTK") //此处表示代币的名称和缩写
        Ownable()
    {
        // mint init supply to owner
        _mint(msg.sender, initialSupply);
    }

    function mint(address to, uint256 amount) external onlyOwner {
        uint256 ts = totalSupply();
        if((ts+amount) > maxSupply){
            amount = maxSupply - ts;
        }

        _mint(to, amount);
    }

    // everyone can burn its value; need it or add it in transfer
    // gas: 38435
    function burn(uint256 burnAmount) external {
        _burn(msg.sender, burnAmount);
    }
}