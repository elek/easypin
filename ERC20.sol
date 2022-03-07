// SPDX-License-Identifier: UNLICENSED

interface ERC20 {

    function balanceOf(address who) external view returns (uint);

    function allowance(address owner, address spender) external view returns (uint);

    function transfer(address to, uint value) external returns (bool ok);

    function transferFrom(address from, address to, uint value) external returns (bool ok);

    function approve(address spender, uint value) external returns (bool ok);

    event Transfer(address indexed from, address indexed to, uint value);
    event Approval(address indexed owner, address indexed spender, uint value);
}

