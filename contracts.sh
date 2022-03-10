#!/usr/bin/env bash

solc --base-path `pwd`  --include-path node_modules --overwrite --abi --bin -o contract/ contract/StorjPin.sol
solc --base-path `pwd`  --include-path node_modules --overwrite --abi --bin -o test/ test/Gold.sol
solc --base-path `pwd`  --include-path node_modules --overwrite --abi --bin -o test/ test/Faucet.sol
solc --base-path `pwd` --include-path contract  --include-path node_modules --overwrite --abi --bin -o test/ test/NFT.sol
