# generate-block-for-testing

This application is used to perform block generation in a testing network for the bitcoin or the liquid network.

## function

- generatetoaddress
  - use address
  - (in the future) generate address by bip32
  - (in the future) show the generated block's hash.
- generate block for dynamic federation (liquid network)
  - set fedpeg script
  - set pak entry
  - (in the future) block sign
  - (in the future) maximum witness size
  - (in the future) show the generated block's hash.

## usage

### command-line argument

```sh
Usage: generateblock.exe [--host HOST] [--fedpegscript FEDPEGSCRIPT] [--pak PAK] [--network NETWORK] [--address ADDRESS] [--rpcuserid RPCUSERID] [--rpcpassword RPCPASSWORD] [--logging]

Options:
  --host HOST            connection host & port
  --fedpegscript FEDPEGSCRIPT, -s FEDPEGSCRIPT
                         fedpeg script on dynafed
  --pak PAK, -p PAK      pak entries
  --network NETWORK, -n NETWORK
                         network. (bitcoin:mainnet/testnet/regtest, liquid:liquidv1/liquidregtest/elementsregtest)
  --address ADDRESS, -a ADDRESS
                         bitcoin address for generatetoaddress
  --rpcuserid RPCUSERID
                         connection rpc userID
  --rpcpassword RPCPASSWORD
                         connection rpc password
  --logging, -l          log output
  --help, -h             display this help and exit
```

### environment variable

- GENERATE_BLOCK_CONNECTION_HOST: host & port.
- GENERATE_BLOCK_CONNECTION_NETWORK: network type.
  - bitcoin: mainnet, testnet, regtest
  - liquid network: liquidv1, liquidregtest, (elementsregtest)
- CONNECTION_PRC_USERID: connection rpc userID
- CONNECTION_PRC_PASSWORD: connection rpc password
- (for generatetoaddress)
  - GENERATE_BLOCK_GENERATETOADDRESS: bitcoin address
- (liquid network parameter)
  - DYNAFED_FEDPEG_SCRIPT: fedpeg script.
  - DYNAFED_PAK: pak entry. To set multiple items, separate them with commas.

If both environment variable and command-line argument are set, the value set for command-line argument will take precedence.

## build

```go
go build ./cmd/generateblock/
```

```sh
make build
```

### format

```sh
make gettools format
```
