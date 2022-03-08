docker-compose exec geth geth attach --exec "eth.sendTransaction({from:eth.coinbase, to:'$1', value: web3.toWei(100, \"ether\")})" /test-chain-dir/geth.ipc
