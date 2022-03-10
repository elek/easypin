// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Faucet is Ownable {

    IERC20 token;

    constructor(address _tokenAddress) Ownable(){
        token = IERC20(_tokenAddress);
    }

    function get() public {
        require(token.transfer(msg.sender, 1000000000000000000));
    }

    function withdraw(address target) public onlyOwner {
        uint256 balance = token.balanceOf(address(this));
        token.transfer(target, balance);
    }
}
