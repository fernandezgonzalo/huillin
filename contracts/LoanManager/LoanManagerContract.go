// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LoanManager

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

// LoanManagerABI is the input ABI used to generate the binding from.
const LoanManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint128\"},{\"name\":\"_model\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_borrower\",\"type\":\"address\"},{\"name\":\"_callback\",\"type\":\"address\"},{\"name\":\"_salt\",\"type\":\"uint256\"},{\"name\":\"_expiration\",\"type\":\"uint64\"},{\"name\":\"_loanData\",\"type\":\"bytes\"}],\"name\":\"requestLoan\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestData\",\"type\":\"bytes\"},{\"name\":\"_loanData\",\"type\":\"bytes\"},{\"name\":\"_cosigner\",\"type\":\"address\"},{\"name\":\"_maxCosignerCost\",\"type\":\"uint256\"},{\"name\":\"_cosignerData\",\"type\":\"bytes\"},{\"name\":\"_oracleData\",\"type\":\"bytes\"},{\"name\":\"_creatorSig\",\"type\":\"bytes\"},{\"name\":\"_borrowerSig\",\"type\":\"bytes\"},{\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"settleLend\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getClosingObligation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestData\",\"type\":\"bytes\"},{\"name\":\"_loanData\",\"type\":\"bytes\"}],\"name\":\"settleCancel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDueTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint128\"},{\"name\":\"_model\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_borrower\",\"type\":\"address\"},{\"name\":\"_callback\",\"type\":\"address\"},{\"name\":\"_salt\",\"type\":\"uint256\"},{\"name\":\"_expiration\",\"type\":\"uint64\"},{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_loanData\",\"type\":\"bytes\"}],\"name\":\"encodeRequest\",\"outputs\":[{\"name\":\"requestData\",\"type\":\"bytes\"},{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"internalSalt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getCallback\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getLoanData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint128\"},{\"name\":\"_borrower\",\"type\":\"address\"},{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_model\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_callback\",\"type\":\"address\"},{\"name\":\"_salt\",\"type\":\"uint256\"},{\"name\":\"_expiration\",\"type\":\"uint64\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"calcId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"cosign\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getBorrower\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"approveRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debtEngine\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getBorrower\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerApproveRequest\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getCosigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getClosingObligation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"name\":\"open\",\"type\":\"bool\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"expiration\",\"type\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint128\"},{\"name\":\"cosigner\",\"type\":\"address\"},{\"name\":\"model\",\"type\":\"address\"},{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"callback\",\"type\":\"address\"},{\"name\":\"salt\",\"type\":\"uint256\"},{\"name\":\"loanData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getCosigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GAS_CALLBACK\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"canceledSettles\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getLoanData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getExpirationRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getDueTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getCurrency\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getCurrency\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_oracleData\",\"type\":\"bytes\"},{\"name\":\"_cosigner\",\"type\":\"address\"},{\"name\":\"_cosignerLimit\",\"type\":\"uint256\"},{\"name\":\"_cosignerData\",\"type\":\"bytes\"},{\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"lend\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint128\"},{\"name\":\"_borrower\",\"type\":\"address\"},{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_callback\",\"type\":\"address\"},{\"name\":\"_salt\",\"type\":\"uint256\"},{\"name\":\"_expiration\",\"type\":\"uint64\"}],\"name\":\"buildInternalSalt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getExpirationRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_debtEngine\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"_model\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_oracle\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_borrower\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_callback\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_salt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_loanData\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"Requested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_lender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"Lent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_cosigner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"Cosigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_canceler\",\"type\":\"address\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oracle\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalent\",\"type\":\"uint256\"}],\"name\":\"ReadedOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_response\",\"type\":\"bytes32\"}],\"name\":\"ApprovedRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_response\",\"type\":\"bytes32\"}],\"name\":\"ApprovedError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"ApprovedByCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"ApprovedBySignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"CreatorByCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"BorrowerByCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"CreatorBySignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"BorrowerBySignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_lender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"SettledLend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_canceler\",\"type\":\"address\"}],\"name\":\"SettledCancel\",\"type\":\"event\"}]"

// LoanManager is an auto generated Go binding around an Ethereum contract.
type LoanManager struct {
	LoanManagerCaller     // Read-only binding to the contract
	LoanManagerTransactor // Write-only binding to the contract
	LoanManagerFilterer   // Log filterer for contract events
}

// LoanManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoanManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanManagerSession struct {
	Contract     *LoanManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanManagerCallerSession struct {
	Contract *LoanManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LoanManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanManagerTransactorSession struct {
	Contract     *LoanManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LoanManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoanManagerRaw struct {
	Contract *LoanManager // Generic contract binding to access the raw methods on
}

// LoanManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanManagerCallerRaw struct {
	Contract *LoanManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LoanManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanManagerTransactorRaw struct {
	Contract *LoanManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoanManager creates a new instance of LoanManager, bound to a specific deployed contract.
func NewLoanManager(address common.Address, backend bind.ContractBackend) (*LoanManager, error) {
	contract, err := bindLoanManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoanManager{LoanManagerCaller: LoanManagerCaller{contract: contract}, LoanManagerTransactor: LoanManagerTransactor{contract: contract}, LoanManagerFilterer: LoanManagerFilterer{contract: contract}}, nil
}

// NewLoanManagerCaller creates a new read-only instance of LoanManager, bound to a specific deployed contract.
func NewLoanManagerCaller(address common.Address, caller bind.ContractCaller) (*LoanManagerCaller, error) {
	contract, err := bindLoanManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanManagerCaller{contract: contract}, nil
}

// NewLoanManagerTransactor creates a new write-only instance of LoanManager, bound to a specific deployed contract.
func NewLoanManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoanManagerTransactor, error) {
	contract, err := bindLoanManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanManagerTransactor{contract: contract}, nil
}

// NewLoanManagerFilterer creates a new log filterer instance of LoanManager, bound to a specific deployed contract.
func NewLoanManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoanManagerFilterer, error) {
	contract, err := bindLoanManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanManagerFilterer{contract: contract}, nil
}

// bindLoanManager binds a generic wrapper to an already deployed contract.
func bindLoanManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanManager *LoanManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanManager.Contract.LoanManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanManager *LoanManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanManager.Contract.LoanManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanManager *LoanManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanManager.Contract.LoanManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanManager *LoanManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanManager *LoanManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanManager *LoanManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanManager.Contract.contract.Transact(opts, method, params...)
}

// GASCALLBACK is a free data retrieval call binding the contract method 0xa9f649e8.
//
// Solidity: function GAS_CALLBACK() constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GASCALLBACK(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "GAS_CALLBACK")
	return *ret0, err
}

// GASCALLBACK is a free data retrieval call binding the contract method 0xa9f649e8.
//
// Solidity: function GAS_CALLBACK() constant returns(uint256)
func (_LoanManager *LoanManagerSession) GASCALLBACK() (*big.Int, error) {
	return _LoanManager.Contract.GASCALLBACK(&_LoanManager.CallOpts)
}

// GASCALLBACK is a free data retrieval call binding the contract method 0xa9f649e8.
//
// Solidity: function GAS_CALLBACK() constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GASCALLBACK() (*big.Int, error) {
	return _LoanManager.Contract.GASCALLBACK(&_LoanManager.CallOpts)
}

// BuildInternalSalt is a free data retrieval call binding the contract method 0xe3bab7ce.
//
// Solidity: function buildInternalSalt(uint128 _amount, address _borrower, address _creator, address _callback, uint256 _salt, uint64 _expiration) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) BuildInternalSalt(opts *bind.CallOpts, _amount *big.Int, _borrower common.Address, _creator common.Address, _callback common.Address, _salt *big.Int, _expiration uint64) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "buildInternalSalt", _amount, _borrower, _creator, _callback, _salt, _expiration)
	return *ret0, err
}

// BuildInternalSalt is a free data retrieval call binding the contract method 0xe3bab7ce.
//
// Solidity: function buildInternalSalt(uint128 _amount, address _borrower, address _creator, address _callback, uint256 _salt, uint64 _expiration) constant returns(uint256)
func (_LoanManager *LoanManagerSession) BuildInternalSalt(_amount *big.Int, _borrower common.Address, _creator common.Address, _callback common.Address, _salt *big.Int, _expiration uint64) (*big.Int, error) {
	return _LoanManager.Contract.BuildInternalSalt(&_LoanManager.CallOpts, _amount, _borrower, _creator, _callback, _salt, _expiration)
}

// BuildInternalSalt is a free data retrieval call binding the contract method 0xe3bab7ce.
//
// Solidity: function buildInternalSalt(uint128 _amount, address _borrower, address _creator, address _callback, uint256 _salt, uint64 _expiration) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) BuildInternalSalt(_amount *big.Int, _borrower common.Address, _creator common.Address, _callback common.Address, _salt *big.Int, _expiration uint64) (*big.Int, error) {
	return _LoanManager.Contract.BuildInternalSalt(&_LoanManager.CallOpts, _amount, _borrower, _creator, _callback, _salt, _expiration)
}

// CalcId is a free data retrieval call binding the contract method 0x629d2f34.
//
// Solidity: function calcId(uint128 _amount, address _borrower, address _creator, address _model, address _oracle, address _callback, uint256 _salt, uint64 _expiration, bytes _data) constant returns(bytes32)
func (_LoanManager *LoanManagerCaller) CalcId(opts *bind.CallOpts, _amount *big.Int, _borrower common.Address, _creator common.Address, _model common.Address, _oracle common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "calcId", _amount, _borrower, _creator, _model, _oracle, _callback, _salt, _expiration, _data)
	return *ret0, err
}

