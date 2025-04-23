// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// VRFCoordinatorV2MetaData contains all meta data concerning the VRFCoordinatorV2 contract.
var VRFCoordinatorV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"numWords\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"numWords\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RandomWordsFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomWordsRequested\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"numWords\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061082d8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806338ba46141461004e57806381d12c581461006a578063973a814e1461009c578063e726f2e1146100ba575b5f5ffd5b61006860048036038101906100639190610477565b6100ea565b005b610084600480360381019061007f91906104d4565b610236565b60405161009393929190610576565b60405180910390f35b6100a4610295565b6040516100b191906105ba565b60405180910390f35b6100d460048036038101906100cf91906105fd565b61029b565b6040516100e191906105ba565b60405180910390f35b5f5f5f8581526020019081526020015f209050805f0160189054906101000a900460ff161561014e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014590610682565b60405180910390fd5b6001815f0160186101000a81548160ff021916908315150217905550805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631fe543e38585856040518463ffffffff1660e01b81526004016101c993929190610718565b5f604051808303815f87803b1580156101e0575f5ffd5b505af11580156101f2573d5f5f3e3d5ffd5b50505050837ff45ee76115b0ed5f4ebe293254449fbe612bad36a53d52b87b6a40687adc48de8484604051610228929190610748565b60405180910390a250505050565b5f602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690805f0160149054906101000a900463ffffffff1690805f0160189054906101000a900460ff16905083565b60015481565b5f60015f81546102aa90610797565b919050819055905060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018363ffffffff1681526020015f15158152505f5f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151815f0160146101000a81548163ffffffff021916908363ffffffff1602179055506040820151815f0160186101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff16817fb1ef6f09514dece37223c1ec7728e9c8c30e319e98425b5a2570f5123587377c846040516103ce91906107de565b60405180910390a3919050565b5f5ffd5b5f5ffd5b5f819050919050565b6103f5816103e3565b81146103ff575f5ffd5b50565b5f81359050610410816103ec565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261043757610436610416565b5b8235905067ffffffffffffffff8111156104545761045361041a565b5b6020830191508360208202830111156104705761046f61041e565b5b9250929050565b5f5f5f6040848603121561048e5761048d6103db565b5b5f61049b86828701610402565b935050602084013567ffffffffffffffff8111156104bc576104bb6103df565b5b6104c886828701610422565b92509250509250925092565b5f602082840312156104e9576104e86103db565b5b5f6104f684828501610402565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610528826104ff565b9050919050565b6105388161051e565b82525050565b5f63ffffffff82169050919050565b6105568161053e565b82525050565b5f8115159050919050565b6105708161055c565b82525050565b5f6060820190506105895f83018661052f565b610596602083018561054d565b6105a36040830184610567565b949350505050565b6105b4816103e3565b82525050565b5f6020820190506105cd5f8301846105ab565b92915050565b6105dc8161053e565b81146105e6575f5ffd5b50565b5f813590506105f7816105d3565b92915050565b5f60208284031215610612576106116103db565b5b5f61061f848285016105e9565b91505092915050565b5f82825260208201905092915050565b7f416c72656164792066756c66696c6c65640000000000000000000000000000005f82015250565b5f61066c601183610628565b915061067782610638565b602082019050919050565b5f6020820190508181035f83015261069981610660565b9050919050565b5f82825260208201905092915050565b5f5ffd5b82818337505050565b5f6106c883856106a0565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156106fb576106fa6106b0565b5b60208302925061070c8385846106b4565b82840190509392505050565b5f60408201905061072b5f8301866105ab565b818103602083015261073e8184866106bd565b9050949350505050565b5f6020820190508181035f8301526107618184866106bd565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6107a1826103e3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107d3576107d261076a565b5b600182019050919050565b5f6020820190506107f15f83018461054d565b9291505056fea2646970667358221220fd69074cdadd3a9c0cde0a32db8061c247d8d8f67d667c8aaeac9b666ed92d7464736f6c634300081c0033",
}

// VRFCoordinatorV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFCoordinatorV2MetaData.ABI instead.
var VRFCoordinatorV2ABI = VRFCoordinatorV2MetaData.ABI

