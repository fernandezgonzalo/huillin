// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Collateral

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

// CollateralABI is the input ABI used to generate the binding from.
const CollateralABI = "[{\"inputs\":[{\"internalType\":\"contractLoanManager\",\"name\":\"_loanManager\",\"type\":\"address\"},{\"internalType\":\"contractCollateralAuction\",\"name\":\"_auction\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractCollateralHandler\",\"name\":\"_handler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newAmount\",\"type\":\"uint256\"}],\"name\":\"BorrowCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dueTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_obligation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_obligationTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_marketValue\",\"type\":\"uint256\"}],\"name\":\"ClaimedExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_marketValue\",\"type\":\"uint256\"}],\"name\":\"ClaimedLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_received\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_leftover\",\"type\":\"uint256\"}],\"name\":\"ClosedAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_debtId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractRateOracle\",\"name\":\"_oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_liquidationRatio\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_balanceRatio\",\"type\":\"uint96\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_uriProvider\",\"type\":\"address\"}],\"name\":\"SetURIProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"SetUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"}],\"name\":\"Started\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"assetsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"contractCollateralAuction\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_leftover\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_received\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"auctionClosed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionToEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"internalType\":\"contractCollateralHandler\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"borrowCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_debtId\",\"type\":\"bytes32\"},{\"internalType\":\"contractRateOracle\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_liquidationRatio\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_balanceRatio\",\"type\":\"uint96\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"entryId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"debtToEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"debtId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractRateOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"liquidationRatio\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"balanceRatio\",\"type\":\"uint96\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entryToAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEntriesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"}],\"name\":\"inAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetHolder\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanManager\",\"outputs\":[{\"internalType\":\"contractLoanManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanManagerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"requestCosign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_userData\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_authorized\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Collateral is an auto generated Go binding around an Ethereum contract.
type Collateral struct {
	CollateralCaller     // Read-only binding to the contract
	CollateralTransactor // Write-only binding to the contract
	CollateralFilterer   // Log filterer for contract events
}

// CollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralSession struct {
	Contract     *Collateral       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralCallerSession struct {
	Contract *CollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralTransactorSession struct {
	Contract     *CollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralRaw struct {
	Contract *Collateral // Generic contract binding to access the raw methods on
}

// CollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralCallerRaw struct {
	Contract *CollateralCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralTransactorRaw struct {
	Contract *CollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateral creates a new instance of Collateral, bound to a specific deployed contract.
func NewCollateral(address common.Address, backend bind.ContractBackend) (*Collateral, error) {
	contract, err := bindCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Collateral{CollateralCaller: CollateralCaller{contract: contract}, CollateralTransactor: CollateralTransactor{contract: contract}, CollateralFilterer: CollateralFilterer{contract: contract}}, nil
}

// NewCollateralCaller creates a new read-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralCaller(address common.Address, caller bind.ContractCaller) (*CollateralCaller, error) {
	contract, err := bindCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralCaller{contract: contract}, nil
}

// NewCollateralTransactor creates a new write-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralTransactor, error) {
	contract, err := bindCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralTransactor{contract: contract}, nil
}

// NewCollateralFilterer creates a new log filterer instance of Collateral, bound to a specific deployed contract.
func NewCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralFilterer, error) {
	contract, err := bindCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralFilterer{contract: contract}, nil
}

