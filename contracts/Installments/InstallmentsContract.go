// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Installments

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	//_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// InstallmentsABI is the input ABI used to generate the binding from.
const InstallmentsABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_descriptor\",\"type\":\"address\"}],\"name\":\"setDescriptor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_engine\",\"type\":\"address\"}],\"name\":\"setEngine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"simFirstObligation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"descriptor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modelId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"simTotalObligation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"}],\"name\":\"fixClock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFinalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"simPunitiveInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"punitiveInterestRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getInstallments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"real\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getClosingObligation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getObligation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_PAID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_cuota\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"_installments\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"_duration\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"_timeUnit\",\"type\":\"uint32\"}],\"name\":\"encodeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_ONGOING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"L_DATA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"engine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getDueTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"configs\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"installments\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"timeUnit\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"duration\",\"type\":\"uint40\"},{\"internalType\":\"uint64\",\"name\":\"lentTime\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"cuota\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"simInstallments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"installments\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFrequency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"simDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"simFrequency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"frequency\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getEstimateObligation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"run\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"states\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"clock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"paid\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"paidBase\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"interest\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_engine\",\"type\":\"address\"}],\"name\":\"_setEngine\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_descriptor\",\"type\":\"address\"}],\"name\":\"_setDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_to\",\"type\":\"uint64\"}],\"name\":\"_setClock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_paidBase\",\"type\":\"uint128\"}],\"name\":\"_setPaidBase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_interest\",\"type\":\"uint128\"}],\"name\":\"_setInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"ChangedStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_debt\",\"type\":\"uint256\"}],\"name\":\"ChangedObligation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frequency\",\"type\":\"uint256\"}],\"name\":\"ChangedFrequency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"ChangedDueTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_dueTime\",\"type\":\"uint64\"}],\"name\":\"ChangedFinalTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AddedDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_paid\",\"type\":\"uint256\"}],\"name\":\"AddedPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Installments is an auto generated Go binding around an Ethereum contract.
type Installments struct {
	InstallmentsCaller     // Read-only binding to the contract
	InstallmentsTransactor // Write-only binding to the contract
	InstallmentsFilterer   // Log filterer for contract events
}

// InstallmentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstallmentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstallmentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstallmentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstallmentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstallmentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstallmentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstallmentsSession struct {
	Contract     *Installments     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InstallmentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstallmentsCallerSession struct {
	Contract *InstallmentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InstallmentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstallmentsTransactorSession struct {
	Contract     *InstallmentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InstallmentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstallmentsRaw struct {
	Contract *Installments // Generic contract binding to access the raw methods on
}

// InstallmentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstallmentsCallerRaw struct {
	Contract *InstallmentsCaller // Generic read-only contract binding to access the raw methods on
}

// InstallmentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstallmentsTransactorRaw struct {
	Contract *InstallmentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstallments creates a new instance of Installments, bound to a specific deployed contract.
func NewInstallments(address common.Address, backend bind.ContractBackend) (*Installments, error) {
	contract, err := bindInstallments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Installments{InstallmentsCaller: InstallmentsCaller{contract: contract}, InstallmentsTransactor: InstallmentsTransactor{contract: contract}, InstallmentsFilterer: InstallmentsFilterer{contract: contract}}, nil
}

// NewInstallmentsCaller creates a new read-only instance of Installments, bound to a specific deployed contract.
func NewInstallmentsCaller(address common.Address, caller bind.ContractCaller) (*InstallmentsCaller, error) {
	contract, err := bindInstallments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstallmentsCaller{contract: contract}, nil
}

// NewInstallmentsTransactor creates a new write-only instance of Installments, bound to a specific deployed contract.
func NewInstallmentsTransactor(address common.Address, transactor bind.ContractTransactor) (*InstallmentsTransactor, error) {
	contract, err := bindInstallments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstallmentsTransactor{contract: contract}, nil
}

// NewInstallmentsFilterer creates a new log filterer instance of Installments, bound to a specific deployed contract.
func NewInstallmentsFilterer(address common.Address, filterer bind.ContractFilterer) (*InstallmentsFilterer, error) {
	contract, err := bindInstallments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstallmentsFilterer{contract: contract}, nil
}

// bindInstallments binds a generic wrapper to an already deployed contract.
func bindInstallments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InstallmentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Installments *InstallmentsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Installments.Contract.InstallmentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Installments *InstallmentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Installments.Contract.InstallmentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Installments *InstallmentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Installments.Contract.InstallmentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Installments *InstallmentsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Installments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Installments *InstallmentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Installments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Installments *InstallmentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Installments.Contract.contract.Transact(opts, method, params...)
}

// LDATA is a free data retrieval call binding the contract method 0xc61aff5a.
//
// Solidity: function L_DATA() constant returns(uint256)
func (_Installments *InstallmentsCaller) LDATA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "L_DATA")
	return *ret0, err
}

// LDATA is a free data retrieval call binding the contract method 0xc61aff5a.
//
// Solidity: function L_DATA() constant returns(uint256)
func (_Installments *InstallmentsSession) LDATA() (*big.Int, error) {
	return _Installments.Contract.LDATA(&_Installments.CallOpts)
}

// LDATA is a free data retrieval call binding the contract method 0xc61aff5a.
//
// Solidity: function L_DATA() constant returns(uint256)
func (_Installments *InstallmentsCallerSession) LDATA() (*big.Int, error) {
	return _Installments.Contract.LDATA(&_Installments.CallOpts)
}

// STATUSERROR is a free data retrieval call binding the contract method 0xc1a5b260.
//
// Solidity: function STATUS_ERROR() constant returns(uint256)
func (_Installments *InstallmentsCaller) STATUSERROR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "STATUS_ERROR")
	return *ret0, err
}

// STATUSERROR is a free data retrieval call binding the contract method 0xc1a5b260.
//
// Solidity: function STATUS_ERROR() constant returns(uint256)
func (_Installments *InstallmentsSession) STATUSERROR() (*big.Int, error) {
	return _Installments.Contract.STATUSERROR(&_Installments.CallOpts)
}

// STATUSERROR is a free data retrieval call binding the contract method 0xc1a5b260.
//
// Solidity: function STATUS_ERROR() constant returns(uint256)
func (_Installments *InstallmentsCallerSession) STATUSERROR() (*big.Int, error) {
	return _Installments.Contract.STATUSERROR(&_Installments.CallOpts)
}

// STATUSONGOING is a free data retrieval call binding the contract method 0xc075947d.
//
// Solidity: function STATUS_ONGOING() constant returns(uint256)
func (_Installments *InstallmentsCaller) STATUSONGOING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "STATUS_ONGOING")
	return *ret0, err
}

// STATUSONGOING is a free data retrieval call binding the contract method 0xc075947d.
//
// Solidity: function STATUS_ONGOING() constant returns(uint256)
func (_Installments *InstallmentsSession) STATUSONGOING() (*big.Int, error) {
	return _Installments.Contract.STATUSONGOING(&_Installments.CallOpts)
}

// STATUSONGOING is a free data retrieval call binding the contract method 0xc075947d.
//
// Solidity: function STATUS_ONGOING() constant returns(uint256)
func (_Installments *InstallmentsCallerSession) STATUSONGOING() (*big.Int, error) {
	return _Installments.Contract.STATUSONGOING(&_Installments.CallOpts)
}

// STATUSPAID is a free data retrieval call binding the contract method 0xb317d031.
//
// Solidity: function STATUS_PAID() constant returns(uint256)
func (_Installments *InstallmentsCaller) STATUSPAID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "STATUS_PAID")
	return *ret0, err
}

// STATUSPAID is a free data retrieval call binding the contract method 0xb317d031.
//
// Solidity: function STATUS_PAID() constant returns(uint256)
func (_Installments *InstallmentsSession) STATUSPAID() (*big.Int, error) {
	return _Installments.Contract.STATUSPAID(&_Installments.CallOpts)
}

// STATUSPAID is a free data retrieval call binding the contract method 0xb317d031.
//
// Solidity: function STATUS_PAID() constant returns(uint256)
func (_Installments *InstallmentsCallerSession) STATUSPAID() (*big.Int, error) {
	return _Installments.Contract.STATUSPAID(&_Installments.CallOpts)
}

