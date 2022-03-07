// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "Ownable.sol";
import "ERC20.sol";

contract StorjPin is Ownable {

    event Pinned(address indexed owner, uint256 amount, string hash);

    ERC20 token;

    constructor(address _tokenAddress) Ownable(){
        token = ERC20(_tokenAddress);
    }

    function pin(string memory ipfsHash, uint256 tokenAmount) public {
        require(token.transferFrom(msg.sender, address(this), tokenAmount));
        emit Pinned(msg.sender, tokenAmount, ipfsHash);
    }

    function withdraw(address target) public onlyOwner {
        uint256 balance = token.balanceOf(address(this));
        token.transfer(target, balance);
    }

}