// bindCollateral binds a generic wrapper to an already deployed contract.
func bindCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollateralABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.CollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(uint256)
func (_Collateral *CollateralCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(uint256)
func (_Collateral *CollateralSession) VERSION() (*big.Int, error) {
	return _Collateral.Contract.VERSION(&_Collateral.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(uint256)
func (_Collateral *CollateralCallerSession) VERSION() (*big.Int, error) {
	return _Collateral.Contract.VERSION(&_Collateral.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() constant returns(uint256[])
func (_Collateral *CollateralCaller) AllTokens(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "allTokens")
	return *ret0, err
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() constant returns(uint256[])
func (_Collateral *CollateralSession) AllTokens() ([]*big.Int, error) {
	return _Collateral.Contract.AllTokens(&_Collateral.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() constant returns(uint256[])
func (_Collateral *CollateralCallerSession) AllTokens() ([]*big.Int, error) {
	return _Collateral.Contract.AllTokens(&_Collateral.CallOpts)
}

// AssetsOf is a free data retrieval call binding the contract method 0x2c62fa10.
//
// Solidity: function assetsOf(address _owner) constant returns(uint256[])
func (_Collateral *CollateralCaller) AssetsOf(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "assetsOf", _owner)
	return *ret0, err
}

// AssetsOf is a free data retrieval call binding the contract method 0x2c62fa10.
//
// Solidity: function assetsOf(address _owner) constant returns(uint256[])
func (_Collateral *CollateralSession) AssetsOf(_owner common.Address) ([]*big.Int, error) {
	return _Collateral.Contract.AssetsOf(&_Collateral.CallOpts, _owner)
}

// AssetsOf is a free data retrieval call binding the contract method 0x2c62fa10.
//
// Solidity: function assetsOf(address _owner) constant returns(uint256[])
func (_Collateral *CollateralCallerSession) AssetsOf(_owner common.Address) ([]*big.Int, error) {
	return _Collateral.Contract.AssetsOf(&_Collateral.CallOpts, _owner)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() constant returns(address)
func (_Collateral *CollateralCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "auction")
	return *ret0, err
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() constant returns(address)
func (_Collateral *CollateralSession) Auction() (common.Address, error) {
	return _Collateral.Contract.Auction(&_Collateral.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() constant returns(address)
func (_Collateral *CollateralCallerSession) Auction() (common.Address, error) {
	return _Collateral.Contract.Auction(&_Collateral.CallOpts)
}

// AuctionToEntry is a free data retrieval call binding the contract method 0x51f7b9bd.
//
// Solidity: function auctionToEntry(uint256 ) constant returns(uint256)
func (_Collateral *CollateralCaller) AuctionToEntry(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "auctionToEntry", arg0)
	return *ret0, err
}

// AuctionToEntry is a free data retrieval call binding the contract method 0x51f7b9bd.
//
// Solidity: function auctionToEntry(uint256 ) constant returns(uint256)
func (_Collateral *CollateralSession) AuctionToEntry(arg0 *big.Int) (*big.Int, error) {
	return _Collateral.Contract.AuctionToEntry(&_Collateral.CallOpts, arg0)
}

// AuctionToEntry is a free data retrieval call binding the contract method 0x51f7b9bd.
//
// Solidity: function auctionToEntry(uint256 ) constant returns(uint256)
func (_Collateral *CollateralCallerSession) AuctionToEntry(arg0 *big.Int) (*big.Int, error) {
	return _Collateral.Contract.AuctionToEntry(&_Collateral.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Collateral *CollateralCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Collateral *CollateralSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Collateral.Contract.BalanceOf(&_Collateral.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Collateral *CollateralCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Collateral.Contract.BalanceOf(&_Collateral.CallOpts, _owner)
}

// Cost is a free data retrieval call binding the contract method 0x477ce277.
//
// Solidity: function cost(address , uint256 , bytes , bytes ) constant returns(uint256)
func (_Collateral *CollateralCaller) Cost(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte, arg3 []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "cost", arg0, arg1, arg2, arg3)
	return *ret0, err
}

// Cost is a free data retrieval call binding the contract method 0x477ce277.
//
// Solidity: function cost(address , uint256 , bytes , bytes ) constant returns(uint256)
func (_Collateral *CollateralSession) Cost(arg0 common.Address, arg1 *big.Int, arg2 []byte, arg3 []byte) (*big.Int, error) {
	return _Collateral.Contract.Cost(&_Collateral.CallOpts, arg0, arg1, arg2, arg3)
}

// Cost is a free data retrieval call binding the contract method 0x477ce277.
//
// Solidity: function cost(address , uint256 , bytes , bytes ) constant returns(uint256)
func (_Collateral *CollateralCallerSession) Cost(arg0 common.Address, arg1 *big.Int, arg2 []byte, arg3 []byte) (*big.Int, error) {
	return _Collateral.Contract.Cost(&_Collateral.CallOpts, arg0, arg1, arg2, arg3)
}

// DebtToEntry is a free data retrieval call binding the contract method 0x48eab007.
//
// Solidity: function debtToEntry(bytes32 ) constant returns(uint256)
func (_Collateral *CollateralCaller) DebtToEntry(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "debtToEntry", arg0)
	return *ret0, err
}

// DebtToEntry is a free data retrieval call binding the contract method 0x48eab007.
//
// Solidity: function debtToEntry(bytes32 ) constant returns(uint256)
func (_Collateral *CollateralSession) DebtToEntry(arg0 [32]byte) (*big.Int, error) {
	return _Collateral.Contract.DebtToEntry(&_Collateral.CallOpts, arg0)
}

// DebtToEntry is a free data retrieval call binding the contract method 0x48eab007.
//
// Solidity: function debtToEntry(bytes32 ) constant returns(uint256)
func (_Collateral *CollateralCallerSession) DebtToEntry(arg0 [32]byte) (*big.Int, error) {
	return _Collateral.Contract.DebtToEntry(&_Collateral.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) constant returns(bytes32 debtId, uint256 amount, address oracle, address token, uint96 liquidationRatio, uint96 balanceRatio)
func (_Collateral *CollateralCaller) Entries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DebtId           [32]byte
	Amount           *big.Int
	Oracle           common.Address
	Token            common.Address
	LiquidationRatio *big.Int
	BalanceRatio     *big.Int
}, error) {
	ret := new(struct {
		DebtId           [32]byte
		Amount           *big.Int
		Oracle           common.Address
		Token            common.Address
		LiquidationRatio *big.Int
		BalanceRatio     *big.Int
	})
	out := ret
	err := _Collateral.contract.Call(opts, out, "entries", arg0)
	return *ret, err
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) constant returns(bytes32 debtId, uint256 amount, address oracle, address token, uint96 liquidationRatio, uint96 balanceRatio)
func (_Collateral *CollateralSession) Entries(arg0 *big.Int) (struct {
	DebtId           [32]byte
	Amount           *big.Int
	Oracle           common.Address
	Token            common.Address
	LiquidationRatio *big.Int
	BalanceRatio     *big.Int
}, error) {
	return _Collateral.Contract.Entries(&_Collateral.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) constant returns(bytes32 debtId, uint256 amount, address oracle, address token, uint96 liquidationRatio, uint96 balanceRatio)
func (_Collateral *CollateralCallerSession) Entries(arg0 *big.Int) (struct {
	DebtId           [32]byte
	Amount           *big.Int
	Oracle           common.Address
	Token            common.Address
	LiquidationRatio *big.Int
	BalanceRatio     *big.Int
}, error) {
	return _Collateral.Contract.Entries(&_Collateral.CallOpts, arg0)
}

// EntryToAuction is a free data retrieval call binding the contract method 0xf589671b.
//
// Solidity: function entryToAuction(uint256 ) constant returns(uint256)
func (_Collateral *CollateralCaller) EntryToAuction(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "entryToAuction", arg0)
	return *ret0, err
}

// EntryToAuction is a free data retrieval call binding the contract method 0xf589671b.
//
// Solidity: function entryToAuction(uint256 ) constant returns(uint256)
func (_Collateral *CollateralSession) EntryToAuction(arg0 *big.Int) (*big.Int, error) {
	return _Collateral.Contract.EntryToAuction(&_Collateral.CallOpts, arg0)
}

// EntryToAuction is a free data retrieval call binding the contract method 0xf589671b.
//
// Solidity: function entryToAuction(uint256 ) constant returns(uint256)
func (_Collateral *CollateralCallerSession) EntryToAuction(arg0 *big.Int) (*big.Int, error) {
	return _Collateral.Contract.EntryToAuction(&_Collateral.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _assetId) constant returns(address)
func (_Collateral *CollateralCaller) GetApproved(opts *bind.CallOpts, _assetId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "getApproved", _assetId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _assetId) constant returns(address)
func (_Collateral *CollateralSession) GetApproved(_assetId *big.Int) (common.Address, error) {
	return _Collateral.Contract.GetApproved(&_Collateral.CallOpts, _assetId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _assetId) constant returns(address)
func (_Collateral *CollateralCallerSession) GetApproved(_assetId *big.Int) (common.Address, error) {
	return _Collateral.Contract.GetApproved(&_Collateral.CallOpts, _assetId)
}

// GetEntriesLength is a free data retrieval call binding the contract method 0xa0d10c9a.
//
// Solidity: function getEntriesLength() constant returns(uint256)
func (_Collateral *CollateralCaller) GetEntriesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "getEntriesLength")
	return *ret0, err
}

// GetEntriesLength is a free data retrieval call binding the contract method 0xa0d10c9a.
//
// Solidity: function getEntriesLength() constant returns(uint256)
func (_Collateral *CollateralSession) GetEntriesLength() (*big.Int, error) {
	return _Collateral.Contract.GetEntriesLength(&_Collateral.CallOpts)
}

// GetEntriesLength is a free data retrieval call binding the contract method 0xa0d10c9a.
//
// Solidity: function getEntriesLength() constant returns(uint256)
func (_Collateral *CollateralCallerSession) GetEntriesLength() (*big.Int, error) {
	return _Collateral.Contract.GetEntriesLength(&_Collateral.CallOpts)
}

// InAuction is a free data retrieval call binding the contract method 0x392c56ae.
//
// Solidity: function inAuction(uint256 _entryId) constant returns(bool)
func (_Collateral *CollateralCaller) InAuction(opts *bind.CallOpts, _entryId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "inAuction", _entryId)
	return *ret0, err
}

// InAuction is a free data retrieval call binding the contract method 0x392c56ae.
//
// Solidity: function inAuction(uint256 _entryId) constant returns(bool)
func (_Collateral *CollateralSession) InAuction(_entryId *big.Int) (bool, error) {
	return _Collateral.Contract.InAuction(&_Collateral.CallOpts, _entryId)
}

// InAuction is a free data retrieval call binding the contract method 0x392c56ae.
//
// Solidity: function inAuction(uint256 _entryId) constant returns(bool)
func (_Collateral *CollateralCallerSession) InAuction(_entryId *big.Int) (bool, error) {
	return _Collateral.Contract.InAuction(&_Collateral.CallOpts, _entryId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _operator, address _assetHolder) constant returns(bool)
func (_Collateral *CollateralCaller) IsApprovedForAll(opts *bind.CallOpts, _operator common.Address, _assetHolder common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "isApprovedForAll", _operator, _assetHolder)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _operator, address _assetHolder) constant returns(bool)
func (_Collateral *CollateralSession) IsApprovedForAll(_operator common.Address, _assetHolder common.Address) (bool, error) {
	return _Collateral.Contract.IsApprovedForAll(&_Collateral.CallOpts, _operator, _assetHolder)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _operator, address _assetHolder) constant returns(bool)
func (_Collateral *CollateralCallerSession) IsApprovedForAll(_operator common.Address, _assetHolder common.Address) (bool, error) {
	return _Collateral.Contract.IsApprovedForAll(&_Collateral.CallOpts, _operator, _assetHolder)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x2972b0f0.
//
// Solidity: function isAuthorized(address _operator, uint256 _assetId) constant returns(bool)
func (_Collateral *CollateralCaller) IsAuthorized(opts *bind.CallOpts, _operator common.Address, _assetId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "isAuthorized", _operator, _assetId)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x2972b0f0.
//
// Solidity: function isAuthorized(address _operator, uint256 _assetId) constant returns(bool)
func (_Collateral *CollateralSession) IsAuthorized(_operator common.Address, _assetId *big.Int) (bool, error) {
	return _Collateral.Contract.IsAuthorized(&_Collateral.CallOpts, _operator, _assetId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x2972b0f0.
//
// Solidity: function isAuthorized(address _operator, uint256 _assetId) constant returns(bool)
func (_Collateral *CollateralCallerSession) IsAuthorized(_operator common.Address, _assetId *big.Int) (bool, error) {
	return _Collateral.Contract.IsAuthorized(&_Collateral.CallOpts, _operator, _assetId)
}

// LoanManager is a free data retrieval call binding the contract method 0xb3ccbcfe.
//
// Solidity: function loanManager() constant returns(address)
func (_Collateral *CollateralCaller) LoanManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "loanManager")
	return *ret0, err
}

// LoanManager is a free data retrieval call binding the contract method 0xb3ccbcfe.
//
// Solidity: function loanManager() constant returns(address)
func (_Collateral *CollateralSession) LoanManager() (common.Address, error) {
	return _Collateral.Contract.LoanManager(&_Collateral.CallOpts)
}

// LoanManager is a free data retrieval call binding the contract method 0xb3ccbcfe.
//
// Solidity: function loanManager() constant returns(address)
func (_Collateral *CollateralCallerSession) LoanManager() (common.Address, error) {
	return _Collateral.Contract.LoanManager(&_Collateral.CallOpts)
}

// LoanManagerToken is a free data retrieval call binding the contract method 0x283012c4.
//
// Solidity: function loanManagerToken() constant returns(address)
func (_Collateral *CollateralCaller) LoanManagerToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "loanManagerToken")
	return *ret0, err
}

// LoanManagerToken is a free data retrieval call binding the contract method 0x283012c4.
//
// Solidity: function loanManagerToken() constant returns(address)
func (_Collateral *CollateralSession) LoanManagerToken() (common.Address, error) {
	return _Collateral.Contract.LoanManagerToken(&_Collateral.CallOpts)
}

// LoanManagerToken is a free data retrieval call binding the contract method 0x283012c4.
//
// Solidity: function loanManagerToken() constant returns(address)
func (_Collateral *CollateralCallerSession) LoanManagerToken() (common.Address, error) {
	return _Collateral.Contract.LoanManagerToken(&_Collateral.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Collateral *CollateralCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Collateral *CollateralSession) Name() (string, error) {
	return _Collateral.Contract.Name(&_Collateral.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Collateral *CollateralCallerSession) Name() (string, error) {
	return _Collateral.Contract.Name(&_Collateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Collateral *CollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Collateral *CollateralSession) Owner() (common.Address, error) {
	return _Collateral.Contract.Owner(&_Collateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Collateral *CollateralCallerSession) Owner() (common.Address, error) {
	return _Collateral.Contract.Owner(&_Collateral.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _assetId) constant returns(address)
func (_Collateral *CollateralCaller) OwnerOf(opts *bind.CallOpts, _assetId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "ownerOf", _assetId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _assetId) constant returns(address)
func (_Collateral *CollateralSession) OwnerOf(_assetId *big.Int) (common.Address, error) {
	return _Collateral.Contract.OwnerOf(&_Collateral.CallOpts, _assetId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _assetId) constant returns(address)
func (_Collateral *CollateralCallerSession) OwnerOf(_assetId *big.Int) (common.Address, error) {
	return _Collateral.Contract.OwnerOf(&_Collateral.CallOpts, _assetId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Collateral *CollateralCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Collateral *CollateralSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Collateral.Contract.SupportsInterface(&_Collateral.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Collateral *CollateralCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Collateral.Contract.SupportsInterface(&_Collateral.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Collateral *CollateralCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Collateral *CollateralSession) Symbol() (string, error) {
	return _Collateral.Contract.Symbol(&_Collateral.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Collateral *CollateralCallerSession) Symbol() (string, error) {
	return _Collateral.Contract.Symbol(&_Collateral.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_Collateral *CollateralCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_Collateral *CollateralSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Collateral.Contract.TokenByIndex(&_Collateral.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_Collateral *CollateralCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Collateral.Contract.TokenByIndex(&_Collateral.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_Collateral *CollateralCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_Collateral *CollateralSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Collateral.Contract.TokenOfOwnerByIndex(&_Collateral.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_Collateral *CollateralCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Collateral.Contract.TokenOfOwnerByIndex(&_Collateral.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_Collateral *CollateralCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_Collateral *CollateralSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Collateral.Contract.TokenURI(&_Collateral.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_Collateral *CollateralCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Collateral.Contract.TokenURI(&_Collateral.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Collateral *CollateralCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Collateral *CollateralSession) TotalSupply() (*big.Int, error) {
	return _Collateral.Contract.TotalSupply(&_Collateral.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Collateral *CollateralCallerSession) TotalSupply() (*big.Int, error) {
	return _Collateral.Contract.TotalSupply(&_Collateral.CallOpts)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() constant returns(string)
func (_Collateral *CollateralCaller) Url(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Collateral.contract.Call(opts, out, "url")
	return *ret0, err
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() constant returns(string)
func (_Collateral *CollateralSession) Url() (string, error) {
	return _Collateral.Contract.Url(&_Collateral.CallOpts)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() constant returns(string)
func (_Collateral *CollateralCallerSession) Url() (string, error) {
	return _Collateral.Contract.Url(&_Collateral.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _operator, uint256 _assetId) returns()
func (_Collateral *CollateralTransactor) Approve(opts *bind.TransactOpts, _operator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "approve", _operator, _assetId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _operator, uint256 _assetId) returns()
func (_Collateral *CollateralSession) Approve(_operator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Approve(&_Collateral.TransactOpts, _operator, _assetId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _operator, uint256 _assetId) returns()
func (_Collateral *CollateralTransactorSession) Approve(_operator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Approve(&_Collateral.TransactOpts, _operator, _assetId)
}

// AuctionClosed is a paid mutator transaction binding the contract method 0xd80f6c48.
//
// Solidity: function auctionClosed(uint256 _id, uint256 _leftover, uint256 _received, bytes _data) returns()
func (_Collateral *CollateralTransactor) AuctionClosed(opts *bind.TransactOpts, _id *big.Int, _leftover *big.Int, _received *big.Int, _data []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "auctionClosed", _id, _leftover, _received, _data)
}

// AuctionClosed is a paid mutator transaction binding the contract method 0xd80f6c48.
//
// Solidity: function auctionClosed(uint256 _id, uint256 _leftover, uint256 _received, bytes _data) returns()
func (_Collateral *CollateralSession) AuctionClosed(_id *big.Int, _leftover *big.Int, _received *big.Int, _data []byte) (*types.Transaction, error) {
	return _Collateral.Contract.AuctionClosed(&_Collateral.TransactOpts, _id, _leftover, _received, _data)
}

// AuctionClosed is a paid mutator transaction binding the contract method 0xd80f6c48.
//
// Solidity: function auctionClosed(uint256 _id, uint256 _leftover, uint256 _received, bytes _data) returns()
func (_Collateral *CollateralTransactorSession) AuctionClosed(_id *big.Int, _leftover *big.Int, _received *big.Int, _data []byte) (*types.Transaction, error) {
	return _Collateral.Contract.AuctionClosed(&_Collateral.TransactOpts, _id, _leftover, _received, _data)
}

// BorrowCollateral is a paid mutator transaction binding the contract method 0x8c9c294a.
//
// Solidity: function borrowCollateral(uint256 _entryId, address _handler, bytes _data, bytes _oracleData) returns()
func (_Collateral *CollateralTransactor) BorrowCollateral(opts *bind.TransactOpts, _entryId *big.Int, _handler common.Address, _data []byte, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "borrowCollateral", _entryId, _handler, _data, _oracleData)
}

// BorrowCollateral is a paid mutator transaction binding the contract method 0x8c9c294a.
//
// Solidity: function borrowCollateral(uint256 _entryId, address _handler, bytes _data, bytes _oracleData) returns()
func (_Collateral *CollateralSession) BorrowCollateral(_entryId *big.Int, _handler common.Address, _data []byte, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.BorrowCollateral(&_Collateral.TransactOpts, _entryId, _handler, _data, _oracleData)
}

// BorrowCollateral is a paid mutator transaction binding the contract method 0x8c9c294a.
//
// Solidity: function borrowCollateral(uint256 _entryId, address _handler, bytes _data, bytes _oracleData) returns()
func (_Collateral *CollateralTransactorSession) BorrowCollateral(_entryId *big.Int, _handler common.Address, _data []byte, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.BorrowCollateral(&_Collateral.TransactOpts, _entryId, _handler, _data, _oracleData)
}

// Claim is a paid mutator transaction binding the contract method 0x8f0bc152.
//
// Solidity: function claim(address , uint256 _debtId, bytes _oracleData) returns(bool)
func (_Collateral *CollateralTransactor) Claim(opts *bind.TransactOpts, arg0 common.Address, _debtId *big.Int, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "claim", arg0, _debtId, _oracleData)
}

// Claim is a paid mutator transaction binding the contract method 0x8f0bc152.
//
// Solidity: function claim(address , uint256 _debtId, bytes _oracleData) returns(bool)
func (_Collateral *CollateralSession) Claim(arg0 common.Address, _debtId *big.Int, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.Claim(&_Collateral.TransactOpts, arg0, _debtId, _oracleData)
}

// Claim is a paid mutator transaction binding the contract method 0x8f0bc152.
//
// Solidity: function claim(address , uint256 _debtId, bytes _oracleData) returns(bool)
func (_Collateral *CollateralTransactorSession) Claim(arg0 common.Address, _debtId *big.Int, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.Claim(&_Collateral.TransactOpts, arg0, _debtId, _oracleData)
}

// Create is a paid mutator transaction binding the contract method 0x19bb5aaf.
//
// Solidity: function create(bytes32 _debtId, address _oracle, uint256 _amount, uint96 _liquidationRatio, uint96 _balanceRatio) returns(uint256 entryId)
func (_Collateral *CollateralTransactor) Create(opts *bind.TransactOpts, _debtId [32]byte, _oracle common.Address, _amount *big.Int, _liquidationRatio *big.Int, _balanceRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "create", _debtId, _oracle, _amount, _liquidationRatio, _balanceRatio)
}

// Create is a paid mutator transaction binding the contract method 0x19bb5aaf.
//
// Solidity: function create(bytes32 _debtId, address _oracle, uint256 _amount, uint96 _liquidationRatio, uint96 _balanceRatio) returns(uint256 entryId)
func (_Collateral *CollateralSession) Create(_debtId [32]byte, _oracle common.Address, _amount *big.Int, _liquidationRatio *big.Int, _balanceRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Create(&_Collateral.TransactOpts, _debtId, _oracle, _amount, _liquidationRatio, _balanceRatio)
}

// Create is a paid mutator transaction binding the contract method 0x19bb5aaf.
//
// Solidity: function create(bytes32 _debtId, address _oracle, uint256 _amount, uint96 _liquidationRatio, uint96 _balanceRatio) returns(uint256 entryId)
func (_Collateral *CollateralTransactorSession) Create(_debtId [32]byte, _oracle common.Address, _amount *big.Int, _liquidationRatio *big.Int, _balanceRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Create(&_Collateral.TransactOpts, _debtId, _oracle, _amount, _liquidationRatio, _balanceRatio)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _entryId, uint256 _amount) returns()
func (_Collateral *CollateralTransactor) Deposit(opts *bind.TransactOpts, _entryId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "deposit", _entryId, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _entryId, uint256 _amount) returns()
func (_Collateral *CollateralSession) Deposit(_entryId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Deposit(&_Collateral.TransactOpts, _entryId, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _entryId, uint256 _amount) returns()
func (_Collateral *CollateralTransactorSession) Deposit(_entryId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Deposit(&_Collateral.TransactOpts, _entryId, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 _entryId, address _to) returns()
func (_Collateral *CollateralTransactor) Redeem(opts *bind.TransactOpts, _entryId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "redeem", _entryId, _to)
}

// Redeem is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 _entryId, address _to) returns()
func (_Collateral *CollateralSession) Redeem(_entryId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Redeem(&_Collateral.TransactOpts, _entryId, _to)
}

// Redeem is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 _entryId, address _to) returns()
func (_Collateral *CollateralTransactorSession) Redeem(_entryId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Redeem(&_Collateral.TransactOpts, _entryId, _to)
}

// RequestCosign is a paid mutator transaction binding the contract method 0xbb26d9d1.
//
// Solidity: function requestCosign(address , uint256 _debtId, bytes _data, bytes _oracleData) returns(bool)
func (_Collateral *CollateralTransactor) RequestCosign(opts *bind.TransactOpts, arg0 common.Address, _debtId *big.Int, _data []byte, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "requestCosign", arg0, _debtId, _data, _oracleData)
}

// RequestCosign is a paid mutator transaction binding the contract method 0xbb26d9d1.
//
// Solidity: function requestCosign(address , uint256 _debtId, bytes _data, bytes _oracleData) returns(bool)
func (_Collateral *CollateralSession) RequestCosign(arg0 common.Address, _debtId *big.Int, _data []byte, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.RequestCosign(&_Collateral.TransactOpts, arg0, _debtId, _data, _oracleData)
}

// RequestCosign is a paid mutator transaction binding the contract method 0xbb26d9d1.
//
// Solidity: function requestCosign(address , uint256 _debtId, bytes _data, bytes _oracleData) returns(bool)
func (_Collateral *CollateralTransactorSession) RequestCosign(arg0 common.Address, _debtId *big.Int, _data []byte, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.RequestCosign(&_Collateral.TransactOpts, arg0, _debtId, _data, _oracleData)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId) returns()
func (_Collateral *CollateralTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "safeTransferFrom", _from, _to, _assetId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId) returns()
func (_Collateral *CollateralSession) SafeTransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SafeTransferFrom(&_Collateral.TransactOpts, _from, _to, _assetId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId) returns()
func (_Collateral *CollateralTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SafeTransferFrom(&_Collateral.TransactOpts, _from, _to, _assetId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId, bytes _userData) returns()
func (_Collateral *CollateralTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _assetId *big.Int, _userData []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "safeTransferFrom0", _from, _to, _assetId, _userData)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId, bytes _userData) returns()
func (_Collateral *CollateralSession) SafeTransferFrom0(_from common.Address, _to common.Address, _assetId *big.Int, _userData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.SafeTransferFrom0(&_Collateral.TransactOpts, _from, _to, _assetId, _userData)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId, bytes _userData) returns()
func (_Collateral *CollateralTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _assetId *big.Int, _userData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.SafeTransferFrom0(&_Collateral.TransactOpts, _from, _to, _assetId, _userData)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _authorized) returns()
func (_Collateral *CollateralTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _authorized bool) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "setApprovalForAll", _operator, _authorized)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _authorized) returns()
func (_Collateral *CollateralSession) SetApprovalForAll(_operator common.Address, _authorized bool) (*types.Transaction, error) {
	return _Collateral.Contract.SetApprovalForAll(&_Collateral.TransactOpts, _operator, _authorized)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _authorized) returns()
func (_Collateral *CollateralTransactorSession) SetApprovalForAll(_operator common.Address, _authorized bool) (*types.Transaction, error) {
	return _Collateral.Contract.SetApprovalForAll(&_Collateral.TransactOpts, _operator, _authorized)
}

// SetUrl is a paid mutator transaction binding the contract method 0x252498a2.
//
// Solidity: function setUrl(string _url) returns()
func (_Collateral *CollateralTransactor) SetUrl(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "setUrl", _url)
}

// SetUrl is a paid mutator transaction binding the contract method 0x252498a2.
//
// Solidity: function setUrl(string _url) returns()
func (_Collateral *CollateralSession) SetUrl(_url string) (*types.Transaction, error) {
	return _Collateral.Contract.SetUrl(&_Collateral.TransactOpts, _url)
}

// SetUrl is a paid mutator transaction binding the contract method 0x252498a2.
//
// Solidity: function setUrl(string _url) returns()
func (_Collateral *CollateralTransactorSession) SetUrl(_url string) (*types.Transaction, error) {
	return _Collateral.Contract.SetUrl(&_Collateral.TransactOpts, _url)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _assetId) returns()
func (_Collateral *CollateralTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "transferFrom", _from, _to, _assetId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _assetId) returns()
func (_Collateral *CollateralSession) TransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.TransferFrom(&_Collateral.TransactOpts, _from, _to, _assetId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _assetId) returns()
func (_Collateral *CollateralTransactorSession) TransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.TransferFrom(&_Collateral.TransactOpts, _from, _to, _assetId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Collateral *CollateralTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Collateral *CollateralSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.TransferOwnership(&_Collateral.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Collateral *CollateralTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.TransferOwnership(&_Collateral.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdcbd7a53.
//
// Solidity: function withdraw(uint256 _entryId, address _to, uint256 _amount, bytes _oracleData) returns()
func (_Collateral *CollateralTransactor) Withdraw(opts *bind.TransactOpts, _entryId *big.Int, _to common.Address, _amount *big.Int, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "withdraw", _entryId, _to, _amount, _oracleData)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdcbd7a53.
//
// Solidity: function withdraw(uint256 _entryId, address _to, uint256 _amount, bytes _oracleData) returns()
func (_Collateral *CollateralSession) Withdraw(_entryId *big.Int, _to common.Address, _amount *big.Int, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.Withdraw(&_Collateral.TransactOpts, _entryId, _to, _amount, _oracleData)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdcbd7a53.
//
// Solidity: function withdraw(uint256 _entryId, address _to, uint256 _amount, bytes _oracleData) returns()
func (_Collateral *CollateralTransactorSession) Withdraw(_entryId *big.Int, _to common.Address, _amount *big.Int, _oracleData []byte) (*types.Transaction, error) {
	return _Collateral.Contract.Withdraw(&_Collateral.TransactOpts, _entryId, _to, _amount, _oracleData)
}

// CollateralApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Collateral contract.
type CollateralApprovalIterator struct {
	Event *CollateralApproval // Event containing the contract specifics and raw log

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
func (it *CollateralApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralApproval)
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
		it.Event = new(CollateralApproval)
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
func (it *CollateralApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralApproval represents a Approval event raised by the Collateral contract.
type CollateralApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Collateral *CollateralFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*CollateralApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralApprovalIterator{contract: _Collateral.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Collateral *CollateralFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CollateralApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralApproval)
				if err := _Collateral.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Collateral *CollateralFilterer) ParseApproval(log types.Log) (*CollateralApproval, error) {
	event := new(CollateralApproval)
	if err := _Collateral.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Collateral contract.
type CollateralApprovalForAllIterator struct {
	Event *CollateralApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CollateralApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralApprovalForAll)
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
		it.Event = new(CollateralApprovalForAll)
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
func (it *CollateralApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralApprovalForAll represents a ApprovalForAll event raised by the Collateral contract.
type CollateralApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Collateral *CollateralFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*CollateralApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &CollateralApprovalForAllIterator{contract: _Collateral.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Collateral *CollateralFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CollateralApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralApprovalForAll)
				if err := _Collateral.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Collateral *CollateralFilterer) ParseApprovalForAll(log types.Log) (*CollateralApprovalForAll, error) {
	event := new(CollateralApprovalForAll)
	if err := _Collateral.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralBorrowCollateralIterator is returned from FilterBorrowCollateral and is used to iterate over the raw logs and unpacked data for BorrowCollateral events raised by the Collateral contract.
type CollateralBorrowCollateralIterator struct {
	Event *CollateralBorrowCollateral // Event containing the contract specifics and raw log

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
func (it *CollateralBorrowCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralBorrowCollateral)
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
		it.Event = new(CollateralBorrowCollateral)
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
func (it *CollateralBorrowCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralBorrowCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralBorrowCollateral represents a BorrowCollateral event raised by the Collateral contract.
type CollateralBorrowCollateral struct {
	EntryId   *big.Int
	Handler   common.Address
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBorrowCollateral is a free log retrieval operation binding the contract event 0x9b33a4c997038b93eec9f3b361e2d2ff5b830ce8a15d2adbb472e1a98f3d2fdf.
//
// Solidity: event BorrowCollateral(uint256 indexed _entryId, address _handler, uint256 _newAmount)
func (_Collateral *CollateralFilterer) FilterBorrowCollateral(opts *bind.FilterOpts, _entryId []*big.Int) (*CollateralBorrowCollateralIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "BorrowCollateral", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralBorrowCollateralIterator{contract: _Collateral.contract, event: "BorrowCollateral", logs: logs, sub: sub}, nil
}

// WatchBorrowCollateral is a free log subscription operation binding the contract event 0x9b33a4c997038b93eec9f3b361e2d2ff5b830ce8a15d2adbb472e1a98f3d2fdf.
//
// Solidity: event BorrowCollateral(uint256 indexed _entryId, address _handler, uint256 _newAmount)
func (_Collateral *CollateralFilterer) WatchBorrowCollateral(opts *bind.WatchOpts, sink chan<- *CollateralBorrowCollateral, _entryId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "BorrowCollateral", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralBorrowCollateral)
				if err := _Collateral.contract.UnpackLog(event, "BorrowCollateral", log); err != nil {
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

// ParseBorrowCollateral is a log parse operation binding the contract event 0x9b33a4c997038b93eec9f3b361e2d2ff5b830ce8a15d2adbb472e1a98f3d2fdf.
//
// Solidity: event BorrowCollateral(uint256 indexed _entryId, address _handler, uint256 _newAmount)
func (_Collateral *CollateralFilterer) ParseBorrowCollateral(log types.Log) (*CollateralBorrowCollateral, error) {
	event := new(CollateralBorrowCollateral)
	if err := _Collateral.contract.UnpackLog(event, "BorrowCollateral", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralClaimedExpiredIterator is returned from FilterClaimedExpired and is used to iterate over the raw logs and unpacked data for ClaimedExpired events raised by the Collateral contract.
type CollateralClaimedExpiredIterator struct {
	Event *CollateralClaimedExpired // Event containing the contract specifics and raw log

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
func (it *CollateralClaimedExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralClaimedExpired)
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
		it.Event = new(CollateralClaimedExpired)
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
func (it *CollateralClaimedExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralClaimedExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralClaimedExpired represents a ClaimedExpired event raised by the Collateral contract.
type CollateralClaimedExpired struct {
	EntryId          *big.Int
	AuctionId        *big.Int
	DueTime          *big.Int
	Obligation       *big.Int
	ObligationTokens *big.Int
	MarketValue      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimedExpired is a free log retrieval operation binding the contract event 0x0ce0f25882fd9e89677eed12d539fa3abc08a0bb845b0c9717fb5cdfa350ab12.
//
// Solidity: event ClaimedExpired(uint256 indexed _entryId, uint256 indexed _auctionId, uint256 _dueTime, uint256 _obligation, uint256 _obligationTokens, uint256 _marketValue)
func (_Collateral *CollateralFilterer) FilterClaimedExpired(opts *bind.FilterOpts, _entryId []*big.Int, _auctionId []*big.Int) (*CollateralClaimedExpiredIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}
	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "ClaimedExpired", _entryIdRule, _auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralClaimedExpiredIterator{contract: _Collateral.contract, event: "ClaimedExpired", logs: logs, sub: sub}, nil
}

// WatchClaimedExpired is a free log subscription operation binding the contract event 0x0ce0f25882fd9e89677eed12d539fa3abc08a0bb845b0c9717fb5cdfa350ab12.
//
// Solidity: event ClaimedExpired(uint256 indexed _entryId, uint256 indexed _auctionId, uint256 _dueTime, uint256 _obligation, uint256 _obligationTokens, uint256 _marketValue)
func (_Collateral *CollateralFilterer) WatchClaimedExpired(opts *bind.WatchOpts, sink chan<- *CollateralClaimedExpired, _entryId []*big.Int, _auctionId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}
	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "ClaimedExpired", _entryIdRule, _auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralClaimedExpired)
				if err := _Collateral.contract.UnpackLog(event, "ClaimedExpired", log); err != nil {
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

// ParseClaimedExpired is a log parse operation binding the contract event 0x0ce0f25882fd9e89677eed12d539fa3abc08a0bb845b0c9717fb5cdfa350ab12.
//
// Solidity: event ClaimedExpired(uint256 indexed _entryId, uint256 indexed _auctionId, uint256 _dueTime, uint256 _obligation, uint256 _obligationTokens, uint256 _marketValue)
func (_Collateral *CollateralFilterer) ParseClaimedExpired(log types.Log) (*CollateralClaimedExpired, error) {
	event := new(CollateralClaimedExpired)
	if err := _Collateral.contract.UnpackLog(event, "ClaimedExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralClaimedLiquidationIterator is returned from FilterClaimedLiquidation and is used to iterate over the raw logs and unpacked data for ClaimedLiquidation events raised by the Collateral contract.
type CollateralClaimedLiquidationIterator struct {
	Event *CollateralClaimedLiquidation // Event containing the contract specifics and raw log

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
func (it *CollateralClaimedLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralClaimedLiquidation)
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
		it.Event = new(CollateralClaimedLiquidation)
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
func (it *CollateralClaimedLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralClaimedLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralClaimedLiquidation represents a ClaimedLiquidation event raised by the Collateral contract.
type CollateralClaimedLiquidation struct {
	EntryId     *big.Int
	AuctionId   *big.Int
	Debt        *big.Int
	Required    *big.Int
	MarketValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimedLiquidation is a free log retrieval operation binding the contract event 0x4df462f6bb661fc317efa25950d47fbd7f4ad62c6b69664c01bab0f119f02d32.
//
// Solidity: event ClaimedLiquidation(uint256 indexed _entryId, uint256 indexed _auctionId, uint256 _debt, uint256 _required, uint256 _marketValue)
func (_Collateral *CollateralFilterer) FilterClaimedLiquidation(opts *bind.FilterOpts, _entryId []*big.Int, _auctionId []*big.Int) (*CollateralClaimedLiquidationIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}
	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "ClaimedLiquidation", _entryIdRule, _auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralClaimedLiquidationIterator{contract: _Collateral.contract, event: "ClaimedLiquidation", logs: logs, sub: sub}, nil
}

// WatchClaimedLiquidation is a free log subscription operation binding the contract event 0x4df462f6bb661fc317efa25950d47fbd7f4ad62c6b69664c01bab0f119f02d32.
//
// Solidity: event ClaimedLiquidation(uint256 indexed _entryId, uint256 indexed _auctionId, uint256 _debt, uint256 _required, uint256 _marketValue)
func (_Collateral *CollateralFilterer) WatchClaimedLiquidation(opts *bind.WatchOpts, sink chan<- *CollateralClaimedLiquidation, _entryId []*big.Int, _auctionId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}
	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "ClaimedLiquidation", _entryIdRule, _auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralClaimedLiquidation)
				if err := _Collateral.contract.UnpackLog(event, "ClaimedLiquidation", log); err != nil {
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

// ParseClaimedLiquidation is a log parse operation binding the contract event 0x4df462f6bb661fc317efa25950d47fbd7f4ad62c6b69664c01bab0f119f02d32.
//
// Solidity: event ClaimedLiquidation(uint256 indexed _entryId, uint256 indexed _auctionId, uint256 _debt, uint256 _required, uint256 _marketValue)
func (_Collateral *CollateralFilterer) ParseClaimedLiquidation(log types.Log) (*CollateralClaimedLiquidation, error) {
	event := new(CollateralClaimedLiquidation)
	if err := _Collateral.contract.UnpackLog(event, "ClaimedLiquidation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralClosedAuctionIterator is returned from FilterClosedAuction and is used to iterate over the raw logs and unpacked data for ClosedAuction events raised by the Collateral contract.
type CollateralClosedAuctionIterator struct {
	Event *CollateralClosedAuction // Event containing the contract specifics and raw log

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
func (it *CollateralClosedAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralClosedAuction)
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
		it.Event = new(CollateralClosedAuction)
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
func (it *CollateralClosedAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralClosedAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralClosedAuction represents a ClosedAuction event raised by the Collateral contract.
type CollateralClosedAuction struct {
	EntryId  *big.Int
	Received *big.Int
	Leftover *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClosedAuction is a free log retrieval operation binding the contract event 0xeff56510b465b07772acd0d1a3723fe4a7cce582a26efe952f226872fbc3595c.
//
// Solidity: event ClosedAuction(uint256 indexed _entryId, uint256 _received, uint256 _leftover)
func (_Collateral *CollateralFilterer) FilterClosedAuction(opts *bind.FilterOpts, _entryId []*big.Int) (*CollateralClosedAuctionIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "ClosedAuction", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralClosedAuctionIterator{contract: _Collateral.contract, event: "ClosedAuction", logs: logs, sub: sub}, nil
}

// WatchClosedAuction is a free log subscription operation binding the contract event 0xeff56510b465b07772acd0d1a3723fe4a7cce582a26efe952f226872fbc3595c.
//
// Solidity: event ClosedAuction(uint256 indexed _entryId, uint256 _received, uint256 _leftover)
func (_Collateral *CollateralFilterer) WatchClosedAuction(opts *bind.WatchOpts, sink chan<- *CollateralClosedAuction, _entryId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "ClosedAuction", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralClosedAuction)
				if err := _Collateral.contract.UnpackLog(event, "ClosedAuction", log); err != nil {
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

// ParseClosedAuction is a log parse operation binding the contract event 0xeff56510b465b07772acd0d1a3723fe4a7cce582a26efe952f226872fbc3595c.
//
// Solidity: event ClosedAuction(uint256 indexed _entryId, uint256 _received, uint256 _leftover)
func (_Collateral *CollateralFilterer) ParseClosedAuction(log types.Log) (*CollateralClosedAuction, error) {
	event := new(CollateralClosedAuction)
	if err := _Collateral.contract.UnpackLog(event, "ClosedAuction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Collateral contract.
type CollateralCreatedIterator struct {
	Event *CollateralCreated // Event containing the contract specifics and raw log

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
func (it *CollateralCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralCreated)
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
		it.Event = new(CollateralCreated)
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
func (it *CollateralCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralCreated represents a Created event raised by the Collateral contract.
type CollateralCreated struct {
	EntryId          *big.Int
	DebtId           [32]byte
	Oracle           common.Address
	Token            common.Address
	Amount           *big.Int
	LiquidationRatio *big.Int
	BalanceRatio     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0xfe5bb50bc3d4e5eea6293acf0cb8bb5ff31625bb76a447fa888dd393b980b52a.
//
// Solidity: event Created(uint256 indexed _entryId, bytes32 indexed _debtId, address _oracle, address _token, uint256 _amount, uint96 _liquidationRatio, uint96 _balanceRatio)
func (_Collateral *CollateralFilterer) FilterCreated(opts *bind.FilterOpts, _entryId []*big.Int, _debtId [][32]byte) (*CollateralCreatedIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}
	var _debtIdRule []interface{}
	for _, _debtIdItem := range _debtId {
		_debtIdRule = append(_debtIdRule, _debtIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Created", _entryIdRule, _debtIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralCreatedIterator{contract: _Collateral.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0xfe5bb50bc3d4e5eea6293acf0cb8bb5ff31625bb76a447fa888dd393b980b52a.
//
// Solidity: event Created(uint256 indexed _entryId, bytes32 indexed _debtId, address _oracle, address _token, uint256 _amount, uint96 _liquidationRatio, uint96 _balanceRatio)
func (_Collateral *CollateralFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *CollateralCreated, _entryId []*big.Int, _debtId [][32]byte) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}
	var _debtIdRule []interface{}
	for _, _debtIdItem := range _debtId {
		_debtIdRule = append(_debtIdRule, _debtIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Created", _entryIdRule, _debtIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralCreated)
				if err := _Collateral.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0xfe5bb50bc3d4e5eea6293acf0cb8bb5ff31625bb76a447fa888dd393b980b52a.
//
// Solidity: event Created(uint256 indexed _entryId, bytes32 indexed _debtId, address _oracle, address _token, uint256 _amount, uint96 _liquidationRatio, uint96 _balanceRatio)
func (_Collateral *CollateralFilterer) ParseCreated(log types.Log) (*CollateralCreated, error) {
	event := new(CollateralCreated)
	if err := _Collateral.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Collateral contract.
type CollateralDepositedIterator struct {
	Event *CollateralDeposited // Event containing the contract specifics and raw log

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
func (it *CollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralDeposited)
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
		it.Event = new(CollateralDeposited)
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
func (it *CollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralDeposited represents a Deposited event raised by the Collateral contract.
type CollateralDeposited struct {
	EntryId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x06da3309189fa49284f335d2c2bcb4cb0b8ad2a59ad92a9bdebeeb8f1ceba511.
//
// Solidity: event Deposited(uint256 indexed _entryId, uint256 _amount)
func (_Collateral *CollateralFilterer) FilterDeposited(opts *bind.FilterOpts, _entryId []*big.Int) (*CollateralDepositedIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Deposited", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralDepositedIterator{contract: _Collateral.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x06da3309189fa49284f335d2c2bcb4cb0b8ad2a59ad92a9bdebeeb8f1ceba511.
//
// Solidity: event Deposited(uint256 indexed _entryId, uint256 _amount)
func (_Collateral *CollateralFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CollateralDeposited, _entryId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Deposited", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralDeposited)
				if err := _Collateral.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x06da3309189fa49284f335d2c2bcb4cb0b8ad2a59ad92a9bdebeeb8f1ceba511.
//
// Solidity: event Deposited(uint256 indexed _entryId, uint256 _amount)
func (_Collateral *CollateralFilterer) ParseDeposited(log types.Log) (*CollateralDeposited, error) {
	event := new(CollateralDeposited)
	if err := _Collateral.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Collateral contract.
type CollateralOwnershipTransferredIterator struct {
	Event *CollateralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollateralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralOwnershipTransferred)
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
		it.Event = new(CollateralOwnershipTransferred)
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
func (it *CollateralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralOwnershipTransferred represents a OwnershipTransferred event raised by the Collateral contract.
type CollateralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_Collateral *CollateralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*CollateralOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralOwnershipTransferredIterator{contract: _Collateral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_Collateral *CollateralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollateralOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralOwnershipTransferred)
				if err := _Collateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Collateral *CollateralFilterer) ParseOwnershipTransferred(log types.Log) (*CollateralOwnershipTransferred, error) {
	event := new(CollateralOwnershipTransferred)
	if err := _Collateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the Collateral contract.
type CollateralRedeemedIterator struct {
	Event *CollateralRedeemed // Event containing the contract specifics and raw log

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
func (it *CollateralRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralRedeemed)
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
		it.Event = new(CollateralRedeemed)
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
func (it *CollateralRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralRedeemed represents a Redeemed event raised by the Collateral contract.
type CollateralRedeemed struct {
	EntryId *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0xc5fcfb68332ef11d542d4ab7e75045a5e4d66eb2d8f846f13a48356b32e534db.
//
// Solidity: event Redeemed(uint256 indexed _entryId, address _to)
func (_Collateral *CollateralFilterer) FilterRedeemed(opts *bind.FilterOpts, _entryId []*big.Int) (*CollateralRedeemedIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Redeemed", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralRedeemedIterator{contract: _Collateral.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0xc5fcfb68332ef11d542d4ab7e75045a5e4d66eb2d8f846f13a48356b32e534db.
//
// Solidity: event Redeemed(uint256 indexed _entryId, address _to)
func (_Collateral *CollateralFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *CollateralRedeemed, _entryId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Redeemed", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralRedeemed)
				if err := _Collateral.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0xc5fcfb68332ef11d542d4ab7e75045a5e4d66eb2d8f846f13a48356b32e534db.
//
// Solidity: event Redeemed(uint256 indexed _entryId, address _to)
func (_Collateral *CollateralFilterer) ParseRedeemed(log types.Log) (*CollateralRedeemed, error) {
	event := new(CollateralRedeemed)
	if err := _Collateral.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralSetURIProviderIterator is returned from FilterSetURIProvider and is used to iterate over the raw logs and unpacked data for SetURIProvider events raised by the Collateral contract.
type CollateralSetURIProviderIterator struct {
	Event *CollateralSetURIProvider // Event containing the contract specifics and raw log

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
func (it *CollateralSetURIProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralSetURIProvider)
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
		it.Event = new(CollateralSetURIProvider)
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
func (it *CollateralSetURIProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralSetURIProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralSetURIProvider represents a SetURIProvider event raised by the Collateral contract.
type CollateralSetURIProvider struct {
	UriProvider common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetURIProvider is a free log retrieval operation binding the contract event 0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec.
//
// Solidity: event SetURIProvider(address _uriProvider)
func (_Collateral *CollateralFilterer) FilterSetURIProvider(opts *bind.FilterOpts) (*CollateralSetURIProviderIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "SetURIProvider")
	if err != nil {
		return nil, err
	}
	return &CollateralSetURIProviderIterator{contract: _Collateral.contract, event: "SetURIProvider", logs: logs, sub: sub}, nil
}

// WatchSetURIProvider is a free log subscription operation binding the contract event 0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec.
//
// Solidity: event SetURIProvider(address _uriProvider)
func (_Collateral *CollateralFilterer) WatchSetURIProvider(opts *bind.WatchOpts, sink chan<- *CollateralSetURIProvider) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "SetURIProvider")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralSetURIProvider)
				if err := _Collateral.contract.UnpackLog(event, "SetURIProvider", log); err != nil {
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

// ParseSetURIProvider is a log parse operation binding the contract event 0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec.
//
// Solidity: event SetURIProvider(address _uriProvider)
func (_Collateral *CollateralFilterer) ParseSetURIProvider(log types.Log) (*CollateralSetURIProvider, error) {
	event := new(CollateralSetURIProvider)
	if err := _Collateral.contract.UnpackLog(event, "SetURIProvider", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralSetUrlIterator is returned from FilterSetUrl and is used to iterate over the raw logs and unpacked data for SetUrl events raised by the Collateral contract.
type CollateralSetUrlIterator struct {
	Event *CollateralSetUrl // Event containing the contract specifics and raw log

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
func (it *CollateralSetUrlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralSetUrl)
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
		it.Event = new(CollateralSetUrl)
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
func (it *CollateralSetUrlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralSetUrlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralSetUrl represents a SetUrl event raised by the Collateral contract.
type CollateralSetUrl struct {
	Url string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetUrl is a free log retrieval operation binding the contract event 0x0c2d986c2593d7ad7d31daddf853e7a3f2a6a91c2dfd596e060df04bc174df5e.
//
// Solidity: event SetUrl(string _url)
func (_Collateral *CollateralFilterer) FilterSetUrl(opts *bind.FilterOpts) (*CollateralSetUrlIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "SetUrl")
	if err != nil {
		return nil, err
	}
	return &CollateralSetUrlIterator{contract: _Collateral.contract, event: "SetUrl", logs: logs, sub: sub}, nil
}

// WatchSetUrl is a free log subscription operation binding the contract event 0x0c2d986c2593d7ad7d31daddf853e7a3f2a6a91c2dfd596e060df04bc174df5e.
//
// Solidity: event SetUrl(string _url)
func (_Collateral *CollateralFilterer) WatchSetUrl(opts *bind.WatchOpts, sink chan<- *CollateralSetUrl) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "SetUrl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralSetUrl)
				if err := _Collateral.contract.UnpackLog(event, "SetUrl", log); err != nil {
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

// ParseSetUrl is a log parse operation binding the contract event 0x0c2d986c2593d7ad7d31daddf853e7a3f2a6a91c2dfd596e060df04bc174df5e.
//
// Solidity: event SetUrl(string _url)
func (_Collateral *CollateralFilterer) ParseSetUrl(log types.Log) (*CollateralSetUrl, error) {
	event := new(CollateralSetUrl)
	if err := _Collateral.contract.UnpackLog(event, "SetUrl", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralStartedIterator is returned from FilterStarted and is used to iterate over the raw logs and unpacked data for Started events raised by the Collateral contract.
type CollateralStartedIterator struct {
	Event *CollateralStarted // Event containing the contract specifics and raw log

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
func (it *CollateralStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralStarted)
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
		it.Event = new(CollateralStarted)
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
func (it *CollateralStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralStarted represents a Started event raised by the Collateral contract.
type CollateralStarted struct {
	EntryId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStarted is a free log retrieval operation binding the contract event 0x006e0c97de781a7389d44ba8fd35d1467cabb17ed04d038d166d34ab819213f3.
//
// Solidity: event Started(uint256 indexed _entryId)
func (_Collateral *CollateralFilterer) FilterStarted(opts *bind.FilterOpts, _entryId []*big.Int) (*CollateralStartedIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Started", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralStartedIterator{contract: _Collateral.contract, event: "Started", logs: logs, sub: sub}, nil
}

// WatchStarted is a free log subscription operation binding the contract event 0x006e0c97de781a7389d44ba8fd35d1467cabb17ed04d038d166d34ab819213f3.
//
// Solidity: event Started(uint256 indexed _entryId)
func (_Collateral *CollateralFilterer) WatchStarted(opts *bind.WatchOpts, sink chan<- *CollateralStarted, _entryId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Started", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralStarted)
				if err := _Collateral.contract.UnpackLog(event, "Started", log); err != nil {
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

// ParseStarted is a log parse operation binding the contract event 0x006e0c97de781a7389d44ba8fd35d1467cabb17ed04d038d166d34ab819213f3.
//
// Solidity: event Started(uint256 indexed _entryId)
func (_Collateral *CollateralFilterer) ParseStarted(log types.Log) (*CollateralStarted, error) {
	event := new(CollateralStarted)
	if err := _Collateral.contract.UnpackLog(event, "Started", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Collateral contract.
type CollateralTransferIterator struct {
	Event *CollateralTransfer // Event containing the contract specifics and raw log

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
func (it *CollateralTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTransfer)
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
		it.Event = new(CollateralTransfer)
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
func (it *CollateralTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTransfer represents a Transfer event raised by the Collateral contract.
type CollateralTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Collateral *CollateralFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*CollateralTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTransferIterator{contract: _Collateral.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Collateral *CollateralFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CollateralTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTransfer)
				if err := _Collateral.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Collateral *CollateralFilterer) ParseTransfer(log types.Log) (*CollateralTransfer, error) {
	event := new(CollateralTransfer)
	if err := _Collateral.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CollateralWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Collateral contract.
type CollateralWithdrawIterator struct {
	Event *CollateralWithdraw // Event containing the contract specifics and raw log

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
func (it *CollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralWithdraw)
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
		it.Event = new(CollateralWithdraw)
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
func (it *CollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralWithdraw represents a Withdraw event raised by the Collateral contract.
type CollateralWithdraw struct {
	EntryId *big.Int
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 indexed _entryId, address _to, uint256 _amount)
func (_Collateral *CollateralFilterer) FilterWithdraw(opts *bind.FilterOpts, _entryId []*big.Int) (*CollateralWithdrawIterator, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Withdraw", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralWithdrawIterator{contract: _Collateral.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 indexed _entryId, address _to, uint256 _amount)
func (_Collateral *CollateralFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CollateralWithdraw, _entryId []*big.Int) (event.Subscription, error) {

	var _entryIdRule []interface{}
	for _, _entryIdItem := range _entryId {
		_entryIdRule = append(_entryIdRule, _entryIdItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Withdraw", _entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralWithdraw)
				if err := _Collateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 indexed _entryId, address _to, uint256 _amount)
func (_Collateral *CollateralFilterer) ParseWithdraw(log types.Log) (*CollateralWithdraw, error) {
	event := new(CollateralWithdraw)
	if err := _Collateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