// Configs is a free data retrieval call binding the contract method 0xce8f6078.
//
// Solidity: function configs(bytes32 ) constant returns(uint24 installments, uint32 timeUnit, uint40 duration, uint64 lentTime, uint128 cuota, uint256 interestRate)
func (_Installments *InstallmentsCaller) Configs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Installments *big.Int
	TimeUnit     uint32
	Duration     *big.Int
	LentTime     uint64
	Cuota        *big.Int
	InterestRate *big.Int
}, error) {
	ret := new(struct {
		Installments *big.Int
		TimeUnit     uint32
		Duration     *big.Int
		LentTime     uint64
		Cuota        *big.Int
		InterestRate *big.Int
	})
	out := ret
	err := _Installments.contract.Call(opts, out, "configs", arg0)
	return *ret, err
}

// Configs is a free data retrieval call binding the contract method 0xce8f6078.
//
// Solidity: function configs(bytes32 ) constant returns(uint24 installments, uint32 timeUnit, uint40 duration, uint64 lentTime, uint128 cuota, uint256 interestRate)
func (_Installments *InstallmentsSession) Configs(arg0 [32]byte) (struct {
	Installments *big.Int
	TimeUnit     uint32
	Duration     *big.Int
	LentTime     uint64
	Cuota        *big.Int
	InterestRate *big.Int
}, error) {
	return _Installments.Contract.Configs(&_Installments.CallOpts, arg0)
}

// Configs is a free data retrieval call binding the contract method 0xce8f6078.
//
// Solidity: function configs(bytes32 ) constant returns(uint24 installments, uint32 timeUnit, uint40 duration, uint64 lentTime, uint128 cuota, uint256 interestRate)
func (_Installments *InstallmentsCallerSession) Configs(arg0 [32]byte) (struct {
	Installments *big.Int
	TimeUnit     uint32
	Duration     *big.Int
	LentTime     uint64
	Cuota        *big.Int
	InterestRate *big.Int
}, error) {
	return _Installments.Contract.Configs(&_Installments.CallOpts, arg0)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() constant returns(address)
func (_Installments *InstallmentsCaller) Descriptor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "descriptor")
	return *ret0, err
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() constant returns(address)
func (_Installments *InstallmentsSession) Descriptor() (common.Address, error) {
	return _Installments.Contract.Descriptor(&_Installments.CallOpts)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() constant returns(address)
func (_Installments *InstallmentsCallerSession) Descriptor() (common.Address, error) {
	return _Installments.Contract.Descriptor(&_Installments.CallOpts)
}

// EncodeData is a free data retrieval call binding the contract method 0xbd623d68.
//
// Solidity: function encodeData(uint128 _cuota, uint256 _interestRate, uint24 _installments, uint40 _duration, uint32 _timeUnit) constant returns(bytes)
func (_Installments *InstallmentsCaller) EncodeData(opts *bind.CallOpts, _cuota *big.Int, _interestRate *big.Int, _installments *big.Int, _duration *big.Int, _timeUnit uint32) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "encodeData", _cuota, _interestRate, _installments, _duration, _timeUnit)
	return *ret0, err
}

// EncodeData is a free data retrieval call binding the contract method 0xbd623d68.
//
// Solidity: function encodeData(uint128 _cuota, uint256 _interestRate, uint24 _installments, uint40 _duration, uint32 _timeUnit) constant returns(bytes)
func (_Installments *InstallmentsSession) EncodeData(_cuota *big.Int, _interestRate *big.Int, _installments *big.Int, _duration *big.Int, _timeUnit uint32) ([]byte, error) {
	return _Installments.Contract.EncodeData(&_Installments.CallOpts, _cuota, _interestRate, _installments, _duration, _timeUnit)
}

// EncodeData is a free data retrieval call binding the contract method 0xbd623d68.
//
// Solidity: function encodeData(uint128 _cuota, uint256 _interestRate, uint24 _installments, uint40 _duration, uint32 _timeUnit) constant returns(bytes)
func (_Installments *InstallmentsCallerSession) EncodeData(_cuota *big.Int, _interestRate *big.Int, _installments *big.Int, _duration *big.Int, _timeUnit uint32) ([]byte, error) {
	return _Installments.Contract.EncodeData(&_Installments.CallOpts, _cuota, _interestRate, _installments, _duration, _timeUnit)
}

// Engine is a free data retrieval call binding the contract method 0xc9d4623f.
//
// Solidity: function engine() constant returns(address)
func (_Installments *InstallmentsCaller) Engine(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "engine")
	return *ret0, err
}

// Engine is a free data retrieval call binding the contract method 0xc9d4623f.
//
// Solidity: function engine() constant returns(address)
func (_Installments *InstallmentsSession) Engine() (common.Address, error) {
	return _Installments.Contract.Engine(&_Installments.CallOpts)
}

// Engine is a free data retrieval call binding the contract method 0xc9d4623f.
//
// Solidity: function engine() constant returns(address)
func (_Installments *InstallmentsCallerSession) Engine() (common.Address, error) {
	return _Installments.Contract.Engine(&_Installments.CallOpts)
}

// GetClosingObligation is a free data retrieval call binding the contract method 0x9a6203e9.
//
// Solidity: function getClosingObligation(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetClosingObligation(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getClosingObligation", id)
	return *ret0, err
}

// GetClosingObligation is a free data retrieval call binding the contract method 0x9a6203e9.
//
// Solidity: function getClosingObligation(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetClosingObligation(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetClosingObligation(&_Installments.CallOpts, id)
}

// GetClosingObligation is a free data retrieval call binding the contract method 0x9a6203e9.
//
// Solidity: function getClosingObligation(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetClosingObligation(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetClosingObligation(&_Installments.CallOpts, id)
}

// GetDueTime is a free data retrieval call binding the contract method 0xcdd8750e.
//
// Solidity: function getDueTime(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetDueTime(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getDueTime", id)
	return *ret0, err
}

// GetDueTime is a free data retrieval call binding the contract method 0xcdd8750e.
//
// Solidity: function getDueTime(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetDueTime(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetDueTime(&_Installments.CallOpts, id)
}

// GetDueTime is a free data retrieval call binding the contract method 0xcdd8750e.
//
// Solidity: function getDueTime(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetDueTime(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetDueTime(&_Installments.CallOpts, id)
}

// GetEstimateObligation is a free data retrieval call binding the contract method 0xeda71f6e.
//
// Solidity: function getEstimateObligation(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetEstimateObligation(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getEstimateObligation", id)
	return *ret0, err
}

// GetEstimateObligation is a free data retrieval call binding the contract method 0xeda71f6e.
//
// Solidity: function getEstimateObligation(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetEstimateObligation(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetEstimateObligation(&_Installments.CallOpts, id)
}

// GetEstimateObligation is a free data retrieval call binding the contract method 0xeda71f6e.
//
// Solidity: function getEstimateObligation(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetEstimateObligation(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetEstimateObligation(&_Installments.CallOpts, id)
}

// GetFinalTime is a free data retrieval call binding the contract method 0x73549604.
//
// Solidity: function getFinalTime(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetFinalTime(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getFinalTime", id)
	return *ret0, err
}

// GetFinalTime is a free data retrieval call binding the contract method 0x73549604.
//
// Solidity: function getFinalTime(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetFinalTime(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetFinalTime(&_Installments.CallOpts, id)
}

// GetFinalTime is a free data retrieval call binding the contract method 0x73549604.
//
// Solidity: function getFinalTime(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetFinalTime(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetFinalTime(&_Installments.CallOpts, id)
}

// GetFrequency is a free data retrieval call binding the contract method 0xd0fc1e7d.
//
// Solidity: function getFrequency(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetFrequency(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getFrequency", id)
	return *ret0, err
}

// GetFrequency is a free data retrieval call binding the contract method 0xd0fc1e7d.
//
// Solidity: function getFrequency(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetFrequency(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetFrequency(&_Installments.CallOpts, id)
}

// GetFrequency is a free data retrieval call binding the contract method 0xd0fc1e7d.
//
// Solidity: function getFrequency(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetFrequency(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetFrequency(&_Installments.CallOpts, id)
}

// GetInstallments is a free data retrieval call binding the contract method 0x87195d23.
//
// Solidity: function getInstallments(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetInstallments(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getInstallments", id)
	return *ret0, err
}

// GetInstallments is a free data retrieval call binding the contract method 0x87195d23.
//
// Solidity: function getInstallments(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetInstallments(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetInstallments(&_Installments.CallOpts, id)
}

