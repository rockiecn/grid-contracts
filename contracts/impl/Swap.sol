// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/access/Ownable.sol";
import "../openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/ICredit.sol";


contract Swap is Ownable {
    using Math for uint256;

    address public creditAddr; // credit token 
    address public gtokenAddr;  // gtoken

    // tokenAmount = credit * rate_numerator / rate_denominator
    uint256 public immutable rate_numerator; 
    uint256 public immutable rate_denominator;

    event Bought(address indexed buyer, uint256 amount);
    event Sold(address indexed seller, uint256 amount);
    event Received(address sender, uint256 amount);

    // token addresses and rate
    constructor(address _credit, address _gtoken, uint256 _rate_numerator, uint256 _rate_denominator) {
        creditAddr = _credit;
        gtokenAddr = _gtoken;

        // rate
        rate_numerator = _rate_numerator;
        rate_denominator = _rate_denominator;
    }

    // use credit to buy token
    function buy(uint256 _creditAmount) payable public {
        require(_creditAmount > 0, "credit amount must larger than 0");

        // check allowance, approve credit to swap must be called first
        uint256 allowance = IERC20(creditAddr).allowance(msg.sender, address(this));
        require(allowance >= _creditAmount);
        // transfer credit to swap
        bool ok = IERC20(creditAddr).transferFrom(msg.sender, address(this), _creditAmount);
        require(ok, "TransferFrom fail");

        // check gtoken balance in swap contract
        uint256 tokenBal = IERC20(gtokenAddr).balanceOf(address(this));
        uint256 tokenAmount = _creditAmount.mulDiv(rate_numerator, rate_denominator);
        require(tokenAmount <= tokenBal, "not enough token in swap contract to buy");
        // swap transfer gtoken to sender
        ok = IERC20(gtokenAddr).transfer(msg.sender, tokenAmount);
        require(ok, "Transfer fail");

        emit Bought(msg.sender, _creditAmount);
    }

    // sell token to get credit
    // caller should approve 'address(this)' of _tokenAmount first
    function sell(uint256 _tokenAmount) public {
        require(_tokenAmount>0, "You need to sell at least some tokens");

        // approve to swap in token must be called first
        uint256 allowance = IERC20(gtokenAddr).allowance(msg.sender, address(this));
        require(allowance >= _tokenAmount, "Check the token allowance");
        // transfer token to swap
        bool ok = IERC20(gtokenAddr).transferFrom(msg.sender, address(this), _tokenAmount);
        require(ok, "transferFrom failed");

        // check credit balance in swap
        uint256 creditAmount = _tokenAmount.mulDiv(rate_denominator, rate_numerator);

        // mint to user directly
        ICredit(creditAddr).mint(msg.sender,creditAmount);
       
        emit Sold(msg.sender, _tokenAmount);
    }

    // send remaining token to owner
    function settlementToken() public onlyOwner {
        uint256 tokenBal = IERC20(gtokenAddr).balanceOf(address(this));
        if(tokenBal>0){
            require(IERC20(gtokenAddr).transfer(owner(), tokenBal), "Transfer fail");
        }
    }
}