// VRFCoordinatorV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFCoordinatorV2MetaData.Bin instead.
var VRFCoordinatorV2Bin = VRFCoordinatorV2MetaData.Bin

// DeployVRFCoordinatorV2 deploys a new Ethereum contract, binding an instance of VRFCoordinatorV2 to it.
func DeployVRFCoordinatorV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFCoordinatorV2, error) {
	parsed, err := VRFCoordinatorV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV2{VRFCoordinatorV2Caller: VRFCoordinatorV2Caller{contract: contract}, VRFCoordinatorV2Transactor: VRFCoordinatorV2Transactor{contract: contract}, VRFCoordinatorV2Filterer: VRFCoordinatorV2Filterer{contract: contract}}, nil
}

// VRFCoordinatorV2 is an auto generated Go binding around an Ethereum contract.
type VRFCoordinatorV2 struct {
	VRFCoordinatorV2Caller     // Read-only binding to the contract
	VRFCoordinatorV2Transactor // Write-only binding to the contract
	VRFCoordinatorV2Filterer   // Log filterer for contract events
}

// VRFCoordinatorV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCoordinatorV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFCoordinatorV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFCoordinatorV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFCoordinatorV2Session struct {
	Contract     *VRFCoordinatorV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCoordinatorV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCoordinatorV2CallerSession struct {
	Contract *VRFCoordinatorV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VRFCoordinatorV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFCoordinatorV2TransactorSession struct {
	Contract     *VRFCoordinatorV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VRFCoordinatorV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VRFCoordinatorV2Raw struct {
	Contract *VRFCoordinatorV2 // Generic contract binding to access the raw methods on
}

// VRFCoordinatorV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCoordinatorV2CallerRaw struct {
	Contract *VRFCoordinatorV2Caller // Generic read-only contract binding to access the raw methods on
}

// VRFCoordinatorV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFCoordinatorV2TransactorRaw struct {
	Contract *VRFCoordinatorV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFCoordinatorV2 creates a new instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2, error) {
	contract, err := bindVRFCoordinatorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2{VRFCoordinatorV2Caller: VRFCoordinatorV2Caller{contract: contract}, VRFCoordinatorV2Transactor: VRFCoordinatorV2Transactor{contract: contract}, VRFCoordinatorV2Filterer: VRFCoordinatorV2Filterer{contract: contract}}, nil
}

// NewVRFCoordinatorV2Caller creates a new read-only instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2Caller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2Caller, error) {
	contract, err := bindVRFCoordinatorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Caller{contract: contract}, nil
}

// NewVRFCoordinatorV2Transactor creates a new write-only instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2Transactor, error) {
	contract, err := bindVRFCoordinatorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Transactor{contract: contract}, nil
}

// NewVRFCoordinatorV2Filterer creates a new log filterer instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2Filterer, error) {
	contract, err := bindVRFCoordinatorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Filterer{contract: contract}, nil
}

// bindVRFCoordinatorV2 binds a generic wrapper to an already deployed contract.
func bindVRFCoordinatorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.contract.Transact(opts, method, params...)
}

// RequestCounter is a free data retrieval call binding the contract method 0x973a814e.
//
// Solidity: function requestCounter() view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) RequestCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "requestCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCounter is a free data retrieval call binding the contract method 0x973a814e.
//
// Solidity: function requestCounter() view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RequestCounter() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.RequestCounter(&_VRFCoordinatorV2.CallOpts)
}