// GetInstallments is a free data retrieval call binding the contract method 0x87195d23.
//
// Solidity: function getInstallments(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetInstallments(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetInstallments(&_Installments.CallOpts, id)
}

// GetObligation is a free data retrieval call binding the contract method 0xa6f6d8bb.
//
// Solidity: function getObligation(bytes32 id, uint64 timestamp) constant returns(uint256, bool)
func (_Installments *InstallmentsCaller) GetObligation(opts *bind.CallOpts, id [32]byte, timestamp uint64) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Installments.contract.Call(opts, out, "getObligation", id, timestamp)
	return *ret0, *ret1, err
}

// GetObligation is a free data retrieval call binding the contract method 0xa6f6d8bb.
//
// Solidity: function getObligation(bytes32 id, uint64 timestamp) constant returns(uint256, bool)
func (_Installments *InstallmentsSession) GetObligation(id [32]byte, timestamp uint64) (*big.Int, bool, error) {
	return _Installments.Contract.GetObligation(&_Installments.CallOpts, id, timestamp)
}

// GetObligation is a free data retrieval call binding the contract method 0xa6f6d8bb.
//
// Solidity: function getObligation(bytes32 id, uint64 timestamp) constant returns(uint256, bool)
func (_Installments *InstallmentsCallerSession) GetObligation(id [32]byte, timestamp uint64) (*big.Int, bool, error) {
	return _Installments.Contract.GetObligation(&_Installments.CallOpts, id, timestamp)
}

// GetPaid is a free data retrieval call binding the contract method 0xa02b1a51.
//
// Solidity: function getPaid(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetPaid(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getPaid", id)
	return *ret0, err
}

// GetPaid is a free data retrieval call binding the contract method 0xa02b1a51.
//
// Solidity: function getPaid(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetPaid(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetPaid(&_Installments.CallOpts, id)
}

// GetPaid is a free data retrieval call binding the contract method 0xa02b1a51.
//
// Solidity: function getPaid(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetPaid(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetPaid(&_Installments.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCaller) GetStatus(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "getStatus", id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsSession) GetStatus(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetStatus(&_Installments.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 id) constant returns(uint256)
func (_Installments *InstallmentsCallerSession) GetStatus(id [32]byte) (*big.Int, error) {
	return _Installments.Contract.GetStatus(&_Installments.CallOpts, id)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _target) constant returns(bool)
func (_Installments *InstallmentsCaller) IsOperator(opts *bind.CallOpts, _target common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "isOperator", _target)
	return *ret0, err
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _target) constant returns(bool)
func (_Installments *InstallmentsSession) IsOperator(_target common.Address) (bool, error) {
	return _Installments.Contract.IsOperator(&_Installments.CallOpts, _target)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _target) constant returns(bool)
func (_Installments *InstallmentsCallerSession) IsOperator(_target common.Address) (bool, error) {
	return _Installments.Contract.IsOperator(&_Installments.CallOpts, _target)
}

// ModelId is a free data retrieval call binding the contract method 0x43cdaad2.
//
// Solidity: function modelId() constant returns(bytes32)
func (_Installments *InstallmentsCaller) ModelId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "modelId")
	return *ret0, err
}

// ModelId is a free data retrieval call binding the contract method 0x43cdaad2.
//
// Solidity: function modelId() constant returns(bytes32)
func (_Installments *InstallmentsSession) ModelId() ([32]byte, error) {
	return _Installments.Contract.ModelId(&_Installments.CallOpts)
}

// ModelId is a free data retrieval call binding the contract method 0x43cdaad2.
//
// Solidity: function modelId() constant returns(bytes32)
func (_Installments *InstallmentsCallerSession) ModelId() ([32]byte, error) {
	return _Installments.Contract.ModelId(&_Installments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Installments *InstallmentsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Installments *InstallmentsSession) Owner() (common.Address, error) {
	return _Installments.Contract.Owner(&_Installments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Installments *InstallmentsCallerSession) Owner() (common.Address, error) {
	return _Installments.Contract.Owner(&_Installments.CallOpts)
}

// SimDuration is a free data retrieval call binding the contract method 0xe54d62e9.
//
// Solidity: function simDuration(bytes _data) constant returns(uint256 duration)
func (_Installments *InstallmentsCaller) SimDuration(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "simDuration", _data)
	return *ret0, err
}

// SimDuration is a free data retrieval call binding the contract method 0xe54d62e9.
//
// Solidity: function simDuration(bytes _data) constant returns(uint256 duration)
func (_Installments *InstallmentsSession) SimDuration(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimDuration(&_Installments.CallOpts, _data)
}

// SimDuration is a free data retrieval call binding the contract method 0xe54d62e9.
//
// Solidity: function simDuration(bytes _data) constant returns(uint256 duration)
func (_Installments *InstallmentsCallerSession) SimDuration(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimDuration(&_Installments.CallOpts, _data)
}

// SimFirstObligation is a free data retrieval call binding the contract method 0x16d63f17.
//
// Solidity: function simFirstObligation(bytes _data) constant returns(uint256 amount, uint256 time)
func (_Installments *InstallmentsCaller) SimFirstObligation(opts *bind.CallOpts, _data []byte) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	ret := new(struct {
		Amount *big.Int
		Time   *big.Int
	})
	out := ret
	err := _Installments.contract.Call(opts, out, "simFirstObligation", _data)
	return *ret, err
}

// SimFirstObligation is a free data retrieval call binding the contract method 0x16d63f17.
//
// Solidity: function simFirstObligation(bytes _data) constant returns(uint256 amount, uint256 time)
func (_Installments *InstallmentsSession) SimFirstObligation(_data []byte) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _Installments.Contract.SimFirstObligation(&_Installments.CallOpts, _data)
}

// SimFirstObligation is a free data retrieval call binding the contract method 0x16d63f17.
//
// Solidity: function simFirstObligation(bytes _data) constant returns(uint256 amount, uint256 time)
func (_Installments *InstallmentsCallerSession) SimFirstObligation(_data []byte) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _Installments.Contract.SimFirstObligation(&_Installments.CallOpts, _data)
}

// SimFrequency is a free data retrieval call binding the contract method 0xebe955fe.
//
// Solidity: function simFrequency(bytes _data) constant returns(uint256 frequency)
func (_Installments *InstallmentsCaller) SimFrequency(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "simFrequency", _data)
	return *ret0, err
}

// SimFrequency is a free data retrieval call binding the contract method 0xebe955fe.
//
// Solidity: function simFrequency(bytes _data) constant returns(uint256 frequency)
func (_Installments *InstallmentsSession) SimFrequency(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimFrequency(&_Installments.CallOpts, _data)
}

// SimFrequency is a free data retrieval call binding the contract method 0xebe955fe.
//
// Solidity: function simFrequency(bytes _data) constant returns(uint256 frequency)
func (_Installments *InstallmentsCallerSession) SimFrequency(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimFrequency(&_Installments.CallOpts, _data)
}

// SimInstallments is a free data retrieval call binding the contract method 0xcf76cb7d.
//
// Solidity: function simInstallments(bytes _data) constant returns(uint256 installments)
func (_Installments *InstallmentsCaller) SimInstallments(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "simInstallments", _data)
	return *ret0, err
}

// SimInstallments is a free data retrieval call binding the contract method 0xcf76cb7d.
//
// Solidity: function simInstallments(bytes _data) constant returns(uint256 installments)
func (_Installments *InstallmentsSession) SimInstallments(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimInstallments(&_Installments.CallOpts, _data)
}

// SimInstallments is a free data retrieval call binding the contract method 0xcf76cb7d.
//
// Solidity: function simInstallments(bytes _data) constant returns(uint256 installments)
func (_Installments *InstallmentsCallerSession) SimInstallments(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimInstallments(&_Installments.CallOpts, _data)
}

// SimPunitiveInterestRate is a free data retrieval call binding the contract method 0x814eba9e.
//
// Solidity: function simPunitiveInterestRate(bytes _data) constant returns(uint256 punitiveInterestRate)
func (_Installments *InstallmentsCaller) SimPunitiveInterestRate(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "simPunitiveInterestRate", _data)
	return *ret0, err
}

