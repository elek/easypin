// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Gold is ERC20 {
    constructor() ERC20("Gold", "GLD")  {
        _mint(msg.sender, 1000000000000000000000000);
    }
}
