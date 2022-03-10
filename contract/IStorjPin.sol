// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

interface IStorjPin {

    function pin(string memory ipfsHash, uint256 tokenAmount, bool parse) external;

}