// CalcId is a free data retrieval call binding the contract method 0x629d2f34.
//
// Solidity: function calcId(uint128 _amount, address _borrower, address _creator, address _model, address _oracle, address _callback, uint256 _salt, uint64 _expiration, bytes _data) constant returns(bytes32)
func (_LoanManager *LoanManagerSession) CalcId(_amount *big.Int, _borrower common.Address, _creator common.Address, _model common.Address, _oracle common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _data []byte) ([32]byte, error) {
	return _LoanManager.Contract.CalcId(&_LoanManager.CallOpts, _amount, _borrower, _creator, _model, _oracle, _callback, _salt, _expiration, _data)
}

// CalcId is a free data retrieval call binding the contract method 0x629d2f34.
//
// Solidity: function calcId(uint128 _amount, address _borrower, address _creator, address _model, address _oracle, address _callback, uint256 _salt, uint64 _expiration, bytes _data) constant returns(bytes32)
func (_LoanManager *LoanManagerCallerSession) CalcId(_amount *big.Int, _borrower common.Address, _creator common.Address, _model common.Address, _oracle common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _data []byte) ([32]byte, error) {
	return _LoanManager.Contract.CalcId(&_LoanManager.CallOpts, _amount, _borrower, _creator, _model, _oracle, _callback, _salt, _expiration, _data)
}

// CanceledSettles is a free data retrieval call binding the contract method 0xabf6ac09.
//
// Solidity: function canceledSettles(bytes32 ) constant returns(bool)
func (_LoanManager *LoanManagerCaller) CanceledSettles(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "canceledSettles", arg0)
	return *ret0, err
}

// CanceledSettles is a free data retrieval call binding the contract method 0xabf6ac09.
//
// Solidity: function canceledSettles(bytes32 ) constant returns(bool)
func (_LoanManager *LoanManagerSession) CanceledSettles(arg0 [32]byte) (bool, error) {
	return _LoanManager.Contract.CanceledSettles(&_LoanManager.CallOpts, arg0)
}

// CanceledSettles is a free data retrieval call binding the contract method 0xabf6ac09.
//
// Solidity: function canceledSettles(bytes32 ) constant returns(bool)
func (_LoanManager *LoanManagerCallerSession) CanceledSettles(arg0 [32]byte) (bool, error) {
	return _LoanManager.Contract.CanceledSettles(&_LoanManager.CallOpts, arg0)
}

// DebtEngine is a free data retrieval call binding the contract method 0x7cdef83c.
//
// Solidity: function debtEngine() constant returns(address)
func (_LoanManager *LoanManagerCaller) DebtEngine(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "debtEngine")
	return *ret0, err
}

// DebtEngine is a free data retrieval call binding the contract method 0x7cdef83c.
//
// Solidity: function debtEngine() constant returns(address)
func (_LoanManager *LoanManagerSession) DebtEngine() (common.Address, error) {
	return _LoanManager.Contract.DebtEngine(&_LoanManager.CallOpts)
}

// DebtEngine is a free data retrieval call binding the contract method 0x7cdef83c.
//
// Solidity: function debtEngine() constant returns(address)
func (_LoanManager *LoanManagerCallerSession) DebtEngine() (common.Address, error) {
	return _LoanManager.Contract.DebtEngine(&_LoanManager.CallOpts)
}

// EncodeRequest is a free data retrieval call binding the contract method 0x5d9c4afe.
//
// Solidity: function encodeRequest(uint128 _amount, address _model, address _oracle, address _borrower, address _callback, uint256 _salt, uint64 _expiration, address _creator, bytes _loanData) constant returns(bytes requestData, bytes32 id)
func (_LoanManager *LoanManagerCaller) EncodeRequest(opts *bind.CallOpts, _amount *big.Int, _model common.Address, _oracle common.Address, _borrower common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _creator common.Address, _loanData []byte) (struct {
	RequestData []byte
	Id          [32]byte
}, error) {
	ret := new(struct {
		RequestData []byte
		Id          [32]byte
	})
	out := ret
	err := _LoanManager.contract.Call(opts, out, "encodeRequest", _amount, _model, _oracle, _borrower, _callback, _salt, _expiration, _creator, _loanData)
	return *ret, err
}

// EncodeRequest is a free data retrieval call binding the contract method 0x5d9c4afe.
//
// Solidity: function encodeRequest(uint128 _amount, address _model, address _oracle, address _borrower, address _callback, uint256 _salt, uint64 _expiration, address _creator, bytes _loanData) constant returns(bytes requestData, bytes32 id)
func (_LoanManager *LoanManagerSession) EncodeRequest(_amount *big.Int, _model common.Address, _oracle common.Address, _borrower common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _creator common.Address, _loanData []byte) (struct {
	RequestData []byte
	Id          [32]byte
}, error) {
	return _LoanManager.Contract.EncodeRequest(&_LoanManager.CallOpts, _amount, _model, _oracle, _borrower, _callback, _salt, _expiration, _creator, _loanData)
}

// EncodeRequest is a free data retrieval call binding the contract method 0x5d9c4afe.
//
// Solidity: function encodeRequest(uint128 _amount, address _model, address _oracle, address _borrower, address _callback, uint256 _salt, uint64 _expiration, address _creator, bytes _loanData) constant returns(bytes requestData, bytes32 id)
func (_LoanManager *LoanManagerCallerSession) EncodeRequest(_amount *big.Int, _model common.Address, _oracle common.Address, _borrower common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _creator common.Address, _loanData []byte) (struct {
	RequestData []byte
	Id          [32]byte
}, error) {
	return _LoanManager.Contract.EncodeRequest(&_LoanManager.CallOpts, _amount, _model, _oracle, _borrower, _callback, _salt, _expiration, _creator, _loanData)
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetAmount(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getAmount", _id)
	return *ret0, err
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetAmount(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetAmount(&_LoanManager.CallOpts, _id)
}

// GetAmount is a free data retrieval call binding the contract method 0x0e0a0d74.
//
// Solidity: function getAmount(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetAmount(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetAmount(&_LoanManager.CallOpts, _id)
}

// GetAmount0 is a free data retrieval call binding the contract method 0x9980ec86.
//
// Solidity: function getAmount(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetAmount0(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getAmount0", _id)
	return *ret0, err
}

// GetAmount0 is a free data retrieval call binding the contract method 0x9980ec86.
//
// Solidity: function getAmount(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetAmount0(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetAmount0(&_LoanManager.CallOpts, _id)
}

// GetAmount0 is a free data retrieval call binding the contract method 0x9980ec86.
//
// Solidity: function getAmount(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetAmount0(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetAmount0(&_LoanManager.CallOpts, _id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _id) constant returns(bool)
func (_LoanManager *LoanManagerCaller) GetApproved(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getApproved", _id)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _id) constant returns(bool)
func (_LoanManager *LoanManagerSession) GetApproved(_id *big.Int) (bool, error) {
	return _LoanManager.Contract.GetApproved(&_LoanManager.CallOpts, _id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _id) constant returns(bool)
func (_LoanManager *LoanManagerCallerSession) GetApproved(_id *big.Int) (bool, error) {
	return _LoanManager.Contract.GetApproved(&_LoanManager.CallOpts, _id)
}

// GetApproved0 is a free data retrieval call binding the contract method 0xef270ee2.
//
// Solidity: function getApproved(bytes32 _id) constant returns(bool)
func (_LoanManager *LoanManagerCaller) GetApproved0(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getApproved0", _id)
	return *ret0, err
}

// GetApproved0 is a free data retrieval call binding the contract method 0xef270ee2.
//
// Solidity: function getApproved(bytes32 _id) constant returns(bool)
func (_LoanManager *LoanManagerSession) GetApproved0(_id [32]byte) (bool, error) {
	return _LoanManager.Contract.GetApproved0(&_LoanManager.CallOpts, _id)
}

// GetApproved0 is a free data retrieval call binding the contract method 0xef270ee2.
//
// Solidity: function getApproved(bytes32 _id) constant returns(bool)
func (_LoanManager *LoanManagerCallerSession) GetApproved0(_id [32]byte) (bool, error) {
	return _LoanManager.Contract.GetApproved0(&_LoanManager.CallOpts, _id)
}

// GetBorrower is a free data retrieval call binding the contract method 0x728c3ea2.
//
// Solidity: function getBorrower(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetBorrower(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getBorrower", _id)
	return *ret0, err
}

// GetBorrower is a free data retrieval call binding the contract method 0x728c3ea2.
//
// Solidity: function getBorrower(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetBorrower(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetBorrower(&_LoanManager.CallOpts, _id)
}

// GetBorrower is a free data retrieval call binding the contract method 0x728c3ea2.
//
// Solidity: function getBorrower(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetBorrower(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetBorrower(&_LoanManager.CallOpts, _id)
}

// GetBorrower0 is a free data retrieval call binding the contract method 0x8500d919.
//
// Solidity: function getBorrower(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetBorrower0(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getBorrower0", _id)
	return *ret0, err
}

// GetBorrower0 is a free data retrieval call binding the contract method 0x8500d919.
//
// Solidity: function getBorrower(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetBorrower0(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetBorrower0(&_LoanManager.CallOpts, _id)
}

// GetBorrower0 is a free data retrieval call binding the contract method 0x8500d919.
//
// Solidity: function getBorrower(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetBorrower0(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetBorrower0(&_LoanManager.CallOpts, _id)
}

// GetCallback is a free data retrieval call binding the contract method 0x6135a0ad.
//
// Solidity: function getCallback(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetCallback(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCallback", _id)
	return *ret0, err
}

// GetCallback is a free data retrieval call binding the contract method 0x6135a0ad.
//
// Solidity: function getCallback(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetCallback(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetCallback(&_LoanManager.CallOpts, _id)
}

// GetCallback is a free data retrieval call binding the contract method 0x6135a0ad.
//
// Solidity: function getCallback(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetCallback(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetCallback(&_LoanManager.CallOpts, _id)
}

// GetClosingObligation is a free data retrieval call binding the contract method 0x51cbb299.
//
// Solidity: function getClosingObligation(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetClosingObligation(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getClosingObligation", _id)
	return *ret0, err
}

// GetClosingObligation is a free data retrieval call binding the contract method 0x51cbb299.
//
// Solidity: function getClosingObligation(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetClosingObligation(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetClosingObligation(&_LoanManager.CallOpts, _id)
}

