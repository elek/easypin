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