// SimPunitiveInterestRate is a free data retrieval call binding the contract method 0x814eba9e.
//
// Solidity: function simPunitiveInterestRate(bytes _data) constant returns(uint256 punitiveInterestRate)
func (_Installments *InstallmentsSession) SimPunitiveInterestRate(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimPunitiveInterestRate(&_Installments.CallOpts, _data)
}

// SimPunitiveInterestRate is a free data retrieval call binding the contract method 0x814eba9e.
//
// Solidity: function simPunitiveInterestRate(bytes _data) constant returns(uint256 punitiveInterestRate)
func (_Installments *InstallmentsCallerSession) SimPunitiveInterestRate(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimPunitiveInterestRate(&_Installments.CallOpts, _data)
}

// SimTotalObligation is a free data retrieval call binding the contract method 0x54392a96.
//
// Solidity: function simTotalObligation(bytes _data) constant returns(uint256 amount)
func (_Installments *InstallmentsCaller) SimTotalObligation(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "simTotalObligation", _data)
	return *ret0, err
}

// SimTotalObligation is a free data retrieval call binding the contract method 0x54392a96.
//
// Solidity: function simTotalObligation(bytes _data) constant returns(uint256 amount)
func (_Installments *InstallmentsSession) SimTotalObligation(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimTotalObligation(&_Installments.CallOpts, _data)
}

// SimTotalObligation is a free data retrieval call binding the contract method 0x54392a96.
//
// Solidity: function simTotalObligation(bytes _data) constant returns(uint256 amount)
func (_Installments *InstallmentsCallerSession) SimTotalObligation(_data []byte) (*big.Int, error) {
	return _Installments.Contract.SimTotalObligation(&_Installments.CallOpts, _data)
}

// States is a free data retrieval call binding the contract method 0xfbdc1ef1.
//
// Solidity: function states(bytes32 ) constant returns(uint8 status, uint64 clock, uint64 lastPayment, uint128 paid, uint128 paidBase, uint128 interest)
func (_Installments *InstallmentsCaller) States(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status      uint8
	Clock       uint64
	LastPayment uint64
	Paid        *big.Int
	PaidBase    *big.Int
	Interest    *big.Int
}, error) {
	ret := new(struct {
		Status      uint8
		Clock       uint64
		LastPayment uint64
		Paid        *big.Int
		PaidBase    *big.Int
		Interest    *big.Int
	})
	out := ret
	err := _Installments.contract.Call(opts, out, "states", arg0)
	return *ret, err
}

// States is a free data retrieval call binding the contract method 0xfbdc1ef1.
//
// Solidity: function states(bytes32 ) constant returns(uint8 status, uint64 clock, uint64 lastPayment, uint128 paid, uint128 paidBase, uint128 interest)
func (_Installments *InstallmentsSession) States(arg0 [32]byte) (struct {
	Status      uint8
	Clock       uint64
	LastPayment uint64
	Paid        *big.Int
	PaidBase    *big.Int
	Interest    *big.Int
}, error) {
	return _Installments.Contract.States(&_Installments.CallOpts, arg0)
}

// States is a free data retrieval call binding the contract method 0xfbdc1ef1.
//
// Solidity: function states(bytes32 ) constant returns(uint8 status, uint64 clock, uint64 lastPayment, uint128 paid, uint128 paidBase, uint128 interest)
func (_Installments *InstallmentsCallerSession) States(arg0 [32]byte) (struct {
	Status      uint8
	Clock       uint64
	LastPayment uint64
	Paid        *big.Int
	PaidBase    *big.Int
	Interest    *big.Int
}, error) {
	return _Installments.Contract.States(&_Installments.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Installments *InstallmentsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Installments *InstallmentsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Installments.Contract.SupportsInterface(&_Installments.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Installments *InstallmentsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Installments.Contract.SupportsInterface(&_Installments.CallOpts, interfaceId)
}

// Validate is a free data retrieval call binding the contract method 0xc16e50ef.
//
// Solidity: function validate(bytes data) constant returns(bool)
func (_Installments *InstallmentsCaller) Validate(opts *bind.CallOpts, data []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Installments.contract.Call(opts, out, "validate", data)
	return *ret0, err
}

// Validate is a free data retrieval call binding the contract method 0xc16e50ef.
//
// Solidity: function validate(bytes data) constant returns(bool)
func (_Installments *InstallmentsSession) Validate(data []byte) (bool, error) {
	return _Installments.Contract.Validate(&_Installments.CallOpts, data)
}

// Validate is a free data retrieval call binding the contract method 0xc16e50ef.
//
// Solidity: function validate(bytes data) constant returns(bool)
func (_Installments *InstallmentsCallerSession) Validate(data []byte) (bool, error) {
	return _Installments.Contract.Validate(&_Installments.CallOpts, data)
}

// AddDebt is a paid mutator transaction binding the contract method 0x90900df4.
//
// Solidity: function addDebt(bytes32 id, uint256 amount) returns(bool)
func (_Installments *InstallmentsTransactor) AddDebt(opts *bind.TransactOpts, id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "addDebt", id, amount)
}

// AddDebt is a paid mutator transaction binding the contract method 0x90900df4.
//
// Solidity: function addDebt(bytes32 id, uint256 amount) returns(bool)
func (_Installments *InstallmentsSession) AddDebt(id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Installments.Contract.AddDebt(&_Installments.TransactOpts, id, amount)
}

// AddDebt is a paid mutator transaction binding the contract method 0x90900df4.
//
// Solidity: function addDebt(bytes32 id, uint256 amount) returns(bool)
func (_Installments *InstallmentsTransactorSession) AddDebt(id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Installments.Contract.AddDebt(&_Installments.TransactOpts, id, amount)
}

// AddPaid is a paid mutator transaction binding the contract method 0x91fa2df4.
//
// Solidity: function addPaid(bytes32 id, uint256 amount) returns(uint256 real)
func (_Installments *InstallmentsTransactor) AddPaid(opts *bind.TransactOpts, id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "addPaid", id, amount)
}

// AddPaid is a paid mutator transaction binding the contract method 0x91fa2df4.
//
// Solidity: function addPaid(bytes32 id, uint256 amount) returns(uint256 real)
func (_Installments *InstallmentsSession) AddPaid(id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Installments.Contract.AddPaid(&_Installments.TransactOpts, id, amount)
}

// AddPaid is a paid mutator transaction binding the contract method 0x91fa2df4.
//
// Solidity: function addPaid(bytes32 id, uint256 amount) returns(uint256 real)
func (_Installments *InstallmentsTransactorSession) AddPaid(id [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Installments.Contract.AddPaid(&_Installments.TransactOpts, id, amount)
}

// Create is a paid mutator transaction binding the contract method 0x5b37e150.
//
// Solidity: function create(bytes32 id, bytes data) returns(bool)
func (_Installments *InstallmentsTransactor) Create(opts *bind.TransactOpts, id [32]byte, data []byte) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "create", id, data)
}

// Create is a paid mutator transaction binding the contract method 0x5b37e150.
//
// Solidity: function create(bytes32 id, bytes data) returns(bool)
func (_Installments *InstallmentsSession) Create(id [32]byte, data []byte) (*types.Transaction, error) {
	return _Installments.Contract.Create(&_Installments.TransactOpts, id, data)
}

// Create is a paid mutator transaction binding the contract method 0x5b37e150.
//
// Solidity: function create(bytes32 id, bytes data) returns(bool)
func (_Installments *InstallmentsTransactorSession) Create(id [32]byte, data []byte) (*types.Transaction, error) {
	return _Installments.Contract.Create(&_Installments.TransactOpts, id, data)
}

// FixClock is a paid mutator transaction binding the contract method 0x625a6726.
//
// Solidity: function fixClock(bytes32 id, uint64 target) returns(bool)
func (_Installments *InstallmentsTransactor) FixClock(opts *bind.TransactOpts, id [32]byte, target uint64) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "fixClock", id, target)
}

// FixClock is a paid mutator transaction binding the contract method 0x625a6726.
//
// Solidity: function fixClock(bytes32 id, uint64 target) returns(bool)
func (_Installments *InstallmentsSession) FixClock(id [32]byte, target uint64) (*types.Transaction, error) {
	return _Installments.Contract.FixClock(&_Installments.TransactOpts, id, target)
}

