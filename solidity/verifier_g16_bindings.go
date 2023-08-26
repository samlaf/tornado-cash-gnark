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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c5f9019157fd3fe134f43421ce03aefa86687c1453e181d0c31af3b59fcb0b9064736f6c63430008150033",
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
	Bin: "0x608060405234801561001057600080fd5b50611b8b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806343753b4d14610030575b600080fd5b61004a600480360381019061004591906113e5565b610060565b6040516100579190611469565b60405180910390f35b600061006a61104b565b60405180604001604052808760006002811061008957610088611484565b5b60200201518152602001876001600281106100a7576100a6611484565b5b6020020151815250816000018190525060405180604001604052806040518060400160405280886000600281106100e1576100e0611484565b5b60200201516000600281106100f9576100f8611484565b5b602002015181526020018860006002811061011757610116611484565b5b602002015160016002811061012f5761012e611484565b5b6020020151815250815260200160405180604001604052808860016002811061015b5761015a611484565b5b602002015160006002811061017357610172611484565b5b602002015181526020018860016002811061019157610190611484565b5b60200201516001600281106101a9576101a8611484565b5b602002015181525081525081602001819052506040518060400160405280856000600281106101db576101da611484565b5b60200201518152602001856001600281106101f9576101f8611484565b5b602002015181525081604001819052507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781600001516000015110610273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026a90611510565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816000015160200151106102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d49061157c565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160006002811061031a57610319611484565b5b60200201511061035f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610356906115e8565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516020015160006002811061039c5761039b611484565b5b6020020151106103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611654565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160016002811061041e5761041d611484565b5b602002015110610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a906116c0565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160200151602001516001600281106104a05761049f611484565b5b6020020151106104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc9061172c565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160400151600001511061054f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054690611798565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816040015160200151106105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b090611804565b60405180910390fd5b60005b6001811015610651577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018482600181106105f9576105f8611484565b5b60200201351061063e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063590611870565b60405180910390fd5b8080610649906118bf565b9150506105bc565b50600061065c6107fa565b905060006040518060400160405280600081526020016000815250905061068161107e565b6106896110a0565b6000604051806040016040528060008152602001600081525090507f1f98f92fc64ededc98941be5298297611f269b6d77b232ca95d080b6b80e4f788460000181815250507f2a21d8652fef257ce92fcc1a61c37f19d7322082930b4ec471b7882a1d9d6a928460200181815250507f25e998d29d433daaa572a8784d1e5ad80131ce4d03b22e18c20637bdd3fcd1d68260006003811061072d5761072c611484565b5b6020020181815250507f027a999b777ae24fad862b422dc5de4a14d599d54011189820f0e8e35887c9b18260016003811061076b5761076a611484565b5b6020020181815250508760006001811061078857610787611484565b5b6020020135826002600381106107a1576107a0611484565b5b6020020181815250506107b682828587610aaf565b6107ea6107c68760000151610b51565b876020015187600001518860200151888a604001518c604001518c60600151610c0f565b9650505050505050949350505050565b6108026110c2565b60405180604001604052807f1938577543dca51ad865de535db42f9c1b0930c156f52180a8df9162184a321681526020017f1149a781df2ac851f9c5f53518dc415795bc53b1bf7df04727358bc141bbe1818152508160000181905250604051806040016040528060405180604001604052807f259e91caf7fef28e177ab826cf28308d36d44134ed27f5a117ce21364b54c8ff81526020017e2aa0dbf986d67f8599a628488e6f2a7b6a849f00ce4099fdff9267e945e3f4815250815260200160405180604001604052807f2876ef1051af8cfb14d56ee21311854b9b293d355c9837d384dcadd4208737f381526020017f019dd3978f2e8fcbda6e32b34615fd51aa1946c9b85e89da5fbb06f7393d0b0a8152508152508160200181905250604051806040016040528060405180604001604052807f16127450e15ef36f8944860a40e92f14d92294244df385f4afe3ec900184910481526020017e8cbc6c872c96c4b12320fe486532bfa3aba8c12ae523063b188c870ca97d8b815250815260200160405180604001604052807f23a38c49a05b7508a44409538730ad5cf69c03fe38439c13513f355bdcdbd87d81526020017f20c7aefa6d36f8c2c921c4eef08f1183b02045c79bfaec1556257cbfd83d971b8152508152508160400181905250604051806040016040528060405180604001604052807f1e318414a288d9dcb939b5a7afe65f14ad02fa9f596ba8b78bf92ee1c411ba7881526020017f103c9d808b04a089c113a980e5f5cb619511fa3c13c36d7540ec3b75e02a8423815250815260200160405180604001604052807f0e8a452db2a2aa752a52224c2e6fbd54f049ac88b346911b5bf320fe660119e381526020017f1fdf46ec513216474302a3290129889d431bb9dbe2beaf1e46f7dddb7e3ebad0815250815250816060018190525090565b610ab98484610f85565b806000015182600060048110610ad257610ad1611484565b5b602002018181525050806020015182600160048110610af457610af3611484565b5b602002018181525050826000015182600260048110610b1657610b15611484565b5b602002018181525050826020015182600360048110610b3857610b37611484565b5b602002018181525050610b4b8282610fe8565b50505050565b610b59611102565b60008260000151148015610b71575060008260200151145b15610b945760405180604001604052806000815260200160008152509050610c0a565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478460200151610bd99190611936565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610c049190611967565b81525090505b919050565b60008060405180608001604052808b8152602001898152602001878152602001858152509050600060405180608001604052808b815260200189815260200187815260200185815250905060006018905060008167ffffffffffffffff811115610c7c57610c7b6111ab565b5b604051908082528060200260200182016040528015610caa5781602001602082028036833780820191505090505b50905060005b6004811015610ee8576000600682610cc8919061199b565b9050858260048110610cdd57610cdc611484565b5b60200201516000015183600083610cf491906119dd565b81518110610d0557610d04611484565b5b602002602001018181525050858260048110610d2457610d23611484565b5b60200201516020015183600183610d3b91906119dd565b81518110610d4c57610d4b611484565b5b602002602001018181525050848260048110610d6b57610d6a611484565b5b602002015160000151600060028110610d8757610d86611484565b5b602002015183600283610d9a91906119dd565b81518110610dab57610daa611484565b5b602002602001018181525050848260048110610dca57610dc9611484565b5b602002015160000151600160028110610de657610de5611484565b5b602002015183600383610df991906119dd565b81518110610e0a57610e09611484565b5b602002602001018181525050848260048110610e2957610e28611484565b5b602002015160200151600060028110610e4557610e44611484565b5b602002015183600483610e5891906119dd565b81518110610e6957610e68611484565b5b602002602001018181525050848260048110610e8857610e87611484565b5b602002015160200151600160028110610ea457610ea3611484565b5b602002015183600583610eb791906119dd565b81518110610ec857610ec7611484565b5b602002602001018181525050508080610ee0906118bf565b915050610cb0565b50610ef161111c565b6000602082602086026020860160086107d05a03fa90508060008103610f1357fe5b5080610f54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4b90611a5d565b60405180910390fd5b600082600060018110610f6a57610f69611484565b5b60200201511415965050505050505098975050505050505050565b600060608260808560076107d05a03fa90508060008103610fa257fe5b5080610fe3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fda90611ac9565b60405180910390fd5b505050565b600060608260c08560066107d05a03fa9050806000810361100557fe5b5080611046576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103d90611b35565b60405180910390fd5b505050565b604051806060016040528061105e611102565b815260200161106b61113e565b8152602001611078611102565b81525090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b60405180608001604052806110d5611102565b81526020016110e261113e565b81526020016110ef61113e565b81526020016110fc61113e565b81525090565b604051806040016040528060008152602001600081525090565b6040518060200160405280600190602082028036833780820191505090505090565b6040518060400160405280611151611164565b815260200161115e611164565b81525090565b6040518060400160405280600290602082028036833780820191505090505090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6111e38261119a565b810181811067ffffffffffffffff82111715611202576112016111ab565b5b80604052505050565b6000611215611186565b905061122182826111da565b919050565b600067ffffffffffffffff821115611241576112406111ab565b5b602082029050919050565b600080fd5b6000819050919050565b61126481611251565b811461126f57600080fd5b50565b6000813590506112818161125b565b92915050565b600061129a61129584611226565b61120b565b905080602084028301858111156112b4576112b361124c565b5b835b818110156112dd57806112c98882611272565b8452602084019350506020810190506112b6565b5050509392505050565b600082601f8301126112fc576112fb611195565b5b6002611309848285611287565b91505092915050565b600067ffffffffffffffff82111561132d5761132c6111ab565b5b602082029050919050565b600061134b61134684611312565b61120b565b905080604084028301858111156113655761136461124c565b5b835b8181101561138e578061137a88826112e7565b845260208401935050604081019050611367565b5050509392505050565b600082601f8301126113ad576113ac611195565b5b60026113ba848285611338565b91505092915050565b6000819050826020600102820111156113df576113de61124c565b5b92915050565b6000806000806101208587031215611400576113ff611190565b5b600061140e878288016112e7565b945050604061141f87828801611398565b93505060c0611430878288016112e7565b925050610100611442878288016113c3565b91505092959194509250565b60008115159050919050565b6114638161144e565b82525050565b600060208201905061147e600083018461145a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f76657269666965722d61582d6774652d7072696d652d71000000000000000000600082015250565b60006114fa6017836114b3565b9150611505826114c4565b602082019050919050565b60006020820190508181036000830152611529816114ed565b9050919050565b7f76657269666965722d61592d6774652d7072696d652d71000000000000000000600082015250565b60006115666017836114b3565b915061157182611530565b602082019050919050565b6000602082019050818103600083015261159581611559565b9050919050565b7f76657269666965722d6258302d6774652d7072696d652d710000000000000000600082015250565b60006115d26018836114b3565b91506115dd8261159c565b602082019050919050565b60006020820190508181036000830152611601816115c5565b9050919050565b7f76657269666965722d6259302d6774652d7072696d652d710000000000000000600082015250565b600061163e6018836114b3565b915061164982611608565b602082019050919050565b6000602082019050818103600083015261166d81611631565b9050919050565b7f76657269666965722d6258312d6774652d7072696d652d710000000000000000600082015250565b60006116aa6018836114b3565b91506116b582611674565b602082019050919050565b600060208201905081810360008301526116d98161169d565b9050919050565b7f76657269666965722d6259312d6774652d7072696d652d710000000000000000600082015250565b60006117166018836114b3565b9150611721826116e0565b602082019050919050565b6000602082019050818103600083015261174581611709565b9050919050565b7f76657269666965722d63582d6774652d7072696d652d71000000000000000000600082015250565b60006117826017836114b3565b915061178d8261174c565b602082019050919050565b600060208201905081810360008301526117b181611775565b9050919050565b7f76657269666965722d63592d6774652d7072696d652d71000000000000000000600082015250565b60006117ee6017836114b3565b91506117f9826117b8565b602082019050919050565b6000602082019050818103600083015261181d816117e1565b9050919050565b7f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400600082015250565b600061185a601f836114b3565b915061186582611824565b602082019050919050565b600060208201905081810360008301526118898161184d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118ca82611251565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118fc576118fb611890565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061194182611251565b915061194c83611251565b92508261195c5761195b611907565b5b828206905092915050565b600061197282611251565b915061197d83611251565b925082820390508181111561199557611994611890565b5b92915050565b60006119a682611251565b91506119b183611251565b92508282026119bf81611251565b915082820484148315176119d6576119d5611890565b5b5092915050565b60006119e882611251565b91506119f383611251565b9250828201905080821115611a0b57611a0a611890565b5b92915050565b7f70616972696e672d6f70636f64652d6661696c65640000000000000000000000600082015250565b6000611a476015836114b3565b9150611a5282611a11565b602082019050919050565b60006020820190508181036000830152611a7681611a3a565b9050919050565b7f70616972696e672d6d756c2d6661696c65640000000000000000000000000000600082015250565b6000611ab36012836114b3565b9150611abe82611a7d565b602082019050919050565b60006020820190508181036000830152611ae281611aa6565b9050919050565b7f70616972696e672d6164642d6661696c65640000000000000000000000000000600082015250565b6000611b1f6012836114b3565b9150611b2a82611ae9565b602082019050919050565b60006020820190508181036000830152611b4e81611b12565b905091905056fea264697066735822122095580c45265341de5324f773e9f055b0972d309c3505ca3b373ca3857d14958764736f6c63430008150033",
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
