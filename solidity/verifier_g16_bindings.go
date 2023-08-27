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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f389592bbc15a25a30e8d360dd33bc376412f6c46da157359354be7969e4f1a864736f6c63430008150033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[3]\",\"name\":\"input\",\"type\":\"uint256[3]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611d08806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806311479fea14610030575b600080fd5b61004a60048036038101906100459190611562565b610060565b60405161005791906115e6565b60405180910390f35b600061006a6111c8565b60405180604001604052808760006002811061008957610088611601565b5b60200201518152602001876001600281106100a7576100a6611601565b5b6020020151815250816000018190525060405180604001604052806040518060400160405280886000600281106100e1576100e0611601565b5b60200201516000600281106100f9576100f8611601565b5b602002015181526020018860006002811061011757610116611601565b5b602002015160016002811061012f5761012e611601565b5b6020020151815250815260200160405180604001604052808860016002811061015b5761015a611601565b5b602002015160006002811061017357610172611601565b5b602002015181526020018860016002811061019157610190611601565b5b60200201516001600281106101a9576101a8611601565b5b602002015181525081525081602001819052506040518060400160405280856000600281106101db576101da611601565b5b60200201518152602001856001600281106101f9576101f8611601565b5b602002015181525081604001819052507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781600001516000015110610273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026a9061168d565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816000015160200151106102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d4906116f9565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160006002811061031a57610319611601565b5b60200201511061035f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035690611765565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516020015160006002811061039c5761039b611601565b5b6020020151106103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d8906117d1565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160016002811061041e5761041d611601565b5b602002015110610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a9061183d565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160200151602001516001600281106104a05761049f611601565b5b6020020151106104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc906118a9565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160400151600001511061054f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054690611915565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816040015160200151106105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b090611981565b60405180910390fd5b60005b6003811015610651577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018482600381106105f9576105f8611601565b5b60200201351061063e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610635906119ed565b60405180910390fd5b808061064990611a3c565b9150506105bc565b50600061065c610975565b90506000604051806040016040528060008152602001600081525090506106816111fb565b61068961121d565b6000604051806040016040528060008152602001600081525090507f2c1b7cc4c55e65f1dfa8d0271f7c58c317687a59348467f7a53382ae6c2c7d438460000181815250507e5d796fe65fe6cf47aeea5bcf107e6febb8f9aefd580fe6403582ee4b7a71618460200181815250507f0381ca9780997a022b7893cffc349e92fc18136c330ad934317649d18d3771b88260006003811061072c5761072b611601565b5b6020020181815250507f2fb10b23ef55ee23a86abcf7fbb2f51936a9fe8eecf43e2ed3c6fdc1dcb666768260016003811061076a57610769611601565b5b6020020181815250508760006003811061078757610786611601565b5b6020020135826002600381106107a05761079f611601565b5b6020020181815250506107b582828587610c2c565b7f2c86c2953ae290fb974dbebe5ec91b255399991f2cd8145032fbcbaf29ea4658826000600381106107ea576107e9611601565b5b6020020181815250507f2610e5b8b7733de5cf9486d81fd6d1853bfd0d5d35f9d5453cfcb472b1d547928260016003811061082857610827611601565b5b6020020181815250508760016003811061084557610844611601565b5b60200201358260026003811061085e5761085d611601565b5b60200201818152505061087382828587610c2c565b7f17c6359a5af92579c87faa049752435bafb35caedf76049699e7b1ee971ab1b6826000600381106108a8576108a7611601565b5b6020020181815250507f0ed07b16747d5f84cbea981283af56539a6eebdde972f2a3cf932fd013e43528826001600381106108e6576108e5611601565b5b6020020181815250508760026003811061090357610902611601565b5b60200201358260026003811061091c5761091b611601565b5b60200201818152505061093182828587610c2c565b6109656109418760000151610cce565b876020015187600001518860200151888a604001518c604001518c60600151610d8c565b9650505050505050949350505050565b61097d61123f565b60405180604001604052807f184428254274aebf8b38d3c35272f667365f15b3a44c64d499cb53d08da4c4ef81526020017f20ef73dca4d0e42c5fbc2846c2e43c076d5201698269ad28cb3a962a9bc6ad148152508160000181905250604051806040016040528060405180604001604052807f21786790018c0c355a6a3200e0026f9679fc3be1b547c6af8f20a3f135bdd24481526020017f0324e1a0aa92891fee3a7490de4c7c396d6c97b89cd09690dcb3771a3cb82644815250815260200160405180604001604052807f1d91921be2ed4f7f3d4fdd9b7f99b0f19c0876e29cc0b8b4c0982b20817f154881526020017f2e65d906065594e34d028524a02a704f4da7e032bcd3ee4e65f05e46e2defea38152508152508160200181905250604051806040016040528060405180604001604052807f23dc393ae7f9376eb931dc0eb2f35ad6fab731afb2dd76ea238bedebefb121d981526020017f160059af32fe9381d4d4f0871d65ed870994ea27d8f6e0c5bd3e341f42301f50815250815260200160405180604001604052807f268b12838a29764d02abe6aee1a174e05e164bafee3a16fc8c47a89b59c75ffd81526020017f056631b0163c7148cdd840b6c553ffb9737bf50a4699fa6f99226e4198709b0d8152508152508160400181905250604051806040016040528060405180604001604052807f0e92e45ab8e2d5471a6d7b87a634a53c0d52f11c41fe94ab751cbcde25f9295881526020017f242fcb49d9cebda175d518e84edff883bea6d448e88ef9d2c5028efaa6be604d815250815260200160405180604001604052807f073277faee7527b2388548f7634f004bfcccd8b01dd616f9acddeff1075e9adb81526020017f03598a719a12bf6ee06e32f1fc58ebddc1bf4794aa70e9db5ac39458f618fb28815250815250816060018190525090565b610c368484611102565b806000015182600060048110610c4f57610c4e611601565b5b602002018181525050806020015182600160048110610c7157610c70611601565b5b602002018181525050826000015182600260048110610c9357610c92611601565b5b602002018181525050826020015182600360048110610cb557610cb4611601565b5b602002018181525050610cc88282611165565b50505050565b610cd661127f565b60008260000151148015610cee575060008260200151145b15610d115760405180604001604052806000815260200160008152509050610d87565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478460200151610d569190611ab3565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610d819190611ae4565b81525090505b919050565b60008060405180608001604052808b8152602001898152602001878152602001858152509050600060405180608001604052808b815260200189815260200187815260200185815250905060006018905060008167ffffffffffffffff811115610df957610df8611328565b5b604051908082528060200260200182016040528015610e275781602001602082028036833780820191505090505b50905060005b6004811015611065576000600682610e459190611b18565b9050858260048110610e5a57610e59611601565b5b60200201516000015183600083610e719190611b5a565b81518110610e8257610e81611601565b5b602002602001018181525050858260048110610ea157610ea0611601565b5b60200201516020015183600183610eb89190611b5a565b81518110610ec957610ec8611601565b5b602002602001018181525050848260048110610ee857610ee7611601565b5b602002015160000151600060028110610f0457610f03611601565b5b602002015183600283610f179190611b5a565b81518110610f2857610f27611601565b5b602002602001018181525050848260048110610f4757610f46611601565b5b602002015160000151600160028110610f6357610f62611601565b5b602002015183600383610f769190611b5a565b81518110610f8757610f86611601565b5b602002602001018181525050848260048110610fa657610fa5611601565b5b602002015160200151600060028110610fc257610fc1611601565b5b602002015183600483610fd59190611b5a565b81518110610fe657610fe5611601565b5b60200260200101818152505084826004811061100557611004611601565b5b60200201516020015160016002811061102157611020611601565b5b6020020151836005836110349190611b5a565b8151811061104557611044611601565b5b60200260200101818152505050808061105d90611a3c565b915050610e2d565b5061106e611299565b6000602082602086026020860160086107d05a03fa9050806000810361109057fe5b50806110d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c890611bda565b60405180910390fd5b6000826000600181106110e7576110e6611601565b5b60200201511415965050505050505098975050505050505050565b600060608260808560076107d05a03fa9050806000810361111f57fe5b5080611160576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115790611c46565b60405180910390fd5b505050565b600060608260c08560066107d05a03fa9050806000810361118257fe5b50806111c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ba90611cb2565b60405180910390fd5b505050565b60405180606001604052806111db61127f565b81526020016111e86112bb565b81526020016111f561127f565b81525090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b604051806080016040528061125261127f565b815260200161125f6112bb565b815260200161126c6112bb565b81526020016112796112bb565b81525090565b604051806040016040528060008152602001600081525090565b6040518060200160405280600190602082028036833780820191505090505090565b60405180604001604052806112ce6112e1565b81526020016112db6112e1565b81525090565b6040518060400160405280600290602082028036833780820191505090505090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61136082611317565b810181811067ffffffffffffffff8211171561137f5761137e611328565b5b80604052505050565b6000611392611303565b905061139e8282611357565b919050565b600067ffffffffffffffff8211156113be576113bd611328565b5b602082029050919050565b600080fd5b6000819050919050565b6113e1816113ce565b81146113ec57600080fd5b50565b6000813590506113fe816113d8565b92915050565b6000611417611412846113a3565b611388565b90508060208402830185811115611431576114306113c9565b5b835b8181101561145a578061144688826113ef565b845260208401935050602081019050611433565b5050509392505050565b600082601f83011261147957611478611312565b5b6002611486848285611404565b91505092915050565b600067ffffffffffffffff8211156114aa576114a9611328565b5b602082029050919050565b60006114c86114c38461148f565b611388565b905080604084028301858111156114e2576114e16113c9565b5b835b8181101561150b57806114f78882611464565b8452602084019350506040810190506114e4565b5050509392505050565b600082601f83011261152a57611529611312565b5b60026115378482856114b5565b91505092915050565b60008190508260206003028201111561155c5761155b6113c9565b5b92915050565b600080600080610160858703121561157d5761157c61130d565b5b600061158b87828801611464565b945050604061159c87828801611515565b93505060c06115ad87828801611464565b9250506101006115bf87828801611540565b91505092959194509250565b60008115159050919050565b6115e0816115cb565b82525050565b60006020820190506115fb60008301846115d7565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f76657269666965722d61582d6774652d7072696d652d71000000000000000000600082015250565b6000611677601783611630565b915061168282611641565b602082019050919050565b600060208201905081810360008301526116a68161166a565b9050919050565b7f76657269666965722d61592d6774652d7072696d652d71000000000000000000600082015250565b60006116e3601783611630565b91506116ee826116ad565b602082019050919050565b60006020820190508181036000830152611712816116d6565b9050919050565b7f76657269666965722d6258302d6774652d7072696d652d710000000000000000600082015250565b600061174f601883611630565b915061175a82611719565b602082019050919050565b6000602082019050818103600083015261177e81611742565b9050919050565b7f76657269666965722d6259302d6774652d7072696d652d710000000000000000600082015250565b60006117bb601883611630565b91506117c682611785565b602082019050919050565b600060208201905081810360008301526117ea816117ae565b9050919050565b7f76657269666965722d6258312d6774652d7072696d652d710000000000000000600082015250565b6000611827601883611630565b9150611832826117f1565b602082019050919050565b600060208201905081810360008301526118568161181a565b9050919050565b7f76657269666965722d6259312d6774652d7072696d652d710000000000000000600082015250565b6000611893601883611630565b915061189e8261185d565b602082019050919050565b600060208201905081810360008301526118c281611886565b9050919050565b7f76657269666965722d63582d6774652d7072696d652d71000000000000000000600082015250565b60006118ff601783611630565b915061190a826118c9565b602082019050919050565b6000602082019050818103600083015261192e816118f2565b9050919050565b7f76657269666965722d63592d6774652d7072696d652d71000000000000000000600082015250565b600061196b601783611630565b915061197682611935565b602082019050919050565b6000602082019050818103600083015261199a8161195e565b9050919050565b7f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400600082015250565b60006119d7601f83611630565b91506119e2826119a1565b602082019050919050565b60006020820190508181036000830152611a06816119ca565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611a47826113ce565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a7957611a78611a0d565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611abe826113ce565b9150611ac9836113ce565b925082611ad957611ad8611a84565b5b828206905092915050565b6000611aef826113ce565b9150611afa836113ce565b9250828203905081811115611b1257611b11611a0d565b5b92915050565b6000611b23826113ce565b9150611b2e836113ce565b9250828202611b3c816113ce565b91508282048414831517611b5357611b52611a0d565b5b5092915050565b6000611b65826113ce565b9150611b70836113ce565b9250828201905080821115611b8857611b87611a0d565b5b92915050565b7f70616972696e672d6f70636f64652d6661696c65640000000000000000000000600082015250565b6000611bc4601583611630565b9150611bcf82611b8e565b602082019050919050565b60006020820190508181036000830152611bf381611bb7565b9050919050565b7f70616972696e672d6d756c2d6661696c65640000000000000000000000000000600082015250565b6000611c30601283611630565b9150611c3b82611bfa565b602082019050919050565b60006020820190508181036000830152611c5f81611c23565b9050919050565b7f70616972696e672d6164642d6661696c65640000000000000000000000000000600082015250565b6000611c9c601283611630565b9150611ca782611c66565b602082019050919050565b60006020820190508181036000830152611ccb81611c8f565b905091905056fea2646970667358221220517fc4361dc296645124b3a120bd1a47fae887b871bda8c0b743e40f009cc68c64736f6c63430008150033",
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

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[3] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[3] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [3]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[3] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [3]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}