// GetClosingObligation is a free data retrieval call binding the contract method 0x51cbb299.
//
// Solidity: function getClosingObligation(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetClosingObligation(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetClosingObligation(&_LoanManager.CallOpts, _id)
}

// GetClosingObligation0 is a free data retrieval call binding the contract method 0x9a6203e9.
//
// Solidity: function getClosingObligation(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetClosingObligation0(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getClosingObligation0", _id)
	return *ret0, err
}

// GetClosingObligation0 is a free data retrieval call binding the contract method 0x9a6203e9.
//
// Solidity: function getClosingObligation(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetClosingObligation0(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetClosingObligation0(&_LoanManager.CallOpts, _id)
}

// GetClosingObligation0 is a free data retrieval call binding the contract method 0x9a6203e9.
//
// Solidity: function getClosingObligation(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetClosingObligation0(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetClosingObligation0(&_LoanManager.CallOpts, _id)
}

// GetCosigner is a free data retrieval call binding the contract method 0x964f61f9.
//
// Solidity: function getCosigner(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetCosigner(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCosigner", _id)
	return *ret0, err
}

// GetCosigner is a free data retrieval call binding the contract method 0x964f61f9.
//
// Solidity: function getCosigner(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetCosigner(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetCosigner(&_LoanManager.CallOpts, _id)
}

// GetCosigner is a free data retrieval call binding the contract method 0x964f61f9.
//
// Solidity: function getCosigner(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetCosigner(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetCosigner(&_LoanManager.CallOpts, _id)
}

// GetCosigner0 is a free data retrieval call binding the contract method 0xa56d9dda.
//
// Solidity: function getCosigner(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetCosigner0(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCosigner0", _id)
	return *ret0, err
}

// GetCosigner0 is a free data retrieval call binding the contract method 0xa56d9dda.
//
// Solidity: function getCosigner(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetCosigner0(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetCosigner0(&_LoanManager.CallOpts, _id)
}

// GetCosigner0 is a free data retrieval call binding the contract method 0xa56d9dda.
//
// Solidity: function getCosigner(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetCosigner0(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetCosigner0(&_LoanManager.CallOpts, _id)
}

// GetCreator is a free data retrieval call binding the contract method 0x060a7ef1.
//
// Solidity: function getCreator(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetCreator(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCreator", _id)
	return *ret0, err
}

// GetCreator is a free data retrieval call binding the contract method 0x060a7ef1.
//
// Solidity: function getCreator(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetCreator(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetCreator(&_LoanManager.CallOpts, _id)
}

// GetCreator is a free data retrieval call binding the contract method 0x060a7ef1.
//
// Solidity: function getCreator(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetCreator(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetCreator(&_LoanManager.CallOpts, _id)
}

// GetCreator0 is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetCreator0(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCreator0", _id)
	return *ret0, err
}

// GetCreator0 is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetCreator0(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetCreator0(&_LoanManager.CallOpts, _id)
}

// GetCreator0 is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetCreator0(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetCreator0(&_LoanManager.CallOpts, _id)
}

// GetCurrency is a free data retrieval call binding the contract method 0xcdf9b77e.
//
// Solidity: function getCurrency(uint256 _id) constant returns(bytes32)
func (_LoanManager *LoanManagerCaller) GetCurrency(opts *bind.CallOpts, _id *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCurrency", _id)
	return *ret0, err
}

// GetCurrency is a free data retrieval call binding the contract method 0xcdf9b77e.
//
// Solidity: function getCurrency(uint256 _id) constant returns(bytes32)
func (_LoanManager *LoanManagerSession) GetCurrency(_id *big.Int) ([32]byte, error) {
	return _LoanManager.Contract.GetCurrency(&_LoanManager.CallOpts, _id)
}

// GetCurrency is a free data retrieval call binding the contract method 0xcdf9b77e.
//
// Solidity: function getCurrency(uint256 _id) constant returns(bytes32)
func (_LoanManager *LoanManagerCallerSession) GetCurrency(_id *big.Int) ([32]byte, error) {
	return _LoanManager.Contract.GetCurrency(&_LoanManager.CallOpts, _id)
}

// GetCurrency0 is a free data retrieval call binding the contract method 0xd383b80d.
//
// Solidity: function getCurrency(bytes32 _id) constant returns(bytes32)
func (_LoanManager *LoanManagerCaller) GetCurrency0(opts *bind.CallOpts, _id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getCurrency0", _id)
	return *ret0, err
}

// GetCurrency0 is a free data retrieval call binding the contract method 0xd383b80d.
//
// Solidity: function getCurrency(bytes32 _id) constant returns(bytes32)
func (_LoanManager *LoanManagerSession) GetCurrency0(_id [32]byte) ([32]byte, error) {
	return _LoanManager.Contract.GetCurrency0(&_LoanManager.CallOpts, _id)
}

// GetCurrency0 is a free data retrieval call binding the contract method 0xd383b80d.
//
// Solidity: function getCurrency(bytes32 _id) constant returns(bytes32)
func (_LoanManager *LoanManagerCallerSession) GetCurrency0(_id [32]byte) ([32]byte, error) {
	return _LoanManager.Contract.GetCurrency0(&_LoanManager.CallOpts, _id)
}

// GetDueTime is a free data retrieval call binding the contract method 0x59357045.
//
// Solidity: function getDueTime(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetDueTime(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getDueTime", _id)
	return *ret0, err
}

// GetDueTime is a free data retrieval call binding the contract method 0x59357045.
//
// Solidity: function getDueTime(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetDueTime(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetDueTime(&_LoanManager.CallOpts, _id)
}

// GetDueTime is a free data retrieval call binding the contract method 0x59357045.
//
// Solidity: function getDueTime(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetDueTime(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetDueTime(&_LoanManager.CallOpts, _id)
}

// GetDueTime0 is a free data retrieval call binding the contract method 0xcdd8750e.
//
// Solidity: function getDueTime(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetDueTime0(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getDueTime0", _id)
	return *ret0, err
}

// GetDueTime0 is a free data retrieval call binding the contract method 0xcdd8750e.
//
// Solidity: function getDueTime(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetDueTime0(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetDueTime0(&_LoanManager.CallOpts, _id)
}

// GetDueTime0 is a free data retrieval call binding the contract method 0xcdd8750e.
//
// Solidity: function getDueTime(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetDueTime0(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetDueTime0(&_LoanManager.CallOpts, _id)
}

// GetExpirationRequest is a free data retrieval call binding the contract method 0xc955bde2.
//
// Solidity: function getExpirationRequest(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetExpirationRequest(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getExpirationRequest", _id)
	return *ret0, err
}

// GetExpirationRequest is a free data retrieval call binding the contract method 0xc955bde2.
//
// Solidity: function getExpirationRequest(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetExpirationRequest(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetExpirationRequest(&_LoanManager.CallOpts, _id)
}

// GetExpirationRequest is a free data retrieval call binding the contract method 0xc955bde2.
//
// Solidity: function getExpirationRequest(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetExpirationRequest(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetExpirationRequest(&_LoanManager.CallOpts, _id)
}

// GetExpirationRequest0 is a free data retrieval call binding the contract method 0xe6c8fcf1.
//
// Solidity: function getExpirationRequest(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetExpirationRequest0(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getExpirationRequest0", _id)
	return *ret0, err
}

// GetExpirationRequest0 is a free data retrieval call binding the contract method 0xe6c8fcf1.
//
// Solidity: function getExpirationRequest(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetExpirationRequest0(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetExpirationRequest0(&_LoanManager.CallOpts, _id)
}

// GetExpirationRequest0 is a free data retrieval call binding the contract method 0xe6c8fcf1.
//
// Solidity: function getExpirationRequest(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetExpirationRequest0(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetExpirationRequest0(&_LoanManager.CallOpts, _id)
}

// GetLoanData is a free data retrieval call binding the contract method 0x622fe39f.
//
// Solidity: function getLoanData(uint256 _id) constant returns(bytes)
func (_LoanManager *LoanManagerCaller) GetLoanData(opts *bind.CallOpts, _id *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getLoanData", _id)
	return *ret0, err
}

// GetLoanData is a free data retrieval call binding the contract method 0x622fe39f.
//
// Solidity: function getLoanData(uint256 _id) constant returns(bytes)
func (_LoanManager *LoanManagerSession) GetLoanData(_id *big.Int) ([]byte, error) {
	return _LoanManager.Contract.GetLoanData(&_LoanManager.CallOpts, _id)
}

// GetLoanData is a free data retrieval call binding the contract method 0x622fe39f.
//
// Solidity: function getLoanData(uint256 _id) constant returns(bytes)
func (_LoanManager *LoanManagerCallerSession) GetLoanData(_id *big.Int) ([]byte, error) {
	return _LoanManager.Contract.GetLoanData(&_LoanManager.CallOpts, _id)
}

// GetLoanData0 is a free data retrieval call binding the contract method 0xc4d2b1b3.
//
// Solidity: function getLoanData(bytes32 _id) constant returns(bytes)
func (_LoanManager *LoanManagerCaller) GetLoanData0(opts *bind.CallOpts, _id [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getLoanData0", _id)
	return *ret0, err
}

// GetLoanData0 is a free data retrieval call binding the contract method 0xc4d2b1b3.
//
// Solidity: function getLoanData(bytes32 _id) constant returns(bytes)
func (_LoanManager *LoanManagerSession) GetLoanData0(_id [32]byte) ([]byte, error) {
	return _LoanManager.Contract.GetLoanData0(&_LoanManager.CallOpts, _id)
}

// GetLoanData0 is a free data retrieval call binding the contract method 0xc4d2b1b3.
//
// Solidity: function getLoanData(bytes32 _id) constant returns(bytes)
func (_LoanManager *LoanManagerCallerSession) GetLoanData0(_id [32]byte) ([]byte, error) {
	return _LoanManager.Contract.GetLoanData0(&_LoanManager.CallOpts, _id)
}

// GetOracle is a free data retrieval call binding the contract method 0x10a9de60.
//
// Solidity: function getOracle(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetOracle(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getOracle", _id)
	return *ret0, err
}

