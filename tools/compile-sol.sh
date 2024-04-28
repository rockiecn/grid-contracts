#! /bin/bash

# solc version 0.8.0+commit.c7dfd78e.Darwin.appleclang
# abigen version 1.10.4-stable
# openzeppelin v4.8.0

## 如果使用下面这种：增加'--libraries'选项，之后使用'abigen --abi --bin'的编译方式，则会导致部署合约多约30万gas，调用函数也会多约1万gas
## solc contracts/impl/Access.sol --libraries "contracts/Recover.sol:Recover=0x21818Ceb3e8992F15899582456592307Ec08F1fe" --bin --abi -o ~/Documents/ContractABI/ --overwrite
## solc contracts/impl/Auth.sol --libraries "contracts/Recover.sol:Recover=0x21818Ceb3e8992F15899582456592307Ec08F1fe" --bin --abi -o ~/Documents/ContractABI/ --overwrite

solc ../contracts/GToken.sol --bin --abi -o ../abi/GToken --overwrite
solc ../contracts/Registry.sol --bin --abi -o ../abi/Registry --overwrite

abigen --sol ../contracts/GToken.sol --out ../go/gtoken/gtoken.go --pkg gtoken --bin ../abi/gtoken.bin
abigen --sol ../contracts/Registry.sol --out ../go/registry/registry.go --pkg registry --bin ../abi/registry.bin 