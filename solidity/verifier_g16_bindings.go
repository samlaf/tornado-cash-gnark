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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cc9cdffd9ab24a0d130af7cd6628aee5ffa4547c513064d3ba21fc421cd8f35e64736f6c63430008150033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611c4a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f5c9d69e14610030575b600080fd5b61004a600480360381019061004591906114a4565b610060565b6040516100579190611528565b60405180910390f35b600061006a61110a565b60405180604001604052808760006002811061008957610088611543565b5b60200201518152602001876001600281106100a7576100a6611543565b5b6020020151815250816000018190525060405180604001604052806040518060400160405280886000600281106100e1576100e0611543565b5b60200201516000600281106100f9576100f8611543565b5b602002015181526020018860006002811061011757610116611543565b5b602002015160016002811061012f5761012e611543565b5b6020020151815250815260200160405180604001604052808860016002811061015b5761015a611543565b5b602002015160006002811061017357610172611543565b5b602002015181526020018860016002811061019157610190611543565b5b60200201516001600281106101a9576101a8611543565b5b602002015181525081525081602001819052506040518060400160405280856000600281106101db576101da611543565b5b60200201518152602001856001600281106101f9576101f8611543565b5b602002015181525081604001819052507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781600001516000015110610273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026a906115cf565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816000015160200151106102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d49061163b565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160006002811061031a57610319611543565b5b60200201511061035f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610356906116a7565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516020015160006002811061039c5761039b611543565b5b6020020151106103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611713565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160016002811061041e5761041d611543565b5b602002015110610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a9061177f565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160200151602001516001600281106104a05761049f611543565b5b6020020151106104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc906117eb565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160400151600001511061054f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054690611857565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816040015160200151106105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b0906118c3565b60405180910390fd5b60005b6002811015610651577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018482600281106105f9576105f8611543565b5b60200201351061063e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106359061192f565b60405180910390fd5b80806106499061197e565b9150506105bc565b50600061065c6108b8565b905060006040518060400160405280600081526020016000815250905061068161113d565b61068961115f565b6000604051806040016040528060008152602001600081525090507f061775c1c2c316bcc23b0addf38775b015a4c2dc29fc13ada26b68a35e5ccbf78460000181815250507f2d35eae924c6b97e06de2edcce4c5c02d2d44acc29fd4cb36df9acebbe17851e8460200181815250507f02c0e1014eee9d602bb8145175d2d38486f337e49f82357c55132b5a14223ae38260006003811061072d5761072c611543565b5b6020020181815250507f1b5ead385e9a164399e6858a0e4a30b29fa9215de27130de6bbc90710234b2878260016003811061076b5761076a611543565b5b6020020181815250508760006002811061078857610787611543565b5b6020020135826002600381106107a1576107a0611543565b5b6020020181815250506107b682828587610b6e565b7f06b86e686fc5afa9940015f8552635594e6033c697dd5c81ead52b7ed08ca8a3826000600381106107eb576107ea611543565b5b6020020181815250507f2876dd301f1961cc7594fc3378e9355e63f6297be0d18002337c97c013b69d9a8260016003811061082957610828611543565b5b6020020181815250508760016002811061084657610845611543565b5b60200201358260026003811061085f5761085e611543565b5b60200201818152505061087482828587610b6e565b6108a86108848760000151610c10565b876020015187600001518860200151888a604001518c604001518c60600151610cce565b9650505050505050949350505050565b6108c0611181565b60405180604001604052807f18a704164a5139da3580b7586eac5dc4c32ab49490562494533cfef51f668e1e81526020017f158dc5351a2badda9a34ea7a9a212c4a60103bcd79fc76261ac9cfaf20f2e2e48152508160000181905250604051806040016040528060405180604001604052807f28eb79a0c9fc3bd626572330efc542de6124cb0f610b7e775a46a078630c03f381526020017f302d43d1e997c40763a83ab74ad3dbbc3b40b868a09fb9f99b2ae83114837876815250815260200160405180604001604052807f119a950d9a5c56bf8200b4b77ca7c135355a9d3b296822e5f3391d1c7adc02e481526020017f0c6bb8d94b2bce220e546422547fe8166a2e5bc9315f76b8427b3c18af01f6ca8152508152508160200181905250604051806040016040528060405180604001604052807f2616e45cacaef34d46c945a5fc7da880b9fd9bac6a31592573541435daebfab781526020017f23f76e44f03d92d34e013e1a8347ffc552bfc7dc50c2cfda088cceeef1fbc9b6815250815260200160405180604001604052807f119316265b53da6e07fcadbf578185b20c51f6a15545d6603c9be6eb7456e9d381526020017f23da3e1e66ce16b6428d119d926b244762c1c7212f4cc14cdf597dd3df0840428152508152508160400181905250604051806040016040528060405180604001604052807f15a7fe9adb233dd32597f8d5ab75c1980d67f9957cce70613f1adaca19ad844a81526020017f084d537c4269d5b3efb0add599c7f71b9391c6cf1405b934cb07ac13b9783541815250815260200160405180604001604052807f148a84f4ca40fb64bfe5e81aca80ab034cde52488faf95893545fb785b2e744d81526020017e4c1c809ff1000568be302e5741306563a59cf60836cefdef4686d5b6ab14cf815250815250816060018190525090565b610b788484611044565b806000015182600060048110610b9157610b90611543565b5b602002018181525050806020015182600160048110610bb357610bb2611543565b5b602002018181525050826000015182600260048110610bd557610bd4611543565b5b602002018181525050826020015182600360048110610bf757610bf6611543565b5b602002018181525050610c0a82826110a7565b50505050565b610c186111c1565b60008260000151148015610c30575060008260200151145b15610c535760405180604001604052806000815260200160008152509050610cc9565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478460200151610c9891906119f5565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610cc39190611a26565b81525090505b919050565b60008060405180608001604052808b8152602001898152602001878152602001858152509050600060405180608001604052808b815260200189815260200187815260200185815250905060006018905060008167ffffffffffffffff811115610d3b57610d3a61126a565b5b604051908082528060200260200182016040528015610d695781602001602082028036833780820191505090505b50905060005b6004811015610fa7576000600682610d879190611a5a565b9050858260048110610d9c57610d9b611543565b5b60200201516000015183600083610db39190611a9c565b81518110610dc457610dc3611543565b5b602002602001018181525050858260048110610de357610de2611543565b5b60200201516020015183600183610dfa9190611a9c565b81518110610e0b57610e0a611543565b5b602002602001018181525050848260048110610e2a57610e29611543565b5b602002015160000151600060028110610e4657610e45611543565b5b602002015183600283610e599190611a9c565b81518110610e6a57610e69611543565b5b602002602001018181525050848260048110610e8957610e88611543565b5b602002015160000151600160028110610ea557610ea4611543565b5b602002015183600383610eb89190611a9c565b81518110610ec957610ec8611543565b5b602002602001018181525050848260048110610ee857610ee7611543565b5b602002015160200151600060028110610f0457610f03611543565b5b602002015183600483610f179190611a9c565b81518110610f2857610f27611543565b5b602002602001018181525050848260048110610f4757610f46611543565b5b602002015160200151600160028110610f6357610f62611543565b5b602002015183600583610f769190611a9c565b81518110610f8757610f86611543565b5b602002602001018181525050508080610f9f9061197e565b915050610d6f565b50610fb06111db565b6000602082602086026020860160086107d05a03fa90508060008103610fd257fe5b5080611013576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100a90611b1c565b60405180910390fd5b60008260006001811061102957611028611543565b5b60200201511415965050505050505098975050505050505050565b600060608260808560076107d05a03fa9050806000810361106157fe5b50806110a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109990611b88565b60405180910390fd5b505050565b600060608260c08560066107d05a03fa905080600081036110c457fe5b5080611105576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110fc90611bf4565b60405180910390fd5b505050565b604051806060016040528061111d6111c1565b815260200161112a6111fd565b81526020016111376111c1565b81525090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b60405180608001604052806111946111c1565b81526020016111a16111fd565b81526020016111ae6111fd565b81526020016111bb6111fd565b81525090565b604051806040016040528060008152602001600081525090565b6040518060200160405280600190602082028036833780820191505090505090565b6040518060400160405280611210611223565b815260200161121d611223565b81525090565b6040518060400160405280600290602082028036833780820191505090505090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6112a282611259565b810181811067ffffffffffffffff821117156112c1576112c061126a565b5b80604052505050565b60006112d4611245565b90506112e08282611299565b919050565b600067ffffffffffffffff821115611300576112ff61126a565b5b602082029050919050565b600080fd5b6000819050919050565b61132381611310565b811461132e57600080fd5b50565b6000813590506113408161131a565b92915050565b6000611359611354846112e5565b6112ca565b905080602084028301858111156113735761137261130b565b5b835b8181101561139c57806113888882611331565b845260208401935050602081019050611375565b5050509392505050565b600082601f8301126113bb576113ba611254565b5b60026113c8848285611346565b91505092915050565b600067ffffffffffffffff8211156113ec576113eb61126a565b5b602082029050919050565b600061140a611405846113d1565b6112ca565b905080604084028301858111156114245761142361130b565b5b835b8181101561144d578061143988826113a6565b845260208401935050604081019050611426565b5050509392505050565b600082601f83011261146c5761146b611254565b5b60026114798482856113f7565b91505092915050565b60008190508260206002028201111561149e5761149d61130b565b5b92915050565b60008060008061014085870312156114bf576114be61124f565b5b60006114cd878288016113a6565b94505060406114de87828801611457565b93505060c06114ef878288016113a6565b92505061010061150187828801611482565b91505092959194509250565b60008115159050919050565b6115228161150d565b82525050565b600060208201905061153d6000830184611519565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f76657269666965722d61582d6774652d7072696d652d71000000000000000000600082015250565b60006115b9601783611572565b91506115c482611583565b602082019050919050565b600060208201905081810360008301526115e8816115ac565b9050919050565b7f76657269666965722d61592d6774652d7072696d652d71000000000000000000600082015250565b6000611625601783611572565b9150611630826115ef565b602082019050919050565b6000602082019050818103600083015261165481611618565b9050919050565b7f76657269666965722d6258302d6774652d7072696d652d710000000000000000600082015250565b6000611691601883611572565b915061169c8261165b565b602082019050919050565b600060208201905081810360008301526116c081611684565b9050919050565b7f76657269666965722d6259302d6774652d7072696d652d710000000000000000600082015250565b60006116fd601883611572565b9150611708826116c7565b602082019050919050565b6000602082019050818103600083015261172c816116f0565b9050919050565b7f76657269666965722d6258312d6774652d7072696d652d710000000000000000600082015250565b6000611769601883611572565b915061177482611733565b602082019050919050565b600060208201905081810360008301526117988161175c565b9050919050565b7f76657269666965722d6259312d6774652d7072696d652d710000000000000000600082015250565b60006117d5601883611572565b91506117e08261179f565b602082019050919050565b60006020820190508181036000830152611804816117c8565b9050919050565b7f76657269666965722d63582d6774652d7072696d652d71000000000000000000600082015250565b6000611841601783611572565b915061184c8261180b565b602082019050919050565b6000602082019050818103600083015261187081611834565b9050919050565b7f76657269666965722d63592d6774652d7072696d652d71000000000000000000600082015250565b60006118ad601783611572565b91506118b882611877565b602082019050919050565b600060208201905081810360008301526118dc816118a0565b9050919050565b7f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400600082015250565b6000611919601f83611572565b9150611924826118e3565b602082019050919050565b600060208201905081810360008301526119488161190c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061198982611310565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119bb576119ba61194f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611a0082611310565b9150611a0b83611310565b925082611a1b57611a1a6119c6565b5b828206905092915050565b6000611a3182611310565b9150611a3c83611310565b9250828203905081811115611a5457611a5361194f565b5b92915050565b6000611a6582611310565b9150611a7083611310565b9250828202611a7e81611310565b91508282048414831517611a9557611a9461194f565b5b5092915050565b6000611aa782611310565b9150611ab283611310565b9250828201905080821115611aca57611ac961194f565b5b92915050565b7f70616972696e672d6f70636f64652d6661696c65640000000000000000000000600082015250565b6000611b06601583611572565b9150611b1182611ad0565b602082019050919050565b60006020820190508181036000830152611b3581611af9565b9050919050565b7f70616972696e672d6d756c2d6661696c65640000000000000000000000000000600082015250565b6000611b72601283611572565b9150611b7d82611b3c565b602082019050919050565b60006020820190508181036000830152611ba181611b65565b9050919050565b7f70616972696e672d6164642d6661696c65640000000000000000000000000000600082015250565b6000611bde601283611572565b9150611be982611ba8565b602082019050919050565b60006020820190508181036000830152611c0d81611bd1565b905091905056fea26469706673582212206cd29dd576c942ec25c0eae8d5633f15ef6d692f46254863383d99a44396867664736f6c63430008150033",
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

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}
