############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help bindings mocks install-tests-deps tests tests-cover
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


gen: ## generate files (pk, vk, solidity verifier, go bindings)
	go generate ./...

clean: ## removes all generate files from gen
	rm -rf solidity/abi
	rm -rf solidity/*.{vk,pk}
	rm solidity/verifier_g16.sol solidity/verifier_g16_bindings.go

test-circuit: ## run circuit unit tests
	cd circuit && go test -v

test-contract: ## run contract verification unit tests
	cd solidity && go test -v

test-all: ## run all tests
	go test -v ./...