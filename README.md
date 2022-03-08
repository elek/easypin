

# Rinbkeby deployment

ERC20 token (GOLD): [0x39AaBAF5DebD0855908D382C81eAD5705C52Fc5a](https://rinkeby.etherscan.io/address/0x39AaBAF5DebD0855908D382C81eAD5705C52Fc5a)

FAUCET contract (call the `get` to get 1 TOKEN): [0x0d5BEf7d1FeAEaF3E28CD1d134ab17a49B9bBF38](https://rinkeby.etherscan.io/address/0x39AaBAF5DebD0855908D382C81eAD5705C52Fc5a)

PIN contract: [0x245a6405aA8377D3cA06160E24D4eb6622a02C66](https://rinkeby.etherscan.io/address/0x245a6405aA8377D3cA06160E24D4eb6622a02C66)


Deployed with:

```
export CETH_ACCOUNT=...
ceth c deploy --name GOLD --abi ../test/Gold.abi ../test/Gold.bin
ceth c deploy --name FAUCET --abi ../test/Faucet.abi ../test/Faucet.bin '(address)' GOLD
ceth c deploy --name PIN --abi ../contract/ ../test/Faucet.bin '(address)' GOLD



```

# Local deployment

## Generate two accounts

```
cethacea account generate
cethacea account generate
```

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
go run ./cmd/easypin/main.go \
   --pin.endpoint=https://mainnet.infura.io/v3/PROJECT_ID \
   --pin.token-address 0x245a6405aA8377D3cA06160E24D4eb6622a02C66 \
   --database 'postgres://root@sd:26257/pin?sslmode=disable' \
   --api.address 127.0.0.1:8787
   --ipfs.address /ip4/127.0.0.1/tcp/5001
```

Requirements:

 1. Working ethereum node where contracts are deployed (you can use infura rinkeby URL or a local chain from `./testchain`)
 2. Cockroach (or postgres). Please create the scheme (`pin` in our example)
 3. Running IPFS node (`ipfs daemon`) 

```
curl localhost:8787/api/v0/pin/all
[{"Cid":"QmcUkP3BMDkKNp2V6FTojXqgCMWdpDFQb74uphQVdWpi9Z","TokenValue":100000000000000000,"Transaction":"0xace7f0e4ad0a2a9e93e14db6e870b4f6dbcbea88b2ddf016f37d6b63909887ad"}]
```

# Web development

Developer version (proxy is activated for `/api/*` to proxy requests to `127.0.0.1:8787`: it requires running go lang instance ):

```
cd web
npm install
npx vite
```

Production version:

```
cd web
npm build
```

It saves the final artifacts to `web/dist` which is exposed by the golang api.
