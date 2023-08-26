package solidity

//go:generate go run contract/main.go
//go:generate solc --evm-version paris --combined-json abi,bin verifier_g16.sol -o abi --overwrite
//go:generate abigen --combined-json abi/combined.json --pkg solidity --out verifier_g16_bindings.go