// RequestCounter is a free data retrieval call binding the contract method 0x973a814e.
//
// Solidity: function requestCounter() view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) RequestCounter() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.RequestCounter(&_VRFCoordinatorV2.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(address requester, uint32 numWords, bool fulfilled)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requester common.Address
	NumWords  uint32
	Fulfilled bool
}, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Requester common.Address
		NumWords  uint32
		Fulfilled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NumWords = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Fulfilled = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(address requester, uint32 numWords, bool fulfilled)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) Requests(arg0 *big.Int) (struct {
	Requester common.Address
	NumWords  uint32
	Fulfilled bool
}, error) {
	return _VRFCoordinatorV2.Contract.Requests(&_VRFCoordinatorV2.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(address requester, uint32 numWords, bool fulfilled)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) Requests(arg0 *big.Int) (struct {
	Requester common.Address
	NumWords  uint32
	Fulfilled bool
}, error) {
	return _VRFCoordinatorV2.Contract.Requests(&_VRFCoordinatorV2.CallOpts, arg0)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) FulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "fulfillRandomWords", requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) FulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.FulfillRandomWords(&_VRFCoordinatorV2.TransactOpts, requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) FulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.FulfillRandomWords(&_VRFCoordinatorV2.TransactOpts, requestId, randomWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0xe726f2e1.
//
// Solidity: function requestRandomWords(uint32 numWords) returns(uint256 requestId)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RequestRandomWords(opts *bind.TransactOpts, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "requestRandomWords", numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0xe726f2e1.
//
// Solidity: function requestRandomWords(uint32 numWords) returns(uint256 requestId)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RequestRandomWords(numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestRandomWords(&_VRFCoordinatorV2.TransactOpts, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0xe726f2e1.
//
// Solidity: function requestRandomWords(uint32 numWords) returns(uint256 requestId)
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RequestRandomWords(numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestRandomWords(&_VRFCoordinatorV2.TransactOpts, numWords)
}

// VRFCoordinatorV2RandomWordsFulfilledIterator is returned from FilterRandomWordsFulfilled and is used to iterate over the raw logs and unpacked data for RandomWordsFulfilled events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV2RandomWordsFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VRFCoordinatorV2RandomWordsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2RandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2RandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VRFCoordinatorV2RandomWordsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2RandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2RandomWordsFulfilled represents a RandomWordsFulfilled event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RandomWordsFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsFulfilled is a free log retrieval operation binding the contract event 0xf45ee76115b0ed5f4ebe293254449fbe612bad36a53d52b87b6a40687adc48de.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256[] randomWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV2RandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2RandomWordsFulfilledIterator{contract: _VRFCoordinatorV2.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomWordsFulfilled is a free log subscription operation binding the contract event 0xf45ee76115b0ed5f4ebe293254449fbe612bad36a53d52b87b6a40687adc48de.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256[] randomWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2RandomWordsFulfilled)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordsFulfilled is a log parse operation binding the contract event 0xf45ee76115b0ed5f4ebe293254449fbe612bad36a53d52b87b6a40687adc48de.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256[] randomWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2RandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV2RandomWordsFulfilled)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFCoordinatorV2RandomWordsRequestedIterator is returned from FilterRandomWordsRequested and is used to iterate over the raw logs and unpacked data for RandomWordsRequested events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV2RandomWordsRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VRFCoordinatorV2RandomWordsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2RandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2RandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VRFCoordinatorV2RandomWordsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2RandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2RandomWordsRequested represents a RandomWordsRequested event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RandomWordsRequested struct {
	RequestId *big.Int
	Requester common.Address
	NumWords  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsRequested is a free log retrieval operation binding the contract event 0xb1ef6f09514dece37223c1ec7728e9c8c30e319e98425b5a2570f5123587377c.
//
// Solidity: event RandomWordsRequested(uint256 indexed requestId, address indexed requester, uint32 numWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterRandomWordsRequested(opts *bind.FilterOpts, requestId []*big.Int, requester []common.Address) (*VRFCoordinatorV2RandomWordsRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "RandomWordsRequested", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2RandomWordsRequestedIterator{contract: _VRFCoordinatorV2.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordsRequested is a free log subscription operation binding the contract event 0xb1ef6f09514dece37223c1ec7728e9c8c30e319e98425b5a2570f5123587377c.
//
// Solidity: event RandomWordsRequested(uint256 indexed requestId, address indexed requester, uint32 numWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RandomWordsRequested, requestId []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "RandomWordsRequested", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2RandomWordsRequested)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordsRequested is a log parse operation binding the contract event 0xb1ef6f09514dece37223c1ec7728e9c8c30e319e98425b5a2570f5123587377c.
//
// Solidity: event RandomWordsRequested(uint256 indexed requestId, address indexed requester, uint32 numWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2RandomWordsRequested, error) {
	event := new(VRFCoordinatorV2RandomWordsRequested)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
