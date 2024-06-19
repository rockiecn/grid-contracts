// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IAccess.sol";


contract Credit is ERC20, Ownable {
    // access contract address to manage mint right
    address public access;

    //uint256 private _totalSupply; // 上限6亿；初始发行3亿
    uint256 public constant initialSupply = 3 * 10**26;
    uint256 public constant maxSupply = 6 * 10**26;

    // need access contract address
    constructor(address _a)
        ERC20("Credit", "CRE") //此处表示代币的名称和缩写
        Ownable()
    {
        access = _a;
        _mint(msg.sender, initialSupply);
    }


    // credit can be mint by admin without limit
    function mint(address to, uint256 amount) external {
        // check access right for sender
        require(IAccess(access).access(msg.sender), "caller no mint"); 

        _mint(to, amount);
    }

    // everyone can burn its value; need it or add it in transfer
    // gas: 38435
    function burn(uint256 burnAmount) external {
        _burn(msg.sender, burnAmount);
    }
}