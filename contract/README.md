## Deploy GOLD token (key1)

```
cethacea contract deploy --account key1 --name GOLD --abi test/Gold.abi test/Gold.bin
cethacea token balance --account key1 --contract GOLD
```

## Deploy FAUCET (key1)

```
cethacea contract deploy --account key1 --name FAUCET --abi test/Faucet.abi test/Faucet.bin '(address)' GOLD 
cethacea token transfer --contract GOLD --account key1 100000000000000000000 FAUCET 
cethacea token balance --account key1 --contract GOLD FAUCET
cethacea token balance --account key1 --contract GOLD key1

```

## Self-request token (key2)

```
cethacea contract call --account key2 --contract FAUCET 'get()'
cethacea token balance --contract GOLD key2
```

## Deploy the pin contract (key1)

```
cethacea contract deploy --account key1 --name PIN --abi StorjPin.abi StorjPin.bin '(address)' GOLD
```

## Use the ping contract (key2)

```
cethacea contract call --account key2 --contract GOLD approve PIN 10000000000000000000000 
cethacea contract call --account key2 --contract PIN pin 'QmcUkP3BMDkKNp2V6FTojXqgCMWdpDFQb74uphQVdWpi9Z' 100000000000000000
cethacea log --contract PIN
```

## Starting the API server:

```
go run ./cmd/easypin/main.go --pin.endpoint=http://sd:8545 \
   --pin.token-address D3154BE2863cc6FCA8dE0F0D37a82c0cce15Fc35 \
   --database 'postgres://root@sd:26257/pin?sslmode=disable' \
   --api.address 127.0.0.1:8787
```

```
curl localhost:8787/api/v0/pin/all
[{"Cid":"QmcUkP3BMDkKNp2V6FTojXqgCMWdpDFQb74uphQVdWpi9Z","TokenValue":100000000000000000,"Transaction":"0xace7f0e4ad0a2a9e93e14db6e870b4f6dbcbea88b2ddf016f37d6b63909887ad"}]
```