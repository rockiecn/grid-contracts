#! /bin/bash

# solc 0.8.16+commit.07a7930e.Darwin.appleclang
# abigen version 1.10.4-stable
# openzeppelin v4.8.0

## 如果使用下面这种：增加'--libraries'选项，之后使用'abigen --abi --bin'的编译方式，则会导致部署合约多约30万gas，调用函数也会多约1万gas
## solc contracts/impl/Access.sol --libraries "contracts/Recover.sol:Recover=0x21818Ceb3e8992F15899582456592307Ec08F1fe" --bin --abi -o ~/Documents/ContractABI/ --overwrite
## solc contracts/impl/Auth.sol --libraries "contracts/Recover.sol:Recover=0x21818Ceb3e8992F15899582456592307Ec08F1fe" --bin --abi -o ~/Documents/ContractABI/ --overwrite

solc --base-path . --include-path .. ../contracts/impl/GToken.sol --bin --abi -o ../abi/GToken --overwrite
solc --base-path . --include-path .. ../contracts/impl/Registry.sol --bin --abi -o ../abi/Registry --overwrite
solc --base-path . --include-path .. ../contracts/impl/Market.sol --bin --abi -o ../abi/Market --overwrite
solc --base-path . --include-path .. ../contracts/impl/Pledge.sol --bin --abi -o ../abi/Pledge --overwrite
solc --base-path . --include-path .. ../contracts/impl/Credit.sol --bin --abi -o ../abi/Credit --overwrite
solc --base-path . --include-path .. ../contracts/impl/Swap.sol --bin --abi -o ../abi/Swap --overwrite
solc --base-path . --include-path .. ../contracts/impl/Access.sol --bin --abi -o ../abi/Access --overwrite

abigen --out ../go/gtoken/gtoken.go --pkg gtoken --bin ../abi/GToken/GToken.bin --abi ../abi/GToken/GToken.abi
abigen --out ../go/registry/registry.go --pkg registry --bin ../abi/Registry/Registry.bin --abi ../abi/Registry/Registry.abi
abigen --out ../go/market/market.go --pkg market --bin ../abi/Market/Market.bin --abi ../abi/Market/Market.abi
abigen --out ../go/pledge/pledge.go --pkg pledge --bin ../abi/Pledge/Pledge.bin --abi ../abi/Pledge/Pledge.abi
abigen --out ../go/credit/credit.go --pkg credit --bin ../abi/Credit/Credit.bin --abi ../abi/Credit/Credit.abi
abigen --out ../go/swap/swap.go --pkg swap --bin ../abi/Swap/Swap.bin --abi ../abi/Swap/Swap.abi
abigen --out ../go/access/access.go --pkg access --bin ../abi/Access/Access.bin --abi ../abi/Access/Access.abi