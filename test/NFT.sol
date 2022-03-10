// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

import "IStorjPin.sol";

contract NFT is ERC721URIStorage, AccessControl {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    IStorjPin pin;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    constructor(address pinContract) ERC721("Stamps", "STM") {
        _setupRole(MINTER_ROLE, _msgSender());
        pin = IStorjPin(pinContract);
    }

    function awardItem(address player, string memory tokenURI) public returns (uint256)
    {

        require(hasRole(MINTER_ROLE, msg.sender), "Caller is not a minter");
        _tokenIds.increment();

        pin.pin(tokenURI, 10000000);
        uint256 newItemId = _tokenIds.current();
        _safeMint(player, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }

    function _baseURI() internal view virtual override returns (string memory) {
        return "https://gateway.ipfs.io/ipfs/";
    }


    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}