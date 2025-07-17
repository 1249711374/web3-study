// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package visit

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

// VisitMetaData contains all meta data concerning the Visit contract.
var VisitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"visit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f5f553480156011575f5ffd5b506101608061001f5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806306661abd146100385780630cc4330c14610056575b5f5ffd5b610040610074565b60405161004d91906100b1565b60405180910390f35b61005e610079565b60405161006b91906100b1565b60405180910390f35b5f5481565b5f60015f5f82825461008b91906100f7565b925050819055505f54905090565b5f819050919050565b6100ab81610099565b82525050565b5f6020820190506100c45f8301846100a2565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61010182610099565b915061010c83610099565b9250828201905080821115610124576101236100ca565b5b9291505056fea26469706673582212203329fe086c1af12fdb3733053591760027b7b89a815f9117bdfb075ebacf6c4964736f6c634300081d0033",
}

// VisitABI is the input ABI used to generate the binding from.
// Deprecated: Use VisitMetaData.ABI instead.
var VisitABI = VisitMetaData.ABI

// VisitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VisitMetaData.Bin instead.
var VisitBin = VisitMetaData.Bin

// DeployVisit deploys a new Ethereum contract, binding an instance of Visit to it.
func DeployVisit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Visit, error) {
	parsed, err := VisitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VisitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Visit{VisitCaller: VisitCaller{contract: contract}, VisitTransactor: VisitTransactor{contract: contract}, VisitFilterer: VisitFilterer{contract: contract}}, nil
}

// Visit is an auto generated Go binding around an Ethereum contract.
type Visit struct {
	VisitCaller     // Read-only binding to the contract
	VisitTransactor // Write-only binding to the contract
	VisitFilterer   // Log filterer for contract events
}

// VisitCaller is an auto generated read-only Go binding around an Ethereum contract.
type VisitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VisitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VisitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VisitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VisitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VisitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VisitSession struct {
	Contract     *Visit            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VisitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VisitCallerSession struct {
	Contract *VisitCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VisitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VisitTransactorSession struct {
	Contract     *VisitTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VisitRaw is an auto generated low-level Go binding around an Ethereum contract.
type VisitRaw struct {
	Contract *Visit // Generic contract binding to access the raw methods on
}

// VisitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VisitCallerRaw struct {
	Contract *VisitCaller // Generic read-only contract binding to access the raw methods on
}

// VisitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VisitTransactorRaw struct {
	Contract *VisitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVisit creates a new instance of Visit, bound to a specific deployed contract.
func NewVisit(address common.Address, backend bind.ContractBackend) (*Visit, error) {
	contract, err := bindVisit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Visit{VisitCaller: VisitCaller{contract: contract}, VisitTransactor: VisitTransactor{contract: contract}, VisitFilterer: VisitFilterer{contract: contract}}, nil
}

// NewVisitCaller creates a new read-only instance of Visit, bound to a specific deployed contract.
func NewVisitCaller(address common.Address, caller bind.ContractCaller) (*VisitCaller, error) {
	contract, err := bindVisit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VisitCaller{contract: contract}, nil
}

// NewVisitTransactor creates a new write-only instance of Visit, bound to a specific deployed contract.
func NewVisitTransactor(address common.Address, transactor bind.ContractTransactor) (*VisitTransactor, error) {
	contract, err := bindVisit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VisitTransactor{contract: contract}, nil
}

// NewVisitFilterer creates a new log filterer instance of Visit, bound to a specific deployed contract.
func NewVisitFilterer(address common.Address, filterer bind.ContractFilterer) (*VisitFilterer, error) {
	contract, err := bindVisit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VisitFilterer{contract: contract}, nil
}

// bindVisit binds a generic wrapper to an already deployed contract.
func bindVisit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VisitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Visit *VisitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Visit.Contract.VisitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Visit *VisitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Visit.Contract.VisitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Visit *VisitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Visit.Contract.VisitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Visit *VisitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Visit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Visit *VisitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Visit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Visit *VisitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Visit.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Visit *VisitCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Visit.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Visit *VisitSession) Count() (*big.Int, error) {
	return _Visit.Contract.Count(&_Visit.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Visit *VisitCallerSession) Count() (*big.Int, error) {
	return _Visit.Contract.Count(&_Visit.CallOpts)
}

// Visit is a paid mutator transaction binding the contract method 0x0cc4330c.
//
// Solidity: function visit() returns(uint256)
func (_Visit *VisitTransactor) Visit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Visit.contract.Transact(opts, "visit")
}

// Visit is a paid mutator transaction binding the contract method 0x0cc4330c.
//
// Solidity: function visit() returns(uint256)
func (_Visit *VisitSession) Visit() (*types.Transaction, error) {
	return _Visit.Contract.Visit(&_Visit.TransactOpts)
}

// Visit is a paid mutator transaction binding the contract method 0x0cc4330c.
//
// Solidity: function visit() returns(uint256)
func (_Visit *VisitTransactorSession) Visit() (*types.Transaction, error) {
	return _Visit.Contract.Visit(&_Visit.TransactOpts)
}
