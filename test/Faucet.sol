// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "ERC20.sol";
import "./Ownable.sol";

contract Faucet is Ownable {

    ERC20 token;

    constructor(address _tokenAddress) Ownable(){
        token = ERC20(_tokenAddress);
    }

    function get() public {
        require(token.transfer(msg.sender, 1000000000000000000));
    }

    function withdraw(address target) public onlyOwner {
        uint256 balance = token.balanceOf(address(this));
        token.transfer(target, balance);
    }
}
