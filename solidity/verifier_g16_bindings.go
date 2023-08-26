// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PairingMetaData contains all meta data concerning the Pairing contract.
var PairingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206d40ce6c1737e79d62670d7ba2126a75abe19c7572dbe371d43dff6a4e8d749264736f6c63430008150033",
}

// PairingABI is the input ABI used to generate the binding from.
// Deprecated: Use PairingMetaData.ABI instead.
var PairingABI = PairingMetaData.ABI

// PairingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PairingMetaData.Bin instead.
var PairingBin = PairingMetaData.Bin

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// VerifierMetaData contains all meta data concerning the Verifier contract.
var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b8d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806343753b4d14610030575b600080fd5b61004a600480360381019061004591906113e7565b610060565b604051610057919061146b565b60405180910390f35b600061006a61104d565b60405180604001604052808760006002811061008957610088611486565b5b60200201518152602001876001600281106100a7576100a6611486565b5b6020020151815250816000018190525060405180604001604052806040518060400160405280886000600281106100e1576100e0611486565b5b60200201516000600281106100f9576100f8611486565b5b602002015181526020018860006002811061011757610116611486565b5b602002015160016002811061012f5761012e611486565b5b6020020151815250815260200160405180604001604052808860016002811061015b5761015a611486565b5b602002015160006002811061017357610172611486565b5b602002015181526020018860016002811061019157610190611486565b5b60200201516001600281106101a9576101a8611486565b5b602002015181525081525081602001819052506040518060400160405280856000600281106101db576101da611486565b5b60200201518152602001856001600281106101f9576101f8611486565b5b602002015181525081604001819052507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781600001516000015110610273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026a90611512565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816000015160200151106102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d49061157e565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160006002811061031a57610319611486565b5b60200201511061035f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610356906115ea565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516020015160006002811061039c5761039b611486565b5b6020020151106103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611656565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160016002811061041e5761041d611486565b5b602002015110610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a906116c2565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160200151602001516001600281106104a05761049f611486565b5b6020020151106104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc9061172e565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160400151600001511061054f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105469061179a565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816040015160200151106105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b090611806565b60405180910390fd5b60005b6001811015610651577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018482600181106105f9576105f8611486565b5b60200201351061063e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063590611872565b60405180910390fd5b8080610649906118c1565b9150506105bc565b50600061065c6107fa565b9050600060405180604001604052806000815260200160008152509050610681611080565b6106896110a2565b6000604051806040016040528060008152602001600081525090507f0358c3a2fd1ae5575da891c869dad67726aa760b7083b97a43b09720f71e93408460000181815250507f112da6123be5b05db5f43b6f39bd9ca0b841543ce764d46dd12a10173a7b89f78460200181815250507f1128fac33b354b214f1bc999d968167281a21fc8f0a6e5c0c1ebff19f16e28568260006003811061072d5761072c611486565b5b6020020181815250507f1096bb5619c49d35829fa977d23560acb7c4b21131fbdfb576b3749dc8b23e9a8260016003811061076b5761076a611486565b5b6020020181815250508760006001811061078857610787611486565b5b6020020135826002600381106107a1576107a0611486565b5b6020020181815250506107b682828587610ab1565b6107ea6107c68760000151610b53565b876020015187600001518860200151888a604001518c604001518c60600151610c11565b9650505050505050949350505050565b6108026110c4565b60405180604001604052807f2d665e52dfd5da1539ae1e94b652b86c6fdff3a8f90c51f8a683bf62eed98cfe81526020017f0f186a3104a8448b9b67d4028c17a707b270784d988c405b5943a6a67b8316f88152508160000181905250604051806040016040528060405180604001604052807f12c40932c631bbaa18e45ac86e0f6974b95904b6abbd0ad69ed29b3f9554b67e81526020017f03ed7ec6917390fc1ad1f76e8b4198d8dbc4dc3cf4963682f5cb92a3fbf731dc815250815260200160405180604001604052807f095be75e32a09d91970cd2e8c784894dc598b82d6d672fd1e52f6498a4c1fc0a81526020017f29894c0ae12d73c4b99f6a9001a69caa3a8cb2a8cd7c3566a42e89927a8ae0478152508152508160200181905250604051806040016040528060405180604001604052807f26cbc2ade2a16bcb67a5793440815f948ecef18f69bd561d567b7e30129b5a9e81526020017f27b516abb68fc09b2f0862f98fd6ba71e2086d744794f47cc5a952ae7b3da32b815250815260200160405180604001604052807f26ee9df4dc436fd3f63b4aeaa9b0f3b809856a349c8c20b5107f3de47e0ad8d481526020017f073f39005e2851b1c7b9fd2213a15a62a1f484642703272b3c3b0c447ff3f1b58152508152508160400181905250604051806040016040528060405180604001604052807f15d361fed101638c52c6433170dd56781555c3a91f7da944b39f0b7e64dbfe9581526020017f1e5d910e5bc8a4317f0677ad9f628488e4536691bbc697a56ad84fa4a09ca8c1815250815260200160405180604001604052807f023e3c7cb945eb4cd4e186d2412265c8a32148f6d10e124074a6917a2efe6ce381526020017f0dfee8bb797371856425140d29d79e5e17e88468bbb0d02bcbf728907ddaea3d815250815250816060018190525090565b610abb8484610f87565b806000015182600060048110610ad457610ad3611486565b5b602002018181525050806020015182600160048110610af657610af5611486565b5b602002018181525050826000015182600260048110610b1857610b17611486565b5b602002018181525050826020015182600360048110610b3a57610b39611486565b5b602002018181525050610b4d8282610fea565b50505050565b610b5b611104565b60008260000151148015610b73575060008260200151145b15610b965760405180604001604052806000815260200160008152509050610c0c565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478460200151610bdb9190611938565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610c069190611969565b81525090505b919050565b60008060405180608001604052808b8152602001898152602001878152602001858152509050600060405180608001604052808b815260200189815260200187815260200185815250905060006018905060008167ffffffffffffffff811115610c7e57610c7d6111ad565b5b604051908082528060200260200182016040528015610cac5781602001602082028036833780820191505090505b50905060005b6004811015610eea576000600682610cca919061199d565b9050858260048110610cdf57610cde611486565b5b60200201516000015183600083610cf691906119df565b81518110610d0757610d06611486565b5b602002602001018181525050858260048110610d2657610d25611486565b5b60200201516020015183600183610d3d91906119df565b81518110610d4e57610d4d611486565b5b602002602001018181525050848260048110610d6d57610d6c611486565b5b602002015160000151600060028110610d8957610d88611486565b5b602002015183600283610d9c91906119df565b81518110610dad57610dac611486565b5b602002602001018181525050848260048110610dcc57610dcb611486565b5b602002015160000151600160028110610de857610de7611486565b5b602002015183600383610dfb91906119df565b81518110610e0c57610e0b611486565b5b602002602001018181525050848260048110610e2b57610e2a611486565b5b602002015160200151600060028110610e4757610e46611486565b5b602002015183600483610e5a91906119df565b81518110610e6b57610e6a611486565b5b602002602001018181525050848260048110610e8a57610e89611486565b5b602002015160200151600160028110610ea657610ea5611486565b5b602002015183600583610eb991906119df565b81518110610eca57610ec9611486565b5b602002602001018181525050508080610ee2906118c1565b915050610cb2565b50610ef361111e565b6000602082602086026020860160086107d05a03fa90508060008103610f1557fe5b5080610f56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4d90611a5f565b60405180910390fd5b600082600060018110610f6c57610f6b611486565b5b60200201511415965050505050505098975050505050505050565b600060608260808560076107d05a03fa90508060008103610fa457fe5b5080610fe5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fdc90611acb565b60405180910390fd5b505050565b600060608260c08560066107d05a03fa9050806000810361100757fe5b5080611048576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103f90611b37565b60405180910390fd5b505050565b6040518060600160405280611060611104565b815260200161106d611140565b815260200161107a611104565b81525090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b60405180608001604052806110d7611104565b81526020016110e4611140565b81526020016110f1611140565b81526020016110fe611140565b81525090565b604051806040016040528060008152602001600081525090565b6040518060200160405280600190602082028036833780820191505090505090565b6040518060400160405280611153611166565b8152602001611160611166565b81525090565b6040518060400160405280600290602082028036833780820191505090505090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6111e58261119c565b810181811067ffffffffffffffff82111715611204576112036111ad565b5b80604052505050565b6000611217611188565b905061122382826111dc565b919050565b600067ffffffffffffffff821115611243576112426111ad565b5b602082029050919050565b600080fd5b6000819050919050565b61126681611253565b811461127157600080fd5b50565b6000813590506112838161125d565b92915050565b600061129c61129784611228565b61120d565b905080602084028301858111156112b6576112b561124e565b5b835b818110156112df57806112cb8882611274565b8452602084019350506020810190506112b8565b5050509392505050565b600082601f8301126112fe576112fd611197565b5b600261130b848285611289565b91505092915050565b600067ffffffffffffffff82111561132f5761132e6111ad565b5b602082029050919050565b600061134d61134884611314565b61120d565b905080604084028301858111156113675761136661124e565b5b835b81811015611390578061137c88826112e9565b845260208401935050604081019050611369565b5050509392505050565b600082601f8301126113af576113ae611197565b5b60026113bc84828561133a565b91505092915050565b6000819050826020600102820111156113e1576113e061124e565b5b92915050565b600080600080610120858703121561140257611401611192565b5b6000611410878288016112e9565b94505060406114218782880161139a565b93505060c0611432878288016112e9565b925050610100611444878288016113c5565b91505092959194509250565b60008115159050919050565b61146581611450565b82525050565b6000602082019050611480600083018461145c565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f76657269666965722d61582d6774652d7072696d652d71000000000000000000600082015250565b60006114fc6017836114b5565b9150611507826114c6565b602082019050919050565b6000602082019050818103600083015261152b816114ef565b9050919050565b7f76657269666965722d61592d6774652d7072696d652d71000000000000000000600082015250565b60006115686017836114b5565b915061157382611532565b602082019050919050565b600060208201905081810360008301526115978161155b565b9050919050565b7f76657269666965722d6258302d6774652d7072696d652d710000000000000000600082015250565b60006115d46018836114b5565b91506115df8261159e565b602082019050919050565b60006020820190508181036000830152611603816115c7565b9050919050565b7f76657269666965722d6259302d6774652d7072696d652d710000000000000000600082015250565b60006116406018836114b5565b915061164b8261160a565b602082019050919050565b6000602082019050818103600083015261166f81611633565b9050919050565b7f76657269666965722d6258312d6774652d7072696d652d710000000000000000600082015250565b60006116ac6018836114b5565b91506116b782611676565b602082019050919050565b600060208201905081810360008301526116db8161169f565b9050919050565b7f76657269666965722d6259312d6774652d7072696d652d710000000000000000600082015250565b60006117186018836114b5565b9150611723826116e2565b602082019050919050565b600060208201905081810360008301526117478161170b565b9050919050565b7f76657269666965722d63582d6774652d7072696d652d71000000000000000000600082015250565b60006117846017836114b5565b915061178f8261174e565b602082019050919050565b600060208201905081810360008301526117b381611777565b9050919050565b7f76657269666965722d63592d6774652d7072696d652d71000000000000000000600082015250565b60006117f06017836114b5565b91506117fb826117ba565b602082019050919050565b6000602082019050818103600083015261181f816117e3565b9050919050565b7f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400600082015250565b600061185c601f836114b5565b915061186782611826565b602082019050919050565b6000602082019050818103600083015261188b8161184f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118cc82611253565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118fe576118fd611892565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061194382611253565b915061194e83611253565b92508261195e5761195d611909565b5b828206905092915050565b600061197482611253565b915061197f83611253565b925082820390508181111561199757611996611892565b5b92915050565b60006119a882611253565b91506119b383611253565b92508282026119c181611253565b915082820484148315176119d8576119d7611892565b5b5092915050565b60006119ea82611253565b91506119f583611253565b9250828201905080821115611a0d57611a0c611892565b5b92915050565b7f70616972696e672d6f70636f64652d6661696c65640000000000000000000000600082015250565b6000611a496015836114b5565b9150611a5482611a13565b602082019050919050565b60006020820190508181036000830152611a7881611a3c565b9050919050565b7f70616972696e672d6d756c2d6661696c65640000000000000000000000000000600082015250565b6000611ab56012836114b5565b9150611ac082611a7f565b602082019050919050565b60006020820190508181036000830152611ae481611aa8565b9050919050565b7f70616972696e672d6164642d6661696c65640000000000000000000000000000600082015250565b6000611b216012836114b5565b9150611b2c82611aeb565b602082019050919050565b60006020820190508181036000830152611b5081611b14565b905091905056fea264697066735822122052181d13ed6e0f6579e6475e49ad6d9e335449de7307652f9ca87d7a5f23518d64736f6c63430008150033",
}

// VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierMetaData.ABI instead.
var VerifierABI = VerifierMetaData.ABI

// VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierMetaData.Bin instead.
var VerifierBin = VerifierMetaData.Bin

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}