// GetOracle is a free data retrieval call binding the contract method 0x10a9de60.
//
// Solidity: function getOracle(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetOracle(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetOracle(&_LoanManager.CallOpts, _id)
}

// GetOracle is a free data retrieval call binding the contract method 0x10a9de60.
//
// Solidity: function getOracle(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetOracle(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.GetOracle(&_LoanManager.CallOpts, _id)
}

// GetOracle0 is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) GetOracle0(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getOracle0", _id)
	return *ret0, err
}

// GetOracle0 is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) GetOracle0(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetOracle0(&_LoanManager.CallOpts, _id)
}

// GetOracle0 is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) GetOracle0(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.GetOracle0(&_LoanManager.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetStatus(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getStatus", _id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetStatus(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetStatus(&_LoanManager.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetStatus(_id *big.Int) (*big.Int, error) {
	return _LoanManager.Contract.GetStatus(&_LoanManager.CallOpts, _id)
}

// GetStatus0 is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) GetStatus0(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "getStatus0", _id)
	return *ret0, err
}

// GetStatus0 is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) GetStatus0(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetStatus0(&_LoanManager.CallOpts, _id)
}

// GetStatus0 is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) GetStatus0(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.GetStatus0(&_LoanManager.CallOpts, _id)
}

// InternalSalt is a free data retrieval call binding the contract method 0x5f7cfe49.
//
// Solidity: function internalSalt(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCaller) InternalSalt(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "internalSalt", _id)
	return *ret0, err
}

// InternalSalt is a free data retrieval call binding the contract method 0x5f7cfe49.
//
// Solidity: function internalSalt(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerSession) InternalSalt(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.InternalSalt(&_LoanManager.CallOpts, _id)
}

// InternalSalt is a free data retrieval call binding the contract method 0x5f7cfe49.
//
// Solidity: function internalSalt(bytes32 _id) constant returns(uint256)
func (_LoanManager *LoanManagerCallerSession) InternalSalt(_id [32]byte) (*big.Int, error) {
	return _LoanManager.Contract.InternalSalt(&_LoanManager.CallOpts, _id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) OwnerOf(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "ownerOf", _id)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) OwnerOf(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.OwnerOf(&_LoanManager.CallOpts, _id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) OwnerOf(_id *big.Int) (common.Address, error) {
	return _LoanManager.Contract.OwnerOf(&_LoanManager.CallOpts, _id)
}

// OwnerOf0 is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCaller) OwnerOf0(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "ownerOf0", _id)
	return *ret0, err
}

// OwnerOf0 is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerSession) OwnerOf0(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.OwnerOf0(&_LoanManager.CallOpts, _id)
}