// FixClock is a paid mutator transaction binding the contract method 0x625a6726.
//
// Solidity: function fixClock(bytes32 id, uint64 target) returns(bool)
func (_Installments *InstallmentsTransactorSession) FixClock(id [32]byte, target uint64) (*types.Transaction, error) {
	return _Installments.Contract.FixClock(&_Installments.TransactOpts, id, target)
}

// Run is a paid mutator transaction binding the contract method 0xef6ac0f0.
//
// Solidity: function run(bytes32 id) returns(bool)
func (_Installments *InstallmentsTransactor) Run(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "run", id)
}

// Run is a paid mutator transaction binding the contract method 0xef6ac0f0.
//
// Solidity: function run(bytes32 id) returns(bool)
func (_Installments *InstallmentsSession) Run(id [32]byte) (*types.Transaction, error) {
	return _Installments.Contract.Run(&_Installments.TransactOpts, id)
}

// Run is a paid mutator transaction binding the contract method 0xef6ac0f0.
//
// Solidity: function run(bytes32 id) returns(bool)
func (_Installments *InstallmentsTransactorSession) Run(id [32]byte) (*types.Transaction, error) {
	return _Installments.Contract.Run(&_Installments.TransactOpts, id)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns(bool)
func (_Installments *InstallmentsTransactor) SetDescriptor(opts *bind.TransactOpts, _descriptor common.Address) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "setDescriptor", _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns(bool)
func (_Installments *InstallmentsSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _Installments.Contract.SetDescriptor(&_Installments.TransactOpts, _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns(bool)
func (_Installments *InstallmentsTransactorSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _Installments.Contract.SetDescriptor(&_Installments.TransactOpts, _descriptor)
}

// SetEngine is a paid mutator transaction binding the contract method 0x0e830e49.
//
// Solidity: function setEngine(address _engine) returns(bool)
func (_Installments *InstallmentsTransactor) SetEngine(opts *bind.TransactOpts, _engine common.Address) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "setEngine", _engine)
}

// SetEngine is a paid mutator transaction binding the contract method 0x0e830e49.
//
// Solidity: function setEngine(address _engine) returns(bool)
func (_Installments *InstallmentsSession) SetEngine(_engine common.Address) (*types.Transaction, error) {
	return _Installments.Contract.SetEngine(&_Installments.TransactOpts, _engine)
}

// SetEngine is a paid mutator transaction binding the contract method 0x0e830e49.
//
// Solidity: function setEngine(address _engine) returns(bool)
func (_Installments *InstallmentsTransactorSession) SetEngine(_engine common.Address) (*types.Transaction, error) {
	return _Installments.Contract.SetEngine(&_Installments.TransactOpts, _engine)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Installments *InstallmentsTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Installments.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Installments *InstallmentsSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Installments.Contract.TransferOwnership(&_Installments.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Installments *InstallmentsTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Installments.Contract.TransferOwnership(&_Installments.TransactOpts, _newOwner)
}

// InstallmentsAddedDebtIterator is returned from FilterAddedDebt and is used to iterate over the raw logs and unpacked data for AddedDebt events raised by the Installments contract.
type InstallmentsAddedDebtIterator struct {
	Event *InstallmentsAddedDebt // Event containing the contract specifics and raw log

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
func (it *InstallmentsAddedDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsAddedDebt)
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
		it.Event = new(InstallmentsAddedDebt)
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
func (it *InstallmentsAddedDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsAddedDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsAddedDebt represents a AddedDebt event raised by the Installments contract.
type InstallmentsAddedDebt struct {
	Id     [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedDebt is a free log retrieval operation binding the contract event 0x2f39b1101bb100a4485571ad3147db16cd731b8fc83bc513c6b15acf4a24464f.
//
// Solidity: event AddedDebt(bytes32 indexed _id, uint256 _amount)
func (_Installments *InstallmentsFilterer) FilterAddedDebt(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsAddedDebtIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "AddedDebt", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsAddedDebtIterator{contract: _Installments.contract, event: "AddedDebt", logs: logs, sub: sub}, nil
}

// WatchAddedDebt is a free log subscription operation binding the contract event 0x2f39b1101bb100a4485571ad3147db16cd731b8fc83bc513c6b15acf4a24464f.
//
// Solidity: event AddedDebt(bytes32 indexed _id, uint256 _amount)
func (_Installments *InstallmentsFilterer) WatchAddedDebt(opts *bind.WatchOpts, sink chan<- *InstallmentsAddedDebt, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "AddedDebt", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsAddedDebt)
				if err := _Installments.contract.UnpackLog(event, "AddedDebt", log); err != nil {
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

// ParseAddedDebt is a log parse operation binding the contract event 0x2f39b1101bb100a4485571ad3147db16cd731b8fc83bc513c6b15acf4a24464f.
//
// Solidity: event AddedDebt(bytes32 indexed _id, uint256 _amount)
func (_Installments *InstallmentsFilterer) ParseAddedDebt(log types.Log) (*InstallmentsAddedDebt, error) {
	event := new(InstallmentsAddedDebt)
	if err := _Installments.contract.UnpackLog(event, "AddedDebt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsAddedPaidIterator is returned from FilterAddedPaid and is used to iterate over the raw logs and unpacked data for AddedPaid events raised by the Installments contract.
type InstallmentsAddedPaidIterator struct {
	Event *InstallmentsAddedPaid // Event containing the contract specifics and raw log

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
func (it *InstallmentsAddedPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsAddedPaid)
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
		it.Event = new(InstallmentsAddedPaid)
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
func (it *InstallmentsAddedPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsAddedPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsAddedPaid represents a AddedPaid event raised by the Installments contract.
type InstallmentsAddedPaid struct {
	Id   [32]byte
	Paid *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedPaid is a free log retrieval operation binding the contract event 0xc778ddfdb0d4ccac7447183c180c8559ec30cdb51c3726b6c58b3bd342596093.
//
// Solidity: event AddedPaid(bytes32 indexed _id, uint256 _paid)
func (_Installments *InstallmentsFilterer) FilterAddedPaid(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsAddedPaidIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "AddedPaid", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsAddedPaidIterator{contract: _Installments.contract, event: "AddedPaid", logs: logs, sub: sub}, nil
}

// WatchAddedPaid is a free log subscription operation binding the contract event 0xc778ddfdb0d4ccac7447183c180c8559ec30cdb51c3726b6c58b3bd342596093.
//
// Solidity: event AddedPaid(bytes32 indexed _id, uint256 _paid)
func (_Installments *InstallmentsFilterer) WatchAddedPaid(opts *bind.WatchOpts, sink chan<- *InstallmentsAddedPaid, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "AddedPaid", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsAddedPaid)
				if err := _Installments.contract.UnpackLog(event, "AddedPaid", log); err != nil {
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

// ParseAddedPaid is a log parse operation binding the contract event 0xc778ddfdb0d4ccac7447183c180c8559ec30cdb51c3726b6c58b3bd342596093.
//
// Solidity: event AddedPaid(bytes32 indexed _id, uint256 _paid)
func (_Installments *InstallmentsFilterer) ParseAddedPaid(log types.Log) (*InstallmentsAddedPaid, error) {
	event := new(InstallmentsAddedPaid)
	if err := _Installments.contract.UnpackLog(event, "AddedPaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsChangedDueTimeIterator is returned from FilterChangedDueTime and is used to iterate over the raw logs and unpacked data for ChangedDueTime events raised by the Installments contract.
type InstallmentsChangedDueTimeIterator struct {
	Event *InstallmentsChangedDueTime // Event containing the contract specifics and raw log

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
func (it *InstallmentsChangedDueTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsChangedDueTime)
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
		it.Event = new(InstallmentsChangedDueTime)
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
func (it *InstallmentsChangedDueTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsChangedDueTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsChangedDueTime represents a ChangedDueTime event raised by the Installments contract.
type InstallmentsChangedDueTime struct {
	Id        [32]byte
	Timestamp *big.Int
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedDueTime is a free log retrieval operation binding the contract event 0x8108c1693b12a3bca1add9e080d5917b7277a9c7a45f54c13c6cf211565229c9.
//
// Solidity: event ChangedDueTime(bytes32 indexed _id, uint256 _timestamp, uint256 _status)
func (_Installments *InstallmentsFilterer) FilterChangedDueTime(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsChangedDueTimeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "ChangedDueTime", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsChangedDueTimeIterator{contract: _Installments.contract, event: "ChangedDueTime", logs: logs, sub: sub}, nil
}

// WatchChangedDueTime is a free log subscription operation binding the contract event 0x8108c1693b12a3bca1add9e080d5917b7277a9c7a45f54c13c6cf211565229c9.
//
// Solidity: event ChangedDueTime(bytes32 indexed _id, uint256 _timestamp, uint256 _status)
func (_Installments *InstallmentsFilterer) WatchChangedDueTime(opts *bind.WatchOpts, sink chan<- *InstallmentsChangedDueTime, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "ChangedDueTime", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsChangedDueTime)
				if err := _Installments.contract.UnpackLog(event, "ChangedDueTime", log); err != nil {
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

// ParseChangedDueTime is a log parse operation binding the contract event 0x8108c1693b12a3bca1add9e080d5917b7277a9c7a45f54c13c6cf211565229c9.
//
// Solidity: event ChangedDueTime(bytes32 indexed _id, uint256 _timestamp, uint256 _status)
func (_Installments *InstallmentsFilterer) ParseChangedDueTime(log types.Log) (*InstallmentsChangedDueTime, error) {
	event := new(InstallmentsChangedDueTime)
	if err := _Installments.contract.UnpackLog(event, "ChangedDueTime", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsChangedFinalTimeIterator is returned from FilterChangedFinalTime and is used to iterate over the raw logs and unpacked data for ChangedFinalTime events raised by the Installments contract.
type InstallmentsChangedFinalTimeIterator struct {
	Event *InstallmentsChangedFinalTime // Event containing the contract specifics and raw log

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
func (it *InstallmentsChangedFinalTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsChangedFinalTime)
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
		it.Event = new(InstallmentsChangedFinalTime)
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
func (it *InstallmentsChangedFinalTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsChangedFinalTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsChangedFinalTime represents a ChangedFinalTime event raised by the Installments contract.
type InstallmentsChangedFinalTime struct {
	Id        [32]byte
	Timestamp *big.Int
	DueTime   uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedFinalTime is a free log retrieval operation binding the contract event 0x60e093518baf1919f536c54e83aada895ba93b5f32d096d8584e186ed3e317f1.
//
// Solidity: event ChangedFinalTime(bytes32 indexed _id, uint256 _timestamp, uint64 _dueTime)
func (_Installments *InstallmentsFilterer) FilterChangedFinalTime(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsChangedFinalTimeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "ChangedFinalTime", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsChangedFinalTimeIterator{contract: _Installments.contract, event: "ChangedFinalTime", logs: logs, sub: sub}, nil
}

// WatchChangedFinalTime is a free log subscription operation binding the contract event 0x60e093518baf1919f536c54e83aada895ba93b5f32d096d8584e186ed3e317f1.
//
// Solidity: event ChangedFinalTime(bytes32 indexed _id, uint256 _timestamp, uint64 _dueTime)
func (_Installments *InstallmentsFilterer) WatchChangedFinalTime(opts *bind.WatchOpts, sink chan<- *InstallmentsChangedFinalTime, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "ChangedFinalTime", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsChangedFinalTime)
				if err := _Installments.contract.UnpackLog(event, "ChangedFinalTime", log); err != nil {
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

// ParseChangedFinalTime is a log parse operation binding the contract event 0x60e093518baf1919f536c54e83aada895ba93b5f32d096d8584e186ed3e317f1.
//
// Solidity: event ChangedFinalTime(bytes32 indexed _id, uint256 _timestamp, uint64 _dueTime)
func (_Installments *InstallmentsFilterer) ParseChangedFinalTime(log types.Log) (*InstallmentsChangedFinalTime, error) {
	event := new(InstallmentsChangedFinalTime)
	if err := _Installments.contract.UnpackLog(event, "ChangedFinalTime", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsChangedFrequencyIterator is returned from FilterChangedFrequency and is used to iterate over the raw logs and unpacked data for ChangedFrequency events raised by the Installments contract.
type InstallmentsChangedFrequencyIterator struct {
	Event *InstallmentsChangedFrequency // Event containing the contract specifics and raw log

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
func (it *InstallmentsChangedFrequencyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsChangedFrequency)
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
		it.Event = new(InstallmentsChangedFrequency)
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
func (it *InstallmentsChangedFrequencyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsChangedFrequencyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsChangedFrequency represents a ChangedFrequency event raised by the Installments contract.
type InstallmentsChangedFrequency struct {
	Id        [32]byte
	Timestamp *big.Int
	Frequency *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedFrequency is a free log retrieval operation binding the contract event 0xc4a675a21f68bda459f7ad3bdfa98344e1449d550032d36aab2afd69d2f73ed9.
//
// Solidity: event ChangedFrequency(bytes32 indexed _id, uint256 _timestamp, uint256 _frequency)
func (_Installments *InstallmentsFilterer) FilterChangedFrequency(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsChangedFrequencyIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "ChangedFrequency", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsChangedFrequencyIterator{contract: _Installments.contract, event: "ChangedFrequency", logs: logs, sub: sub}, nil
}

// WatchChangedFrequency is a free log subscription operation binding the contract event 0xc4a675a21f68bda459f7ad3bdfa98344e1449d550032d36aab2afd69d2f73ed9.
//
// Solidity: event ChangedFrequency(bytes32 indexed _id, uint256 _timestamp, uint256 _frequency)
func (_Installments *InstallmentsFilterer) WatchChangedFrequency(opts *bind.WatchOpts, sink chan<- *InstallmentsChangedFrequency, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "ChangedFrequency", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsChangedFrequency)
				if err := _Installments.contract.UnpackLog(event, "ChangedFrequency", log); err != nil {
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

// ParseChangedFrequency is a log parse operation binding the contract event 0xc4a675a21f68bda459f7ad3bdfa98344e1449d550032d36aab2afd69d2f73ed9.
//
// Solidity: event ChangedFrequency(bytes32 indexed _id, uint256 _timestamp, uint256 _frequency)
func (_Installments *InstallmentsFilterer) ParseChangedFrequency(log types.Log) (*InstallmentsChangedFrequency, error) {
	event := new(InstallmentsChangedFrequency)
	if err := _Installments.contract.UnpackLog(event, "ChangedFrequency", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsChangedObligationIterator is returned from FilterChangedObligation and is used to iterate over the raw logs and unpacked data for ChangedObligation events raised by the Installments contract.
type InstallmentsChangedObligationIterator struct {
	Event *InstallmentsChangedObligation // Event containing the contract specifics and raw log

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
func (it *InstallmentsChangedObligationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsChangedObligation)
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
		it.Event = new(InstallmentsChangedObligation)
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
func (it *InstallmentsChangedObligationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsChangedObligationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsChangedObligation represents a ChangedObligation event raised by the Installments contract.
type InstallmentsChangedObligation struct {
	Id        [32]byte
	Timestamp *big.Int
	Debt      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedObligation is a free log retrieval operation binding the contract event 0x35a463d00829418b627fa74ee01c7981a42d59eb6b813193ca2b7ecb28c88d6b.
//
// Solidity: event ChangedObligation(bytes32 indexed _id, uint256 _timestamp, uint256 _debt)
func (_Installments *InstallmentsFilterer) FilterChangedObligation(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsChangedObligationIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "ChangedObligation", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsChangedObligationIterator{contract: _Installments.contract, event: "ChangedObligation", logs: logs, sub: sub}, nil
}

// WatchChangedObligation is a free log subscription operation binding the contract event 0x35a463d00829418b627fa74ee01c7981a42d59eb6b813193ca2b7ecb28c88d6b.
//
// Solidity: event ChangedObligation(bytes32 indexed _id, uint256 _timestamp, uint256 _debt)
func (_Installments *InstallmentsFilterer) WatchChangedObligation(opts *bind.WatchOpts, sink chan<- *InstallmentsChangedObligation, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "ChangedObligation", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsChangedObligation)
				if err := _Installments.contract.UnpackLog(event, "ChangedObligation", log); err != nil {
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

// ParseChangedObligation is a log parse operation binding the contract event 0x35a463d00829418b627fa74ee01c7981a42d59eb6b813193ca2b7ecb28c88d6b.
//
// Solidity: event ChangedObligation(bytes32 indexed _id, uint256 _timestamp, uint256 _debt)
func (_Installments *InstallmentsFilterer) ParseChangedObligation(log types.Log) (*InstallmentsChangedObligation, error) {
	event := new(InstallmentsChangedObligation)
	if err := _Installments.contract.UnpackLog(event, "ChangedObligation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsChangedStatusIterator is returned from FilterChangedStatus and is used to iterate over the raw logs and unpacked data for ChangedStatus events raised by the Installments contract.
type InstallmentsChangedStatusIterator struct {
	Event *InstallmentsChangedStatus // Event containing the contract specifics and raw log

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
func (it *InstallmentsChangedStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsChangedStatus)
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
		it.Event = new(InstallmentsChangedStatus)
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
func (it *InstallmentsChangedStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsChangedStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsChangedStatus represents a ChangedStatus event raised by the Installments contract.
type InstallmentsChangedStatus struct {
	Id        [32]byte
	Timestamp *big.Int
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedStatus is a free log retrieval operation binding the contract event 0x03202132f86ea915f0496d740ab57083235a3850003b398f13d3328f966e5e9e.
//
// Solidity: event ChangedStatus(bytes32 indexed _id, uint256 _timestamp, uint256 _status)
func (_Installments *InstallmentsFilterer) FilterChangedStatus(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsChangedStatusIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "ChangedStatus", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsChangedStatusIterator{contract: _Installments.contract, event: "ChangedStatus", logs: logs, sub: sub}, nil
}

// WatchChangedStatus is a free log subscription operation binding the contract event 0x03202132f86ea915f0496d740ab57083235a3850003b398f13d3328f966e5e9e.
//
// Solidity: event ChangedStatus(bytes32 indexed _id, uint256 _timestamp, uint256 _status)
func (_Installments *InstallmentsFilterer) WatchChangedStatus(opts *bind.WatchOpts, sink chan<- *InstallmentsChangedStatus, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "ChangedStatus", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsChangedStatus)
				if err := _Installments.contract.UnpackLog(event, "ChangedStatus", log); err != nil {
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

// ParseChangedStatus is a log parse operation binding the contract event 0x03202132f86ea915f0496d740ab57083235a3850003b398f13d3328f966e5e9e.
//
// Solidity: event ChangedStatus(bytes32 indexed _id, uint256 _timestamp, uint256 _status)
func (_Installments *InstallmentsFilterer) ParseChangedStatus(log types.Log) (*InstallmentsChangedStatus, error) {
	event := new(InstallmentsChangedStatus)
	if err := _Installments.contract.UnpackLog(event, "ChangedStatus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Installments contract.
type InstallmentsCreatedIterator struct {
	Event *InstallmentsCreated // Event containing the contract specifics and raw log

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
func (it *InstallmentsCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsCreated)
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
		it.Event = new(InstallmentsCreated)
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
func (it *InstallmentsCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsCreated represents a Created event raised by the Installments contract.
type InstallmentsCreated struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x102d25c49d33fcdb8976a3f2744e0785c98d9e43b88364859e6aec4ae82eff5c.
//
// Solidity: event Created(bytes32 indexed _id)
func (_Installments *InstallmentsFilterer) FilterCreated(opts *bind.FilterOpts, _id [][32]byte) (*InstallmentsCreatedIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "Created", _idRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsCreatedIterator{contract: _Installments.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x102d25c49d33fcdb8976a3f2744e0785c98d9e43b88364859e6aec4ae82eff5c.
//
// Solidity: event Created(bytes32 indexed _id)
func (_Installments *InstallmentsFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *InstallmentsCreated, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "Created", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsCreated)
				if err := _Installments.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x102d25c49d33fcdb8976a3f2744e0785c98d9e43b88364859e6aec4ae82eff5c.
//
// Solidity: event Created(bytes32 indexed _id)
func (_Installments *InstallmentsFilterer) ParseCreated(log types.Log) (*InstallmentsCreated, error) {
	event := new(InstallmentsCreated)
	if err := _Installments.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Installments contract.
type InstallmentsOwnershipTransferredIterator struct {
	Event *InstallmentsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InstallmentsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsOwnershipTransferred)
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
		it.Event = new(InstallmentsOwnershipTransferred)
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
func (it *InstallmentsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsOwnershipTransferred represents a OwnershipTransferred event raised by the Installments contract.
type InstallmentsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_Installments *InstallmentsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*InstallmentsOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Installments.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InstallmentsOwnershipTransferredIterator{contract: _Installments.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_Installments *InstallmentsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InstallmentsOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Installments.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsOwnershipTransferred)
				if err := _Installments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_Installments *InstallmentsFilterer) ParseOwnershipTransferred(log types.Log) (*InstallmentsOwnershipTransferred, error) {
	event := new(InstallmentsOwnershipTransferred)
	if err := _Installments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsSetClockIterator is returned from FilterSetClock and is used to iterate over the raw logs and unpacked data for SetClock events raised by the Installments contract.
type InstallmentsSetClockIterator struct {
	Event *InstallmentsSetClock // Event containing the contract specifics and raw log

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
func (it *InstallmentsSetClockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsSetClock)
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
		it.Event = new(InstallmentsSetClock)
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
func (it *InstallmentsSetClockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsSetClockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsSetClock represents a SetClock event raised by the Installments contract.
type InstallmentsSetClock struct {
	Id  [32]byte
	To  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetClock is a free log retrieval operation binding the contract event 0x5ac4222bff447e4964c68c4e68ba64f779bf5b79d7ba34c88fc1aa49f18eadbf.
//
// Solidity: event _setClock(bytes32 _id, uint64 _to)
func (_Installments *InstallmentsFilterer) FilterSetClock(opts *bind.FilterOpts) (*InstallmentsSetClockIterator, error) {

	logs, sub, err := _Installments.contract.FilterLogs(opts, "_setClock")
	if err != nil {
		return nil, err
	}
	return &InstallmentsSetClockIterator{contract: _Installments.contract, event: "_setClock", logs: logs, sub: sub}, nil
}

// WatchSetClock is a free log subscription operation binding the contract event 0x5ac4222bff447e4964c68c4e68ba64f779bf5b79d7ba34c88fc1aa49f18eadbf.
//
// Solidity: event _setClock(bytes32 _id, uint64 _to)
func (_Installments *InstallmentsFilterer) WatchSetClock(opts *bind.WatchOpts, sink chan<- *InstallmentsSetClock) (event.Subscription, error) {

	logs, sub, err := _Installments.contract.WatchLogs(opts, "_setClock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsSetClock)
				if err := _Installments.contract.UnpackLog(event, "_setClock", log); err != nil {
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

// ParseSetClock is a log parse operation binding the contract event 0x5ac4222bff447e4964c68c4e68ba64f779bf5b79d7ba34c88fc1aa49f18eadbf.
//
// Solidity: event _setClock(bytes32 _id, uint64 _to)
func (_Installments *InstallmentsFilterer) ParseSetClock(log types.Log) (*InstallmentsSetClock, error) {
	event := new(InstallmentsSetClock)
	if err := _Installments.contract.UnpackLog(event, "_setClock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsSetDescriptorIterator is returned from FilterSetDescriptor and is used to iterate over the raw logs and unpacked data for SetDescriptor events raised by the Installments contract.
type InstallmentsSetDescriptorIterator struct {
	Event *InstallmentsSetDescriptor // Event containing the contract specifics and raw log

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
func (it *InstallmentsSetDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsSetDescriptor)
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
		it.Event = new(InstallmentsSetDescriptor)
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
func (it *InstallmentsSetDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsSetDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsSetDescriptor represents a SetDescriptor event raised by the Installments contract.
type InstallmentsSetDescriptor struct {
	Descriptor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDescriptor is a free log retrieval operation binding the contract event 0xfd7a059fb02438ecc7334e2f3b4dbc55b37aeea735bb0391714041593a5beacd.
//
// Solidity: event _setDescriptor(address _descriptor)
func (_Installments *InstallmentsFilterer) FilterSetDescriptor(opts *bind.FilterOpts) (*InstallmentsSetDescriptorIterator, error) {

	logs, sub, err := _Installments.contract.FilterLogs(opts, "_setDescriptor")
	if err != nil {
		return nil, err
	}
	return &InstallmentsSetDescriptorIterator{contract: _Installments.contract, event: "_setDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetDescriptor is a free log subscription operation binding the contract event 0xfd7a059fb02438ecc7334e2f3b4dbc55b37aeea735bb0391714041593a5beacd.
//
// Solidity: event _setDescriptor(address _descriptor)
func (_Installments *InstallmentsFilterer) WatchSetDescriptor(opts *bind.WatchOpts, sink chan<- *InstallmentsSetDescriptor) (event.Subscription, error) {

	logs, sub, err := _Installments.contract.WatchLogs(opts, "_setDescriptor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsSetDescriptor)
				if err := _Installments.contract.UnpackLog(event, "_setDescriptor", log); err != nil {
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

// ParseSetDescriptor is a log parse operation binding the contract event 0xfd7a059fb02438ecc7334e2f3b4dbc55b37aeea735bb0391714041593a5beacd.
//
// Solidity: event _setDescriptor(address _descriptor)
func (_Installments *InstallmentsFilterer) ParseSetDescriptor(log types.Log) (*InstallmentsSetDescriptor, error) {
	event := new(InstallmentsSetDescriptor)
	if err := _Installments.contract.UnpackLog(event, "_setDescriptor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsSetEngineIterator is returned from FilterSetEngine and is used to iterate over the raw logs and unpacked data for SetEngine events raised by the Installments contract.
type InstallmentsSetEngineIterator struct {
	Event *InstallmentsSetEngine // Event containing the contract specifics and raw log

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
func (it *InstallmentsSetEngineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsSetEngine)
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
		it.Event = new(InstallmentsSetEngine)
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
func (it *InstallmentsSetEngineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsSetEngineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsSetEngine represents a SetEngine event raised by the Installments contract.
type InstallmentsSetEngine struct {
	Engine common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetEngine is a free log retrieval operation binding the contract event 0x0188a486877924f3e0702cef6a0566d0fc77155ab3edb363748f72c6c451ff43.
//
// Solidity: event _setEngine(address _engine)
func (_Installments *InstallmentsFilterer) FilterSetEngine(opts *bind.FilterOpts) (*InstallmentsSetEngineIterator, error) {

	logs, sub, err := _Installments.contract.FilterLogs(opts, "_setEngine")
	if err != nil {
		return nil, err
	}
	return &InstallmentsSetEngineIterator{contract: _Installments.contract, event: "_setEngine", logs: logs, sub: sub}, nil
}

// WatchSetEngine is a free log subscription operation binding the contract event 0x0188a486877924f3e0702cef6a0566d0fc77155ab3edb363748f72c6c451ff43.
//
// Solidity: event _setEngine(address _engine)
func (_Installments *InstallmentsFilterer) WatchSetEngine(opts *bind.WatchOpts, sink chan<- *InstallmentsSetEngine) (event.Subscription, error) {

	logs, sub, err := _Installments.contract.WatchLogs(opts, "_setEngine")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsSetEngine)
				if err := _Installments.contract.UnpackLog(event, "_setEngine", log); err != nil {
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

// ParseSetEngine is a log parse operation binding the contract event 0x0188a486877924f3e0702cef6a0566d0fc77155ab3edb363748f72c6c451ff43.
//
// Solidity: event _setEngine(address _engine)
func (_Installments *InstallmentsFilterer) ParseSetEngine(log types.Log) (*InstallmentsSetEngine, error) {
	event := new(InstallmentsSetEngine)
	if err := _Installments.contract.UnpackLog(event, "_setEngine", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsSetInterestIterator is returned from FilterSetInterest and is used to iterate over the raw logs and unpacked data for SetInterest events raised by the Installments contract.
type InstallmentsSetInterestIterator struct {
	Event *InstallmentsSetInterest // Event containing the contract specifics and raw log

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
func (it *InstallmentsSetInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsSetInterest)
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
		it.Event = new(InstallmentsSetInterest)
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
func (it *InstallmentsSetInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsSetInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsSetInterest represents a SetInterest event raised by the Installments contract.
type InstallmentsSetInterest struct {
	Id       [32]byte
	Interest *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetInterest is a free log retrieval operation binding the contract event 0xfe35a3cf12dd1d691d662ec765e7ad19e6f21c24923847936c5ec4c0cc7186ba.
//
// Solidity: event _setInterest(bytes32 _id, uint128 _interest)
func (_Installments *InstallmentsFilterer) FilterSetInterest(opts *bind.FilterOpts) (*InstallmentsSetInterestIterator, error) {

	logs, sub, err := _Installments.contract.FilterLogs(opts, "_setInterest")
	if err != nil {
		return nil, err
	}
	return &InstallmentsSetInterestIterator{contract: _Installments.contract, event: "_setInterest", logs: logs, sub: sub}, nil
}

// WatchSetInterest is a free log subscription operation binding the contract event 0xfe35a3cf12dd1d691d662ec765e7ad19e6f21c24923847936c5ec4c0cc7186ba.
//
// Solidity: event _setInterest(bytes32 _id, uint128 _interest)
func (_Installments *InstallmentsFilterer) WatchSetInterest(opts *bind.WatchOpts, sink chan<- *InstallmentsSetInterest) (event.Subscription, error) {

	logs, sub, err := _Installments.contract.WatchLogs(opts, "_setInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsSetInterest)
				if err := _Installments.contract.UnpackLog(event, "_setInterest", log); err != nil {
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

// ParseSetInterest is a log parse operation binding the contract event 0xfe35a3cf12dd1d691d662ec765e7ad19e6f21c24923847936c5ec4c0cc7186ba.
//
// Solidity: event _setInterest(bytes32 _id, uint128 _interest)
func (_Installments *InstallmentsFilterer) ParseSetInterest(log types.Log) (*InstallmentsSetInterest, error) {
	event := new(InstallmentsSetInterest)
	if err := _Installments.contract.UnpackLog(event, "_setInterest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstallmentsSetPaidBaseIterator is returned from FilterSetPaidBase and is used to iterate over the raw logs and unpacked data for SetPaidBase events raised by the Installments contract.
type InstallmentsSetPaidBaseIterator struct {
	Event *InstallmentsSetPaidBase // Event containing the contract specifics and raw log

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
func (it *InstallmentsSetPaidBaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstallmentsSetPaidBase)
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
		it.Event = new(InstallmentsSetPaidBase)
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
func (it *InstallmentsSetPaidBaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstallmentsSetPaidBaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstallmentsSetPaidBase represents a SetPaidBase event raised by the Installments contract.
type InstallmentsSetPaidBase struct {
	Id       [32]byte
	PaidBase *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPaidBase is a free log retrieval operation binding the contract event 0xbe7c4be01d96a9021826187a698b198967f64a0d27f8a4e27228ae337c2f380b.
//
// Solidity: event _setPaidBase(bytes32 _id, uint128 _paidBase)
func (_Installments *InstallmentsFilterer) FilterSetPaidBase(opts *bind.FilterOpts) (*InstallmentsSetPaidBaseIterator, error) {

	logs, sub, err := _Installments.contract.FilterLogs(opts, "_setPaidBase")
	if err != nil {
		return nil, err
	}
	return &InstallmentsSetPaidBaseIterator{contract: _Installments.contract, event: "_setPaidBase", logs: logs, sub: sub}, nil
}

// WatchSetPaidBase is a free log subscription operation binding the contract event 0xbe7c4be01d96a9021826187a698b198967f64a0d27f8a4e27228ae337c2f380b.
//
// Solidity: event _setPaidBase(bytes32 _id, uint128 _paidBase)
func (_Installments *InstallmentsFilterer) WatchSetPaidBase(opts *bind.WatchOpts, sink chan<- *InstallmentsSetPaidBase) (event.Subscription, error) {

	logs, sub, err := _Installments.contract.WatchLogs(opts, "_setPaidBase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstallmentsSetPaidBase)
				if err := _Installments.contract.UnpackLog(event, "_setPaidBase", log); err != nil {
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

// ParseSetPaidBase is a log parse operation binding the contract event 0xbe7c4be01d96a9021826187a698b198967f64a0d27f8a4e27228ae337c2f380b.
//
// Solidity: event _setPaidBase(bytes32 _id, uint128 _paidBase)
func (_Installments *InstallmentsFilterer) ParseSetPaidBase(log types.Log) (*InstallmentsSetPaidBase, error) {
	event := new(InstallmentsSetPaidBase)
	if err := _Installments.contract.UnpackLog(event, "_setPaidBase", log); err != nil {
		return nil, err
	}
	return event, nil
}