// OwnerOf0 is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 _id) constant returns(address)
func (_LoanManager *LoanManagerCallerSession) OwnerOf0(_id [32]byte) (common.Address, error) {
	return _LoanManager.Contract.OwnerOf0(&_LoanManager.CallOpts, _id)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(bool open, bool approved, uint64 expiration, uint128 amount, address cosigner, address model, address creator, address oracle, address borrower, address callback, uint256 salt, bytes loanData)
func (_LoanManager *LoanManagerCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Open       bool
	Approved   bool
	Expiration uint64
	Amount     *big.Int
	Cosigner   common.Address
	Model      common.Address
	Creator    common.Address
	Oracle     common.Address
	Borrower   common.Address
	Callback   common.Address
	Salt       *big.Int
	LoanData   []byte
}, error) {
	ret := new(struct {
		Open       bool
		Approved   bool
		Expiration uint64
		Amount     *big.Int
		Cosigner   common.Address
		Model      common.Address
		Creator    common.Address
		Oracle     common.Address
		Borrower   common.Address
		Callback   common.Address
		Salt       *big.Int
		LoanData   []byte
	})
	out := ret
	err := _LoanManager.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(bool open, bool approved, uint64 expiration, uint128 amount, address cosigner, address model, address creator, address oracle, address borrower, address callback, uint256 salt, bytes loanData)
func (_LoanManager *LoanManagerSession) Requests(arg0 [32]byte) (struct {
	Open       bool
	Approved   bool
	Expiration uint64
	Amount     *big.Int
	Cosigner   common.Address
	Model      common.Address
	Creator    common.Address
	Oracle     common.Address
	Borrower   common.Address
	Callback   common.Address
	Salt       *big.Int
	LoanData   []byte
}, error) {
	return _LoanManager.Contract.Requests(&_LoanManager.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(bool open, bool approved, uint64 expiration, uint128 amount, address cosigner, address model, address creator, address oracle, address borrower, address callback, uint256 salt, bytes loanData)
func (_LoanManager *LoanManagerCallerSession) Requests(arg0 [32]byte) (struct {
	Open       bool
	Approved   bool
	Expiration uint64
	Amount     *big.Int
	Cosigner   common.Address
	Model      common.Address
	Creator    common.Address
	Oracle     common.Address
	Borrower   common.Address
	Callback   common.Address
	Salt       *big.Int
	LoanData   []byte
}, error) {
	return _LoanManager.Contract.Requests(&_LoanManager.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_LoanManager *LoanManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanManager.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_LoanManager *LoanManagerSession) Token() (common.Address, error) {
	return _LoanManager.Contract.Token(&_LoanManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_LoanManager *LoanManagerCallerSession) Token() (common.Address, error) {
	return _LoanManager.Contract.Token(&_LoanManager.CallOpts)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x76ba6009.
//
// Solidity: function approveRequest(bytes32 _id) returns(bool)
func (_LoanManager *LoanManagerTransactor) ApproveRequest(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "approveRequest", _id)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x76ba6009.
//
// Solidity: function approveRequest(bytes32 _id) returns(bool)
func (_LoanManager *LoanManagerSession) ApproveRequest(_id [32]byte) (*types.Transaction, error) {
	return _LoanManager.Contract.ApproveRequest(&_LoanManager.TransactOpts, _id)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x76ba6009.
//
// Solidity: function approveRequest(bytes32 _id) returns(bool)
func (_LoanManager *LoanManagerTransactorSession) ApproveRequest(_id [32]byte) (*types.Transaction, error) {
	return _LoanManager.Contract.ApproveRequest(&_LoanManager.TransactOpts, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 _id) returns(bool)
func (_LoanManager *LoanManagerTransactor) Cancel(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "cancel", _id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 _id) returns(bool)
func (_LoanManager *LoanManagerSession) Cancel(_id [32]byte) (*types.Transaction, error) {
	return _LoanManager.Contract.Cancel(&_LoanManager.TransactOpts, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 _id) returns(bool)
func (_LoanManager *LoanManagerTransactorSession) Cancel(_id [32]byte) (*types.Transaction, error) {
	return _LoanManager.Contract.Cancel(&_LoanManager.TransactOpts, _id)
}

// Cosign is a paid mutator transaction binding the contract method 0x6394536d.
//
// Solidity: function cosign(uint256 _id, uint256 _cost) returns(bool)
func (_LoanManager *LoanManagerTransactor) Cosign(opts *bind.TransactOpts, _id *big.Int, _cost *big.Int) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "cosign", _id, _cost)
}

// Cosign is a paid mutator transaction binding the contract method 0x6394536d.
//
// Solidity: function cosign(uint256 _id, uint256 _cost) returns(bool)
func (_LoanManager *LoanManagerSession) Cosign(_id *big.Int, _cost *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.Cosign(&_LoanManager.TransactOpts, _id, _cost)
}

// Cosign is a paid mutator transaction binding the contract method 0x6394536d.
//
// Solidity: function cosign(uint256 _id, uint256 _cost) returns(bool)
func (_LoanManager *LoanManagerTransactorSession) Cosign(_id *big.Int, _cost *big.Int) (*types.Transaction, error) {
	return _LoanManager.Contract.Cosign(&_LoanManager.TransactOpts, _id, _cost)
}

// Lend is a paid mutator transaction binding the contract method 0xdb97eb44.
//
// Solidity: function lend(bytes32 _id, bytes _oracleData, address _cosigner, uint256 _cosignerLimit, bytes _cosignerData, bytes _callbackData) returns(bool)
func (_LoanManager *LoanManagerTransactor) Lend(opts *bind.TransactOpts, _id [32]byte, _oracleData []byte, _cosigner common.Address, _cosignerLimit *big.Int, _cosignerData []byte, _callbackData []byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "lend", _id, _oracleData, _cosigner, _cosignerLimit, _cosignerData, _callbackData)
}

// Lend is a paid mutator transaction binding the contract method 0xdb97eb44.
//
// Solidity: function lend(bytes32 _id, bytes _oracleData, address _cosigner, uint256 _cosignerLimit, bytes _cosignerData, bytes _callbackData) returns(bool)
func (_LoanManager *LoanManagerSession) Lend(_id [32]byte, _oracleData []byte, _cosigner common.Address, _cosignerLimit *big.Int, _cosignerData []byte, _callbackData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.Lend(&_LoanManager.TransactOpts, _id, _oracleData, _cosigner, _cosignerLimit, _cosignerData, _callbackData)
}

// Lend is a paid mutator transaction binding the contract method 0xdb97eb44.
//
// Solidity: function lend(bytes32 _id, bytes _oracleData, address _cosigner, uint256 _cosignerLimit, bytes _cosignerData, bytes _callbackData) returns(bool)
func (_LoanManager *LoanManagerTransactorSession) Lend(_id [32]byte, _oracleData []byte, _cosigner common.Address, _cosignerLimit *big.Int, _cosignerData []byte, _callbackData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.Lend(&_LoanManager.TransactOpts, _id, _oracleData, _cosigner, _cosignerLimit, _cosignerData, _callbackData)
}

// RegisterApproveRequest is a paid mutator transaction binding the contract method 0x90cd0d06.
//
// Solidity: function registerApproveRequest(bytes32 _id, bytes _signature) returns(bool approved)
func (_LoanManager *LoanManagerTransactor) RegisterApproveRequest(opts *bind.TransactOpts, _id [32]byte, _signature []byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "registerApproveRequest", _id, _signature)
}

// RegisterApproveRequest is a paid mutator transaction binding the contract method 0x90cd0d06.
//
// Solidity: function registerApproveRequest(bytes32 _id, bytes _signature) returns(bool approved)
func (_LoanManager *LoanManagerSession) RegisterApproveRequest(_id [32]byte, _signature []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.RegisterApproveRequest(&_LoanManager.TransactOpts, _id, _signature)
}

// RegisterApproveRequest is a paid mutator transaction binding the contract method 0x90cd0d06.
//
// Solidity: function registerApproveRequest(bytes32 _id, bytes _signature) returns(bool approved)
func (_LoanManager *LoanManagerTransactorSession) RegisterApproveRequest(_id [32]byte, _signature []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.RegisterApproveRequest(&_LoanManager.TransactOpts, _id, _signature)
}

// RequestLoan is a paid mutator transaction binding the contract method 0x167d7603.
//
// Solidity: function requestLoan(uint128 _amount, address _model, address _oracle, address _borrower, address _callback, uint256 _salt, uint64 _expiration, bytes _loanData) returns(bytes32 id)
func (_LoanManager *LoanManagerTransactor) RequestLoan(opts *bind.TransactOpts, _amount *big.Int, _model common.Address, _oracle common.Address, _borrower common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _loanData []byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "requestLoan", _amount, _model, _oracle, _borrower, _callback, _salt, _expiration, _loanData)
}

// RequestLoan is a paid mutator transaction binding the contract method 0x167d7603.
//
// Solidity: function requestLoan(uint128 _amount, address _model, address _oracle, address _borrower, address _callback, uint256 _salt, uint64 _expiration, bytes _loanData) returns(bytes32 id)
func (_LoanManager *LoanManagerSession) RequestLoan(_amount *big.Int, _model common.Address, _oracle common.Address, _borrower common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _loanData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.RequestLoan(&_LoanManager.TransactOpts, _amount, _model, _oracle, _borrower, _callback, _salt, _expiration, _loanData)
}

// RequestLoan is a paid mutator transaction binding the contract method 0x167d7603.
//
// Solidity: function requestLoan(uint128 _amount, address _model, address _oracle, address _borrower, address _callback, uint256 _salt, uint64 _expiration, bytes _loanData) returns(bytes32 id)
func (_LoanManager *LoanManagerTransactorSession) RequestLoan(_amount *big.Int, _model common.Address, _oracle common.Address, _borrower common.Address, _callback common.Address, _salt *big.Int, _expiration uint64, _loanData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.RequestLoan(&_LoanManager.TransactOpts, _amount, _model, _oracle, _borrower, _callback, _salt, _expiration, _loanData)
}

// SettleCancel is a paid mutator transaction binding the contract method 0x5900477a.
//
// Solidity: function settleCancel(bytes _requestData, bytes _loanData) returns(bool)
func (_LoanManager *LoanManagerTransactor) SettleCancel(opts *bind.TransactOpts, _requestData []byte, _loanData []byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "settleCancel", _requestData, _loanData)
}

// SettleCancel is a paid mutator transaction binding the contract method 0x5900477a.
//
// Solidity: function settleCancel(bytes _requestData, bytes _loanData) returns(bool)
func (_LoanManager *LoanManagerSession) SettleCancel(_requestData []byte, _loanData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.SettleCancel(&_LoanManager.TransactOpts, _requestData, _loanData)
}

// SettleCancel is a paid mutator transaction binding the contract method 0x5900477a.
//
// Solidity: function settleCancel(bytes _requestData, bytes _loanData) returns(bool)
func (_LoanManager *LoanManagerTransactorSession) SettleCancel(_requestData []byte, _loanData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.SettleCancel(&_LoanManager.TransactOpts, _requestData, _loanData)
}

// SettleLend is a paid mutator transaction binding the contract method 0x385584bf.
//
// Solidity: function settleLend(bytes _requestData, bytes _loanData, address _cosigner, uint256 _maxCosignerCost, bytes _cosignerData, bytes _oracleData, bytes _creatorSig, bytes _borrowerSig, bytes _callbackData) returns(bytes32 id)
func (_LoanManager *LoanManagerTransactor) SettleLend(opts *bind.TransactOpts, _requestData []byte, _loanData []byte, _cosigner common.Address, _maxCosignerCost *big.Int, _cosignerData []byte, _oracleData []byte, _creatorSig []byte, _borrowerSig []byte, _callbackData []byte) (*types.Transaction, error) {
	return _LoanManager.contract.Transact(opts, "settleLend", _requestData, _loanData, _cosigner, _maxCosignerCost, _cosignerData, _oracleData, _creatorSig, _borrowerSig, _callbackData)
}

// SettleLend is a paid mutator transaction binding the contract method 0x385584bf.
//
// Solidity: function settleLend(bytes _requestData, bytes _loanData, address _cosigner, uint256 _maxCosignerCost, bytes _cosignerData, bytes _oracleData, bytes _creatorSig, bytes _borrowerSig, bytes _callbackData) returns(bytes32 id)
func (_LoanManager *LoanManagerSession) SettleLend(_requestData []byte, _loanData []byte, _cosigner common.Address, _maxCosignerCost *big.Int, _cosignerData []byte, _oracleData []byte, _creatorSig []byte, _borrowerSig []byte, _callbackData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.SettleLend(&_LoanManager.TransactOpts, _requestData, _loanData, _cosigner, _maxCosignerCost, _cosignerData, _oracleData, _creatorSig, _borrowerSig, _callbackData)
}

// SettleLend is a paid mutator transaction binding the contract method 0x385584bf.
//
// Solidity: function settleLend(bytes _requestData, bytes _loanData, address _cosigner, uint256 _maxCosignerCost, bytes _cosignerData, bytes _oracleData, bytes _creatorSig, bytes _borrowerSig, bytes _callbackData) returns(bytes32 id)
func (_LoanManager *LoanManagerTransactorSession) SettleLend(_requestData []byte, _loanData []byte, _cosigner common.Address, _maxCosignerCost *big.Int, _cosignerData []byte, _oracleData []byte, _creatorSig []byte, _borrowerSig []byte, _callbackData []byte) (*types.Transaction, error) {
	return _LoanManager.Contract.SettleLend(&_LoanManager.TransactOpts, _requestData, _loanData, _cosigner, _maxCosignerCost, _cosignerData, _oracleData, _creatorSig, _borrowerSig, _callbackData)
}

// LoanManagerApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the LoanManager contract.
type LoanManagerApprovedIterator struct {
	Event *LoanManagerApproved // Event containing the contract specifics and raw log

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
func (it *LoanManagerApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerApproved)
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
		it.Event = new(LoanManagerApproved)
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
func (it *LoanManagerApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerApproved represents a Approved event raised by the LoanManager contract.
type LoanManagerApproved struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x6fb7fd1eda743aa3eb32c69f3b8cf14a5aeadf26db51057a7c5c78ba10eac8a4.
//
// Solidity: event Approved(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterApproved(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerApprovedIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "Approved", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerApprovedIterator{contract: _LoanManager.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x6fb7fd1eda743aa3eb32c69f3b8cf14a5aeadf26db51057a7c5c78ba10eac8a4.
//
// Solidity: event Approved(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *LoanManagerApproved, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "Approved", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerApproved)
				if err := _LoanManager.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x6fb7fd1eda743aa3eb32c69f3b8cf14a5aeadf26db51057a7c5c78ba10eac8a4.
//
// Solidity: event Approved(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseApproved(log types.Log) (*LoanManagerApproved, error) {
	event := new(LoanManagerApproved)
	if err := _LoanManager.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerApprovedByCallbackIterator is returned from FilterApprovedByCallback and is used to iterate over the raw logs and unpacked data for ApprovedByCallback events raised by the LoanManager contract.
type LoanManagerApprovedByCallbackIterator struct {
	Event *LoanManagerApprovedByCallback // Event containing the contract specifics and raw log

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
func (it *LoanManagerApprovedByCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerApprovedByCallback)
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
		it.Event = new(LoanManagerApprovedByCallback)
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
func (it *LoanManagerApprovedByCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerApprovedByCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerApprovedByCallback represents a ApprovedByCallback event raised by the LoanManager contract.
type LoanManagerApprovedByCallback struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApprovedByCallback is a free log retrieval operation binding the contract event 0xbcd1a053efeeac2159254030853620c82ddeffb174f76982dc0dd1083648ade3.
//
// Solidity: event ApprovedByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterApprovedByCallback(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerApprovedByCallbackIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "ApprovedByCallback", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerApprovedByCallbackIterator{contract: _LoanManager.contract, event: "ApprovedByCallback", logs: logs, sub: sub}, nil
}

// WatchApprovedByCallback is a free log subscription operation binding the contract event 0xbcd1a053efeeac2159254030853620c82ddeffb174f76982dc0dd1083648ade3.
//
// Solidity: event ApprovedByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchApprovedByCallback(opts *bind.WatchOpts, sink chan<- *LoanManagerApprovedByCallback, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "ApprovedByCallback", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerApprovedByCallback)
				if err := _LoanManager.contract.UnpackLog(event, "ApprovedByCallback", log); err != nil {
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

// ParseApprovedByCallback is a log parse operation binding the contract event 0xbcd1a053efeeac2159254030853620c82ddeffb174f76982dc0dd1083648ade3.
//
// Solidity: event ApprovedByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseApprovedByCallback(log types.Log) (*LoanManagerApprovedByCallback, error) {
	event := new(LoanManagerApprovedByCallback)
	if err := _LoanManager.contract.UnpackLog(event, "ApprovedByCallback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerApprovedBySignatureIterator is returned from FilterApprovedBySignature and is used to iterate over the raw logs and unpacked data for ApprovedBySignature events raised by the LoanManager contract.
type LoanManagerApprovedBySignatureIterator struct {
	Event *LoanManagerApprovedBySignature // Event containing the contract specifics and raw log

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
func (it *LoanManagerApprovedBySignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerApprovedBySignature)
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
		it.Event = new(LoanManagerApprovedBySignature)
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
func (it *LoanManagerApprovedBySignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerApprovedBySignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerApprovedBySignature represents a ApprovedBySignature event raised by the LoanManager contract.
type LoanManagerApprovedBySignature struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApprovedBySignature is a free log retrieval operation binding the contract event 0xac4faa318b3929891fe5ee57344cce181e72bced7246f5aa4a788ab61e7bb5d2.
//
// Solidity: event ApprovedBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterApprovedBySignature(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerApprovedBySignatureIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "ApprovedBySignature", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerApprovedBySignatureIterator{contract: _LoanManager.contract, event: "ApprovedBySignature", logs: logs, sub: sub}, nil
}

// WatchApprovedBySignature is a free log subscription operation binding the contract event 0xac4faa318b3929891fe5ee57344cce181e72bced7246f5aa4a788ab61e7bb5d2.
//
// Solidity: event ApprovedBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchApprovedBySignature(opts *bind.WatchOpts, sink chan<- *LoanManagerApprovedBySignature, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "ApprovedBySignature", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerApprovedBySignature)
				if err := _LoanManager.contract.UnpackLog(event, "ApprovedBySignature", log); err != nil {
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

// ParseApprovedBySignature is a log parse operation binding the contract event 0xac4faa318b3929891fe5ee57344cce181e72bced7246f5aa4a788ab61e7bb5d2.
//
// Solidity: event ApprovedBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseApprovedBySignature(log types.Log) (*LoanManagerApprovedBySignature, error) {
	event := new(LoanManagerApprovedBySignature)
	if err := _LoanManager.contract.UnpackLog(event, "ApprovedBySignature", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerApprovedErrorIterator is returned from FilterApprovedError and is used to iterate over the raw logs and unpacked data for ApprovedError events raised by the LoanManager contract.
type LoanManagerApprovedErrorIterator struct {
	Event *LoanManagerApprovedError // Event containing the contract specifics and raw log

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
func (it *LoanManagerApprovedErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerApprovedError)
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
		it.Event = new(LoanManagerApprovedError)
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
func (it *LoanManagerApprovedErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerApprovedErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerApprovedError represents a ApprovedError event raised by the LoanManager contract.
type LoanManagerApprovedError struct {
	Id       [32]byte
	Response [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovedError is a free log retrieval operation binding the contract event 0x04843e870d380c72a3ffb37e0afa8b10e7948a34dee4ef9c1d6754cc0699c3d3.
//
// Solidity: event ApprovedError(bytes32 indexed _id, bytes32 _response)
func (_LoanManager *LoanManagerFilterer) FilterApprovedError(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerApprovedErrorIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "ApprovedError", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerApprovedErrorIterator{contract: _LoanManager.contract, event: "ApprovedError", logs: logs, sub: sub}, nil
}

// WatchApprovedError is a free log subscription operation binding the contract event 0x04843e870d380c72a3ffb37e0afa8b10e7948a34dee4ef9c1d6754cc0699c3d3.
//
// Solidity: event ApprovedError(bytes32 indexed _id, bytes32 _response)
func (_LoanManager *LoanManagerFilterer) WatchApprovedError(opts *bind.WatchOpts, sink chan<- *LoanManagerApprovedError, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "ApprovedError", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerApprovedError)
				if err := _LoanManager.contract.UnpackLog(event, "ApprovedError", log); err != nil {
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

// ParseApprovedError is a log parse operation binding the contract event 0x04843e870d380c72a3ffb37e0afa8b10e7948a34dee4ef9c1d6754cc0699c3d3.
//
// Solidity: event ApprovedError(bytes32 indexed _id, bytes32 _response)
func (_LoanManager *LoanManagerFilterer) ParseApprovedError(log types.Log) (*LoanManagerApprovedError, error) {
	event := new(LoanManagerApprovedError)
	if err := _LoanManager.contract.UnpackLog(event, "ApprovedError", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerApprovedRejectedIterator is returned from FilterApprovedRejected and is used to iterate over the raw logs and unpacked data for ApprovedRejected events raised by the LoanManager contract.
type LoanManagerApprovedRejectedIterator struct {
	Event *LoanManagerApprovedRejected // Event containing the contract specifics and raw log

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
func (it *LoanManagerApprovedRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerApprovedRejected)
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
		it.Event = new(LoanManagerApprovedRejected)
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
func (it *LoanManagerApprovedRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerApprovedRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerApprovedRejected represents a ApprovedRejected event raised by the LoanManager contract.
type LoanManagerApprovedRejected struct {
	Id       [32]byte
	Response [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovedRejected is a free log retrieval operation binding the contract event 0x7b550df26fb29cebc893b0071262bded4c36eb15da43317bf72c3f9f5aeb554a.
//
// Solidity: event ApprovedRejected(bytes32 indexed _id, bytes32 _response)
func (_LoanManager *LoanManagerFilterer) FilterApprovedRejected(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerApprovedRejectedIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "ApprovedRejected", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerApprovedRejectedIterator{contract: _LoanManager.contract, event: "ApprovedRejected", logs: logs, sub: sub}, nil
}

// WatchApprovedRejected is a free log subscription operation binding the contract event 0x7b550df26fb29cebc893b0071262bded4c36eb15da43317bf72c3f9f5aeb554a.
//
// Solidity: event ApprovedRejected(bytes32 indexed _id, bytes32 _response)
func (_LoanManager *LoanManagerFilterer) WatchApprovedRejected(opts *bind.WatchOpts, sink chan<- *LoanManagerApprovedRejected, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "ApprovedRejected", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerApprovedRejected)
				if err := _LoanManager.contract.UnpackLog(event, "ApprovedRejected", log); err != nil {
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

// ParseApprovedRejected is a log parse operation binding the contract event 0x7b550df26fb29cebc893b0071262bded4c36eb15da43317bf72c3f9f5aeb554a.
//
// Solidity: event ApprovedRejected(bytes32 indexed _id, bytes32 _response)
func (_LoanManager *LoanManagerFilterer) ParseApprovedRejected(log types.Log) (*LoanManagerApprovedRejected, error) {
	event := new(LoanManagerApprovedRejected)
	if err := _LoanManager.contract.UnpackLog(event, "ApprovedRejected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerBorrowerByCallbackIterator is returned from FilterBorrowerByCallback and is used to iterate over the raw logs and unpacked data for BorrowerByCallback events raised by the LoanManager contract.
type LoanManagerBorrowerByCallbackIterator struct {
	Event *LoanManagerBorrowerByCallback // Event containing the contract specifics and raw log

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
func (it *LoanManagerBorrowerByCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerBorrowerByCallback)
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
		it.Event = new(LoanManagerBorrowerByCallback)
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
func (it *LoanManagerBorrowerByCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerBorrowerByCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerBorrowerByCallback represents a BorrowerByCallback event raised by the LoanManager contract.
type LoanManagerBorrowerByCallback struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBorrowerByCallback is a free log retrieval operation binding the contract event 0x168544dde877630d31a353bc62115efd38d8e072b48761461075d8421620e772.
//
// Solidity: event BorrowerByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterBorrowerByCallback(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerBorrowerByCallbackIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "BorrowerByCallback", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerBorrowerByCallbackIterator{contract: _LoanManager.contract, event: "BorrowerByCallback", logs: logs, sub: sub}, nil
}

// WatchBorrowerByCallback is a free log subscription operation binding the contract event 0x168544dde877630d31a353bc62115efd38d8e072b48761461075d8421620e772.
//
// Solidity: event BorrowerByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchBorrowerByCallback(opts *bind.WatchOpts, sink chan<- *LoanManagerBorrowerByCallback, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "BorrowerByCallback", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerBorrowerByCallback)
				if err := _LoanManager.contract.UnpackLog(event, "BorrowerByCallback", log); err != nil {
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

// ParseBorrowerByCallback is a log parse operation binding the contract event 0x168544dde877630d31a353bc62115efd38d8e072b48761461075d8421620e772.
//
// Solidity: event BorrowerByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseBorrowerByCallback(log types.Log) (*LoanManagerBorrowerByCallback, error) {
	event := new(LoanManagerBorrowerByCallback)
	if err := _LoanManager.contract.UnpackLog(event, "BorrowerByCallback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerBorrowerBySignatureIterator is returned from FilterBorrowerBySignature and is used to iterate over the raw logs and unpacked data for BorrowerBySignature events raised by the LoanManager contract.
type LoanManagerBorrowerBySignatureIterator struct {
	Event *LoanManagerBorrowerBySignature // Event containing the contract specifics and raw log

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
func (it *LoanManagerBorrowerBySignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerBorrowerBySignature)
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
		it.Event = new(LoanManagerBorrowerBySignature)
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
func (it *LoanManagerBorrowerBySignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerBorrowerBySignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerBorrowerBySignature represents a BorrowerBySignature event raised by the LoanManager contract.
type LoanManagerBorrowerBySignature struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBorrowerBySignature is a free log retrieval operation binding the contract event 0x04acecab035917c895d49033509a027f3fbb1913e978eec824926f605370653b.
//
// Solidity: event BorrowerBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterBorrowerBySignature(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerBorrowerBySignatureIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "BorrowerBySignature", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerBorrowerBySignatureIterator{contract: _LoanManager.contract, event: "BorrowerBySignature", logs: logs, sub: sub}, nil
}

// WatchBorrowerBySignature is a free log subscription operation binding the contract event 0x04acecab035917c895d49033509a027f3fbb1913e978eec824926f605370653b.
//
// Solidity: event BorrowerBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchBorrowerBySignature(opts *bind.WatchOpts, sink chan<- *LoanManagerBorrowerBySignature, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "BorrowerBySignature", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerBorrowerBySignature)
				if err := _LoanManager.contract.UnpackLog(event, "BorrowerBySignature", log); err != nil {
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

// ParseBorrowerBySignature is a log parse operation binding the contract event 0x04acecab035917c895d49033509a027f3fbb1913e978eec824926f605370653b.
//
// Solidity: event BorrowerBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseBorrowerBySignature(log types.Log) (*LoanManagerBorrowerBySignature, error) {
	event := new(LoanManagerBorrowerBySignature)
	if err := _LoanManager.contract.UnpackLog(event, "BorrowerBySignature", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the LoanManager contract.
type LoanManagerCanceledIterator struct {
	Event *LoanManagerCanceled // Event containing the contract specifics and raw log

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
func (it *LoanManagerCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerCanceled)
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
		it.Event = new(LoanManagerCanceled)
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
func (it *LoanManagerCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerCanceled represents a Canceled event raised by the LoanManager contract.
type LoanManagerCanceled struct {
	Id       [32]byte
	Canceler common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x5f32c817b1de98e681044c3d78ca791e6018e4c3b4f230e4d85bbcc3815989f3.
//
// Solidity: event Canceled(bytes32 indexed _id, address _canceler)
func (_LoanManager *LoanManagerFilterer) FilterCanceled(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerCanceledIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "Canceled", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerCanceledIterator{contract: _LoanManager.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x5f32c817b1de98e681044c3d78ca791e6018e4c3b4f230e4d85bbcc3815989f3.
//
// Solidity: event Canceled(bytes32 indexed _id, address _canceler)
func (_LoanManager *LoanManagerFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *LoanManagerCanceled, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "Canceled", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerCanceled)
				if err := _LoanManager.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0x5f32c817b1de98e681044c3d78ca791e6018e4c3b4f230e4d85bbcc3815989f3.
//
// Solidity: event Canceled(bytes32 indexed _id, address _canceler)
func (_LoanManager *LoanManagerFilterer) ParseCanceled(log types.Log) (*LoanManagerCanceled, error) {
	event := new(LoanManagerCanceled)
	if err := _LoanManager.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerCosignedIterator is returned from FilterCosigned and is used to iterate over the raw logs and unpacked data for Cosigned events raised by the LoanManager contract.
type LoanManagerCosignedIterator struct {
	Event *LoanManagerCosigned // Event containing the contract specifics and raw log

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
func (it *LoanManagerCosignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerCosigned)
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
		it.Event = new(LoanManagerCosigned)
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
func (it *LoanManagerCosignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerCosignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerCosigned represents a Cosigned event raised by the LoanManager contract.
type LoanManagerCosigned struct {
	Id       [32]byte
	Cosigner common.Address
	Cost     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosigned is a free log retrieval operation binding the contract event 0x8a4448add099432884d9833a6b44c465104441b7659c2c00b037da123a54f9d2.
//
// Solidity: event Cosigned(bytes32 indexed _id, address _cosigner, uint256 _cost)
func (_LoanManager *LoanManagerFilterer) FilterCosigned(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerCosignedIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "Cosigned", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerCosignedIterator{contract: _LoanManager.contract, event: "Cosigned", logs: logs, sub: sub}, nil
}

// WatchCosigned is a free log subscription operation binding the contract event 0x8a4448add099432884d9833a6b44c465104441b7659c2c00b037da123a54f9d2.
//
// Solidity: event Cosigned(bytes32 indexed _id, address _cosigner, uint256 _cost)
func (_LoanManager *LoanManagerFilterer) WatchCosigned(opts *bind.WatchOpts, sink chan<- *LoanManagerCosigned, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "Cosigned", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerCosigned)
				if err := _LoanManager.contract.UnpackLog(event, "Cosigned", log); err != nil {
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

// ParseCosigned is a log parse operation binding the contract event 0x8a4448add099432884d9833a6b44c465104441b7659c2c00b037da123a54f9d2.
//
// Solidity: event Cosigned(bytes32 indexed _id, address _cosigner, uint256 _cost)
func (_LoanManager *LoanManagerFilterer) ParseCosigned(log types.Log) (*LoanManagerCosigned, error) {
	event := new(LoanManagerCosigned)
	if err := _LoanManager.contract.UnpackLog(event, "Cosigned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerCreatorByCallbackIterator is returned from FilterCreatorByCallback and is used to iterate over the raw logs and unpacked data for CreatorByCallback events raised by the LoanManager contract.
type LoanManagerCreatorByCallbackIterator struct {
	Event *LoanManagerCreatorByCallback // Event containing the contract specifics and raw log

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
func (it *LoanManagerCreatorByCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerCreatorByCallback)
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
		it.Event = new(LoanManagerCreatorByCallback)
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
func (it *LoanManagerCreatorByCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerCreatorByCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerCreatorByCallback represents a CreatorByCallback event raised by the LoanManager contract.
type LoanManagerCreatorByCallback struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreatorByCallback is a free log retrieval operation binding the contract event 0x025f21426e185fab3bf170a9374c750b17a510c3ed51549fa4ff068dbca66057.
//
// Solidity: event CreatorByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterCreatorByCallback(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerCreatorByCallbackIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "CreatorByCallback", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerCreatorByCallbackIterator{contract: _LoanManager.contract, event: "CreatorByCallback", logs: logs, sub: sub}, nil
}

// WatchCreatorByCallback is a free log subscription operation binding the contract event 0x025f21426e185fab3bf170a9374c750b17a510c3ed51549fa4ff068dbca66057.
//
// Solidity: event CreatorByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchCreatorByCallback(opts *bind.WatchOpts, sink chan<- *LoanManagerCreatorByCallback, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "CreatorByCallback", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerCreatorByCallback)
				if err := _LoanManager.contract.UnpackLog(event, "CreatorByCallback", log); err != nil {
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

// ParseCreatorByCallback is a log parse operation binding the contract event 0x025f21426e185fab3bf170a9374c750b17a510c3ed51549fa4ff068dbca66057.
//
// Solidity: event CreatorByCallback(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseCreatorByCallback(log types.Log) (*LoanManagerCreatorByCallback, error) {
	event := new(LoanManagerCreatorByCallback)
	if err := _LoanManager.contract.UnpackLog(event, "CreatorByCallback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerCreatorBySignatureIterator is returned from FilterCreatorBySignature and is used to iterate over the raw logs and unpacked data for CreatorBySignature events raised by the LoanManager contract.
type LoanManagerCreatorBySignatureIterator struct {
	Event *LoanManagerCreatorBySignature // Event containing the contract specifics and raw log

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
func (it *LoanManagerCreatorBySignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerCreatorBySignature)
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
		it.Event = new(LoanManagerCreatorBySignature)
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
func (it *LoanManagerCreatorBySignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerCreatorBySignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerCreatorBySignature represents a CreatorBySignature event raised by the LoanManager contract.
type LoanManagerCreatorBySignature struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreatorBySignature is a free log retrieval operation binding the contract event 0x3971d1c591487c96ba2561ab5f74f00001a6134408897ab8b0124c1673055dd2.
//
// Solidity: event CreatorBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) FilterCreatorBySignature(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerCreatorBySignatureIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "CreatorBySignature", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerCreatorBySignatureIterator{contract: _LoanManager.contract, event: "CreatorBySignature", logs: logs, sub: sub}, nil
}

// WatchCreatorBySignature is a free log subscription operation binding the contract event 0x3971d1c591487c96ba2561ab5f74f00001a6134408897ab8b0124c1673055dd2.
//
// Solidity: event CreatorBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) WatchCreatorBySignature(opts *bind.WatchOpts, sink chan<- *LoanManagerCreatorBySignature, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "CreatorBySignature", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerCreatorBySignature)
				if err := _LoanManager.contract.UnpackLog(event, "CreatorBySignature", log); err != nil {
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

// ParseCreatorBySignature is a log parse operation binding the contract event 0x3971d1c591487c96ba2561ab5f74f00001a6134408897ab8b0124c1673055dd2.
//
// Solidity: event CreatorBySignature(bytes32 indexed _id)
func (_LoanManager *LoanManagerFilterer) ParseCreatorBySignature(log types.Log) (*LoanManagerCreatorBySignature, error) {
	event := new(LoanManagerCreatorBySignature)
	if err := _LoanManager.contract.UnpackLog(event, "CreatorBySignature", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerLentIterator is returned from FilterLent and is used to iterate over the raw logs and unpacked data for Lent events raised by the LoanManager contract.
type LoanManagerLentIterator struct {
	Event *LoanManagerLent // Event containing the contract specifics and raw log

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
func (it *LoanManagerLentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerLent)
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
		it.Event = new(LoanManagerLent)
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
func (it *LoanManagerLentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerLentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerLent represents a Lent event raised by the LoanManager contract.
type LoanManagerLent struct {
	Id     [32]byte
	Lender common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLent is a free log retrieval operation binding the contract event 0xf701a0be7203a591a078aa39b506b74dcf910d7a73b1f73a4dff4d0df3968361.
//
// Solidity: event Lent(bytes32 indexed _id, address _lender, uint256 _tokens)
func (_LoanManager *LoanManagerFilterer) FilterLent(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerLentIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "Lent", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerLentIterator{contract: _LoanManager.contract, event: "Lent", logs: logs, sub: sub}, nil
}

// WatchLent is a free log subscription operation binding the contract event 0xf701a0be7203a591a078aa39b506b74dcf910d7a73b1f73a4dff4d0df3968361.
//
// Solidity: event Lent(bytes32 indexed _id, address _lender, uint256 _tokens)
func (_LoanManager *LoanManagerFilterer) WatchLent(opts *bind.WatchOpts, sink chan<- *LoanManagerLent, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "Lent", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerLent)
				if err := _LoanManager.contract.UnpackLog(event, "Lent", log); err != nil {
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

// ParseLent is a log parse operation binding the contract event 0xf701a0be7203a591a078aa39b506b74dcf910d7a73b1f73a4dff4d0df3968361.
//
// Solidity: event Lent(bytes32 indexed _id, address _lender, uint256 _tokens)
func (_LoanManager *LoanManagerFilterer) ParseLent(log types.Log) (*LoanManagerLent, error) {
	event := new(LoanManagerLent)
	if err := _LoanManager.contract.UnpackLog(event, "Lent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerReadedOracleIterator is returned from FilterReadedOracle and is used to iterate over the raw logs and unpacked data for ReadedOracle events raised by the LoanManager contract.
type LoanManagerReadedOracleIterator struct {
	Event *LoanManagerReadedOracle // Event containing the contract specifics and raw log

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
func (it *LoanManagerReadedOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerReadedOracle)
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
		it.Event = new(LoanManagerReadedOracle)
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
func (it *LoanManagerReadedOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerReadedOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerReadedOracle represents a ReadedOracle event raised by the LoanManager contract.
type LoanManagerReadedOracle struct {
	Oracle     common.Address
	Tokens     *big.Int
	Equivalent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReadedOracle is a free log retrieval operation binding the contract event 0x0b53fa0f90e4773f861d84b5da681ae7b3dd6b29358fdb4a38ba8dec13fc80f3.
//
// Solidity: event ReadedOracle(address _oracle, uint256 _tokens, uint256 _equivalent)
func (_LoanManager *LoanManagerFilterer) FilterReadedOracle(opts *bind.FilterOpts) (*LoanManagerReadedOracleIterator, error) {

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "ReadedOracle")
	if err != nil {
		return nil, err
	}
	return &LoanManagerReadedOracleIterator{contract: _LoanManager.contract, event: "ReadedOracle", logs: logs, sub: sub}, nil
}

// WatchReadedOracle is a free log subscription operation binding the contract event 0x0b53fa0f90e4773f861d84b5da681ae7b3dd6b29358fdb4a38ba8dec13fc80f3.
//
// Solidity: event ReadedOracle(address _oracle, uint256 _tokens, uint256 _equivalent)
func (_LoanManager *LoanManagerFilterer) WatchReadedOracle(opts *bind.WatchOpts, sink chan<- *LoanManagerReadedOracle) (event.Subscription, error) {

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "ReadedOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerReadedOracle)
				if err := _LoanManager.contract.UnpackLog(event, "ReadedOracle", log); err != nil {
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

// ParseReadedOracle is a log parse operation binding the contract event 0x0b53fa0f90e4773f861d84b5da681ae7b3dd6b29358fdb4a38ba8dec13fc80f3.
//
// Solidity: event ReadedOracle(address _oracle, uint256 _tokens, uint256 _equivalent)
func (_LoanManager *LoanManagerFilterer) ParseReadedOracle(log types.Log) (*LoanManagerReadedOracle, error) {
	event := new(LoanManagerReadedOracle)
	if err := _LoanManager.contract.UnpackLog(event, "ReadedOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerRequestedIterator is returned from FilterRequested and is used to iterate over the raw logs and unpacked data for Requested events raised by the LoanManager contract.
type LoanManagerRequestedIterator struct {
	Event *LoanManagerRequested // Event containing the contract specifics and raw log

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
func (it *LoanManagerRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerRequested)
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
		it.Event = new(LoanManagerRequested)
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
func (it *LoanManagerRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerRequested represents a Requested event raised by the LoanManager contract.
type LoanManagerRequested struct {
	Id         [32]byte
	Amount     *big.Int
	Model      common.Address
	Creator    common.Address
	Oracle     common.Address
	Borrower   common.Address
	Callback   common.Address
	Salt       *big.Int
	LoanData   []byte
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequested is a free log retrieval operation binding the contract event 0x200404915876dec9fd44568ed7f5358a7fc4b212c1cdfc404ab17be21a1cf4af.
//
// Solidity: event Requested(bytes32 indexed _id, uint128 _amount, address _model, address _creator, address _oracle, address _borrower, address _callback, uint256 _salt, bytes _loanData, uint256 _expiration)
func (_LoanManager *LoanManagerFilterer) FilterRequested(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerRequestedIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "Requested", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerRequestedIterator{contract: _LoanManager.contract, event: "Requested", logs: logs, sub: sub}, nil
}

// WatchRequested is a free log subscription operation binding the contract event 0x200404915876dec9fd44568ed7f5358a7fc4b212c1cdfc404ab17be21a1cf4af.
//
// Solidity: event Requested(bytes32 indexed _id, uint128 _amount, address _model, address _creator, address _oracle, address _borrower, address _callback, uint256 _salt, bytes _loanData, uint256 _expiration)
func (_LoanManager *LoanManagerFilterer) WatchRequested(opts *bind.WatchOpts, sink chan<- *LoanManagerRequested, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "Requested", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerRequested)
				if err := _LoanManager.contract.UnpackLog(event, "Requested", log); err != nil {
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

// ParseRequested is a log parse operation binding the contract event 0x200404915876dec9fd44568ed7f5358a7fc4b212c1cdfc404ab17be21a1cf4af.
//
// Solidity: event Requested(bytes32 indexed _id, uint128 _amount, address _model, address _creator, address _oracle, address _borrower, address _callback, uint256 _salt, bytes _loanData, uint256 _expiration)
func (_LoanManager *LoanManagerFilterer) ParseRequested(log types.Log) (*LoanManagerRequested, error) {
	event := new(LoanManagerRequested)
	if err := _LoanManager.contract.UnpackLog(event, "Requested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerSettledCancelIterator is returned from FilterSettledCancel and is used to iterate over the raw logs and unpacked data for SettledCancel events raised by the LoanManager contract.
type LoanManagerSettledCancelIterator struct {
	Event *LoanManagerSettledCancel // Event containing the contract specifics and raw log

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
func (it *LoanManagerSettledCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerSettledCancel)
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
		it.Event = new(LoanManagerSettledCancel)
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
func (it *LoanManagerSettledCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerSettledCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerSettledCancel represents a SettledCancel event raised by the LoanManager contract.
type LoanManagerSettledCancel struct {
	Id       [32]byte
	Canceler common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSettledCancel is a free log retrieval operation binding the contract event 0x029e3e789ef1a202e83e668d22e38205106d93105847601f49983c21efb8cfbe.
//
// Solidity: event SettledCancel(bytes32 indexed _id, address _canceler)
func (_LoanManager *LoanManagerFilterer) FilterSettledCancel(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerSettledCancelIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "SettledCancel", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerSettledCancelIterator{contract: _LoanManager.contract, event: "SettledCancel", logs: logs, sub: sub}, nil
}

// WatchSettledCancel is a free log subscription operation binding the contract event 0x029e3e789ef1a202e83e668d22e38205106d93105847601f49983c21efb8cfbe.
//
// Solidity: event SettledCancel(bytes32 indexed _id, address _canceler)
func (_LoanManager *LoanManagerFilterer) WatchSettledCancel(opts *bind.WatchOpts, sink chan<- *LoanManagerSettledCancel, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "SettledCancel", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerSettledCancel)
				if err := _LoanManager.contract.UnpackLog(event, "SettledCancel", log); err != nil {
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

// ParseSettledCancel is a log parse operation binding the contract event 0x029e3e789ef1a202e83e668d22e38205106d93105847601f49983c21efb8cfbe.
//
// Solidity: event SettledCancel(bytes32 indexed _id, address _canceler)
func (_LoanManager *LoanManagerFilterer) ParseSettledCancel(log types.Log) (*LoanManagerSettledCancel, error) {
	event := new(LoanManagerSettledCancel)
	if err := _LoanManager.contract.UnpackLog(event, "SettledCancel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanManagerSettledLendIterator is returned from FilterSettledLend and is used to iterate over the raw logs and unpacked data for SettledLend events raised by the LoanManager contract.
type LoanManagerSettledLendIterator struct {
	Event *LoanManagerSettledLend // Event containing the contract specifics and raw log

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
func (it *LoanManagerSettledLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanManagerSettledLend)
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
		it.Event = new(LoanManagerSettledLend)
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
func (it *LoanManagerSettledLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanManagerSettledLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanManagerSettledLend represents a SettledLend event raised by the LoanManager contract.
type LoanManagerSettledLend struct {
	Id     [32]byte
	Lender common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSettledLend is a free log retrieval operation binding the contract event 0xb1d906969e9527a5ad3ac2ca84eb415890ddf44581528edf18d0e7154d7f126d.
//
// Solidity: event SettledLend(bytes32 indexed _id, address _lender, uint256 _tokens)
func (_LoanManager *LoanManagerFilterer) FilterSettledLend(opts *bind.FilterOpts, _id [][32]byte) (*LoanManagerSettledLendIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.FilterLogs(opts, "SettledLend", _idRule)
	if err != nil {
		return nil, err
	}
	return &LoanManagerSettledLendIterator{contract: _LoanManager.contract, event: "SettledLend", logs: logs, sub: sub}, nil
}

// WatchSettledLend is a free log subscription operation binding the contract event 0xb1d906969e9527a5ad3ac2ca84eb415890ddf44581528edf18d0e7154d7f126d.
//
// Solidity: event SettledLend(bytes32 indexed _id, address _lender, uint256 _tokens)
func (_LoanManager *LoanManagerFilterer) WatchSettledLend(opts *bind.WatchOpts, sink chan<- *LoanManagerSettledLend, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _LoanManager.contract.WatchLogs(opts, "SettledLend", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanManagerSettledLend)
				if err := _LoanManager.contract.UnpackLog(event, "SettledLend", log); err != nil {
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

// ParseSettledLend is a log parse operation binding the contract event 0xb1d906969e9527a5ad3ac2ca84eb415890ddf44581528edf18d0e7154d7f126d.
//
// Solidity: event SettledLend(bytes32 indexed _id, address _lender, uint256 _tokens)
func (_LoanManager *LoanManagerFilterer) ParseSettledLend(log types.Log) (*LoanManagerSettledLend, error) {
	event := new(LoanManagerSettledLend)
	if err := _LoanManager.contract.UnpackLog(event, "SettledLend", log); err != nil {
		return nil, err
	}
	return event, nil
}
