// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DebtEngine

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

// DebtEngineABI is the input ABI used to generate the binding from.
const DebtEngineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidToken\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"assetsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractModel\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"create2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"buildId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"payTokenBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"paid\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"paidTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractModel\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"oracleData\",\"type\":\"bytes\"}],\"name\":\"payToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidToken\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getError\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_authorized\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"debts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"error\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"},{\"internalType\":\"contractModel\",\"name\":\"model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractModel\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"create3\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_userData\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"buildId3\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_oracleData\",\"type\":\"bytes\"}],\"name\":\"payBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"paid\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"paidTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractURIProvider\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"setURIProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetHolder\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"run\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPartial\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"buildId2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"Created2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"Created3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestedTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_paid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_equivalent\",\"type\":\"uint256\"}],\"name\":\"ReadedOracleBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_equivalent\",\"type\":\"uint256\"}],\"name\":\"ReadedOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_targetOracle\",\"type\":\"address\"}],\"name\":\"PayBatchError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_result\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"ErrorRecover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_uriProvider\",\"type\":\"address\"}],\"name\":\"SetURIProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"SetName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"SetSymbol\",\"type\":\"event\"}]"

// DebtEngine is an auto generated Go binding around an Ethereum contract.
type DebtEngine struct {
	DebtEngineCaller     // Read-only binding to the contract
	DebtEngineTransactor // Write-only binding to the contract
	DebtEngineFilterer   // Log filterer for contract events
}

// DebtEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebtEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebtEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebtEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebtEngineSession struct {
	Contract     *DebtEngine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebtEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebtEngineCallerSession struct {
	Contract *DebtEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebtEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebtEngineTransactorSession struct {
	Contract     *DebtEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebtEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebtEngineRaw struct {
	Contract *DebtEngine // Generic contract binding to access the raw methods on
}

// DebtEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebtEngineCallerRaw struct {
	Contract *DebtEngineCaller // Generic read-only contract binding to access the raw methods on
}

// DebtEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebtEngineTransactorRaw struct {
	Contract *DebtEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebtEngine creates a new instance of DebtEngine, bound to a specific deployed contract.
func NewDebtEngine(address common.Address, backend bind.ContractBackend) (*DebtEngine, error) {
	contract, err := bindDebtEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebtEngine{DebtEngineCaller: DebtEngineCaller{contract: contract}, DebtEngineTransactor: DebtEngineTransactor{contract: contract}, DebtEngineFilterer: DebtEngineFilterer{contract: contract}}, nil
}

// NewDebtEngineCaller creates a new read-only instance of DebtEngine, bound to a specific deployed contract.
func NewDebtEngineCaller(address common.Address, caller bind.ContractCaller) (*DebtEngineCaller, error) {
	contract, err := bindDebtEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebtEngineCaller{contract: contract}, nil
}

// NewDebtEngineTransactor creates a new write-only instance of DebtEngine, bound to a specific deployed contract.
func NewDebtEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*DebtEngineTransactor, error) {
	contract, err := bindDebtEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebtEngineTransactor{contract: contract}, nil
}

// NewDebtEngineFilterer creates a new log filterer instance of DebtEngine, bound to a specific deployed contract.
func NewDebtEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*DebtEngineFilterer, error) {
	contract, err := bindDebtEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebtEngineFilterer{contract: contract}, nil
}

// bindDebtEngine binds a generic wrapper to an already deployed contract.
func bindDebtEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebtEngineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebtEngine *DebtEngineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebtEngine.Contract.DebtEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebtEngine *DebtEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtEngine.Contract.DebtEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebtEngine *DebtEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebtEngine.Contract.DebtEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebtEngine *DebtEngineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebtEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebtEngine *DebtEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebtEngine *DebtEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebtEngine.Contract.contract.Transact(opts, method, params...)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() constant returns(uint256[])
func (_DebtEngine *DebtEngineCaller) AllTokens(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "allTokens")
	return *ret0, err
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() constant returns(uint256[])
func (_DebtEngine *DebtEngineSession) AllTokens() ([]*big.Int, error) {
	return _DebtEngine.Contract.AllTokens(&_DebtEngine.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() constant returns(uint256[])
func (_DebtEngine *DebtEngineCallerSession) AllTokens() ([]*big.Int, error) {
	return _DebtEngine.Contract.AllTokens(&_DebtEngine.CallOpts)
}

// AssetsOf is a free data retrieval call binding the contract method 0x2c62fa10.
//
// Solidity: function assetsOf(address _owner) constant returns(uint256[])
func (_DebtEngine *DebtEngineCaller) AssetsOf(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "assetsOf", _owner)
	return *ret0, err
}

// AssetsOf is a free data retrieval call binding the contract method 0x2c62fa10.
//
// Solidity: function assetsOf(address _owner) constant returns(uint256[])
func (_DebtEngine *DebtEngineSession) AssetsOf(_owner common.Address) ([]*big.Int, error) {
	return _DebtEngine.Contract.AssetsOf(&_DebtEngine.CallOpts, _owner)
}

// AssetsOf is a free data retrieval call binding the contract method 0x2c62fa10.
//
// Solidity: function assetsOf(address _owner) constant returns(uint256[])
func (_DebtEngine *DebtEngineCallerSession) AssetsOf(_owner common.Address) ([]*big.Int, error) {
	return _DebtEngine.Contract.AssetsOf(&_DebtEngine.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DebtEngine *DebtEngineCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DebtEngine *DebtEngineSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DebtEngine.Contract.BalanceOf(&_DebtEngine.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DebtEngine *DebtEngineCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DebtEngine.Contract.BalanceOf(&_DebtEngine.CallOpts, _owner)
}

// BuildId is a free data retrieval call binding the contract method 0x47ff6d7b.
//
// Solidity: function buildId(address _creator, uint256 _nonce) constant returns(bytes32)
func (_DebtEngine *DebtEngineCaller) BuildId(opts *bind.CallOpts, _creator common.Address, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "buildId", _creator, _nonce)
	return *ret0, err
}

// BuildId is a free data retrieval call binding the contract method 0x47ff6d7b.
//
// Solidity: function buildId(address _creator, uint256 _nonce) constant returns(bytes32)
func (_DebtEngine *DebtEngineSession) BuildId(_creator common.Address, _nonce *big.Int) ([32]byte, error) {
	return _DebtEngine.Contract.BuildId(&_DebtEngine.CallOpts, _creator, _nonce)
}

// BuildId is a free data retrieval call binding the contract method 0x47ff6d7b.
//
// Solidity: function buildId(address _creator, uint256 _nonce) constant returns(bytes32)
func (_DebtEngine *DebtEngineCallerSession) BuildId(_creator common.Address, _nonce *big.Int) ([32]byte, error) {
	return _DebtEngine.Contract.BuildId(&_DebtEngine.CallOpts, _creator, _nonce)
}

// BuildId2 is a free data retrieval call binding the contract method 0xf57ae7b9.
//
// Solidity: function buildId2(address _creator, address _model, address _oracle, uint256 _salt, bytes _data) constant returns(bytes32)
func (_DebtEngine *DebtEngineCaller) BuildId2(opts *bind.CallOpts, _creator common.Address, _model common.Address, _oracle common.Address, _salt *big.Int, _data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "buildId2", _creator, _model, _oracle, _salt, _data)
	return *ret0, err
}

// BuildId2 is a free data retrieval call binding the contract method 0xf57ae7b9.
//
// Solidity: function buildId2(address _creator, address _model, address _oracle, uint256 _salt, bytes _data) constant returns(bytes32)
func (_DebtEngine *DebtEngineSession) BuildId2(_creator common.Address, _model common.Address, _oracle common.Address, _salt *big.Int, _data []byte) ([32]byte, error) {
	return _DebtEngine.Contract.BuildId2(&_DebtEngine.CallOpts, _creator, _model, _oracle, _salt, _data)
}

// BuildId2 is a free data retrieval call binding the contract method 0xf57ae7b9.
//
// Solidity: function buildId2(address _creator, address _model, address _oracle, uint256 _salt, bytes _data) constant returns(bytes32)
func (_DebtEngine *DebtEngineCallerSession) BuildId2(_creator common.Address, _model common.Address, _oracle common.Address, _salt *big.Int, _data []byte) ([32]byte, error) {
	return _DebtEngine.Contract.BuildId2(&_DebtEngine.CallOpts, _creator, _model, _oracle, _salt, _data)
}

// BuildId3 is a free data retrieval call binding the contract method 0xbb1dfeca.
//
// Solidity: function buildId3(address _creator, uint256 _salt) constant returns(bytes32)
func (_DebtEngine *DebtEngineCaller) BuildId3(opts *bind.CallOpts, _creator common.Address, _salt *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "buildId3", _creator, _salt)
	return *ret0, err
}

// BuildId3 is a free data retrieval call binding the contract method 0xbb1dfeca.
//
// Solidity: function buildId3(address _creator, uint256 _salt) constant returns(bytes32)
func (_DebtEngine *DebtEngineSession) BuildId3(_creator common.Address, _salt *big.Int) ([32]byte, error) {
	return _DebtEngine.Contract.BuildId3(&_DebtEngine.CallOpts, _creator, _salt)
}

// BuildId3 is a free data retrieval call binding the contract method 0xbb1dfeca.
//
// Solidity: function buildId3(address _creator, uint256 _salt) constant returns(bytes32)
func (_DebtEngine *DebtEngineCallerSession) BuildId3(_creator common.Address, _salt *big.Int) ([32]byte, error) {
	return _DebtEngine.Contract.BuildId3(&_DebtEngine.CallOpts, _creator, _salt)
}

// Debts is a free data retrieval call binding the contract method 0xa9caa411.
//
// Solidity: function debts(bytes32 ) constant returns(bool error, uint128 balance, address model, address creator, address oracle)
func (_DebtEngine *DebtEngineCaller) Debts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Error   bool
	Balance *big.Int
	Model   common.Address
	Creator common.Address
	Oracle  common.Address
}, error) {
	ret := new(struct {
		Error   bool
		Balance *big.Int
		Model   common.Address
		Creator common.Address
		Oracle  common.Address
	})
	out := ret
	err := _DebtEngine.contract.Call(opts, out, "debts", arg0)
	return *ret, err
}

// Debts is a free data retrieval call binding the contract method 0xa9caa411.
//
// Solidity: function debts(bytes32 ) constant returns(bool error, uint128 balance, address model, address creator, address oracle)
func (_DebtEngine *DebtEngineSession) Debts(arg0 [32]byte) (struct {
	Error   bool
	Balance *big.Int
	Model   common.Address
	Creator common.Address
	Oracle  common.Address
}, error) {
	return _DebtEngine.Contract.Debts(&_DebtEngine.CallOpts, arg0)
}

// Debts is a free data retrieval call binding the contract method 0xa9caa411.
//
// Solidity: function debts(bytes32 ) constant returns(bool error, uint128 balance, address model, address creator, address oracle)
func (_DebtEngine *DebtEngineCallerSession) Debts(arg0 [32]byte) (struct {
	Error   bool
	Balance *big.Int
	Model   common.Address
	Creator common.Address
	Oracle  common.Address
}, error) {
	return _DebtEngine.Contract.Debts(&_DebtEngine.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _assetId) constant returns(address)
func (_DebtEngine *DebtEngineCaller) GetApproved(opts *bind.CallOpts, _assetId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getApproved", _assetId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _assetId) constant returns(address)
func (_DebtEngine *DebtEngineSession) GetApproved(_assetId *big.Int) (common.Address, error) {
	return _DebtEngine.Contract.GetApproved(&_DebtEngine.CallOpts, _assetId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _assetId) constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) GetApproved(_assetId *big.Int) (common.Address, error) {
	return _DebtEngine.Contract.GetApproved(&_DebtEngine.CallOpts, _assetId)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _id) constant returns(uint128)
func (_DebtEngine *DebtEngineCaller) GetBalance(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getBalance", _id)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _id) constant returns(uint128)
func (_DebtEngine *DebtEngineSession) GetBalance(_id [32]byte) (*big.Int, error) {
	return _DebtEngine.Contract.GetBalance(&_DebtEngine.CallOpts, _id)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _id) constant returns(uint128)
func (_DebtEngine *DebtEngineCallerSession) GetBalance(_id [32]byte) (*big.Int, error) {
	return _DebtEngine.Contract.GetBalance(&_DebtEngine.CallOpts, _id)
}

// GetCreator is a free data retrieval call binding the contract method 0x060a7ef1.
//
// Solidity: function getCreator(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineCaller) GetCreator(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getCreator", _id)
	return *ret0, err
}

// GetCreator is a free data retrieval call binding the contract method 0x060a7ef1.
//
// Solidity: function getCreator(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineSession) GetCreator(_id [32]byte) (common.Address, error) {
	return _DebtEngine.Contract.GetCreator(&_DebtEngine.CallOpts, _id)
}

// GetCreator is a free data retrieval call binding the contract method 0x060a7ef1.
//
// Solidity: function getCreator(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) GetCreator(_id [32]byte) (common.Address, error) {
	return _DebtEngine.Contract.GetCreator(&_DebtEngine.CallOpts, _id)
}

// GetError is a free data retrieval call binding the contract method 0x945a6698.
//
// Solidity: function getError(bytes32 _id) constant returns(bool)
func (_DebtEngine *DebtEngineCaller) GetError(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getError", _id)
	return *ret0, err
}

// GetError is a free data retrieval call binding the contract method 0x945a6698.
//
// Solidity: function getError(bytes32 _id) constant returns(bool)
func (_DebtEngine *DebtEngineSession) GetError(_id [32]byte) (bool, error) {
	return _DebtEngine.Contract.GetError(&_DebtEngine.CallOpts, _id)
}

// GetError is a free data retrieval call binding the contract method 0x945a6698.
//
// Solidity: function getError(bytes32 _id) constant returns(bool)
func (_DebtEngine *DebtEngineCallerSession) GetError(_id [32]byte) (bool, error) {
	return _DebtEngine.Contract.GetError(&_DebtEngine.CallOpts, _id)
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineCaller) GetModel(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getModel", _id)
	return *ret0, err
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineSession) GetModel(_id [32]byte) (common.Address, error) {
	return _DebtEngine.Contract.GetModel(&_DebtEngine.CallOpts, _id)
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) GetModel(_id [32]byte) (common.Address, error) {
	return _DebtEngine.Contract.GetModel(&_DebtEngine.CallOpts, _id)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineCaller) GetOracle(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getOracle", _id)
	return *ret0, err
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineSession) GetOracle(_id [32]byte) (common.Address, error) {
	return _DebtEngine.Contract.GetOracle(&_DebtEngine.CallOpts, _id)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _id) constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) GetOracle(_id [32]byte) (common.Address, error) {
	return _DebtEngine.Contract.GetOracle(&_DebtEngine.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _id) constant returns(uint256)
func (_DebtEngine *DebtEngineCaller) GetStatus(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "getStatus", _id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _id) constant returns(uint256)
func (_DebtEngine *DebtEngineSession) GetStatus(_id [32]byte) (*big.Int, error) {
	return _DebtEngine.Contract.GetStatus(&_DebtEngine.CallOpts, _id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 _id) constant returns(uint256)
func (_DebtEngine *DebtEngineCallerSession) GetStatus(_id [32]byte) (*big.Int, error) {
	return _DebtEngine.Contract.GetStatus(&_DebtEngine.CallOpts, _id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _operator, address _assetHolder) constant returns(bool)
func (_DebtEngine *DebtEngineCaller) IsApprovedForAll(opts *bind.CallOpts, _operator common.Address, _assetHolder common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "isApprovedForAll", _operator, _assetHolder)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _operator, address _assetHolder) constant returns(bool)
func (_DebtEngine *DebtEngineSession) IsApprovedForAll(_operator common.Address, _assetHolder common.Address) (bool, error) {
	return _DebtEngine.Contract.IsApprovedForAll(&_DebtEngine.CallOpts, _operator, _assetHolder)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _operator, address _assetHolder) constant returns(bool)
func (_DebtEngine *DebtEngineCallerSession) IsApprovedForAll(_operator common.Address, _assetHolder common.Address) (bool, error) {
	return _DebtEngine.Contract.IsApprovedForAll(&_DebtEngine.CallOpts, _operator, _assetHolder)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x2972b0f0.
//
// Solidity: function isAuthorized(address _operator, uint256 _assetId) constant returns(bool)
func (_DebtEngine *DebtEngineCaller) IsAuthorized(opts *bind.CallOpts, _operator common.Address, _assetId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "isAuthorized", _operator, _assetId)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x2972b0f0.
//
// Solidity: function isAuthorized(address _operator, uint256 _assetId) constant returns(bool)
func (_DebtEngine *DebtEngineSession) IsAuthorized(_operator common.Address, _assetId *big.Int) (bool, error) {
	return _DebtEngine.Contract.IsAuthorized(&_DebtEngine.CallOpts, _operator, _assetId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x2972b0f0.
//
// Solidity: function isAuthorized(address _operator, uint256 _assetId) constant returns(bool)
func (_DebtEngine *DebtEngineCallerSession) IsAuthorized(_operator common.Address, _assetId *big.Int) (bool, error) {
	return _DebtEngine.Contract.IsAuthorized(&_DebtEngine.CallOpts, _operator, _assetId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DebtEngine *DebtEngineCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DebtEngine *DebtEngineSession) Name() (string, error) {
	return _DebtEngine.Contract.Name(&_DebtEngine.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DebtEngine *DebtEngineCallerSession) Name() (string, error) {
	return _DebtEngine.Contract.Name(&_DebtEngine.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) constant returns(uint256)
func (_DebtEngine *DebtEngineCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) constant returns(uint256)
func (_DebtEngine *DebtEngineSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DebtEngine.Contract.Nonces(&_DebtEngine.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) constant returns(uint256)
func (_DebtEngine *DebtEngineCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DebtEngine.Contract.Nonces(&_DebtEngine.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DebtEngine *DebtEngineCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DebtEngine *DebtEngineSession) Owner() (common.Address, error) {
	return _DebtEngine.Contract.Owner(&_DebtEngine.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) Owner() (common.Address, error) {
	return _DebtEngine.Contract.Owner(&_DebtEngine.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _assetId) constant returns(address)
func (_DebtEngine *DebtEngineCaller) OwnerOf(opts *bind.CallOpts, _assetId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "ownerOf", _assetId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _assetId) constant returns(address)
func (_DebtEngine *DebtEngineSession) OwnerOf(_assetId *big.Int) (common.Address, error) {
	return _DebtEngine.Contract.OwnerOf(&_DebtEngine.CallOpts, _assetId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _assetId) constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) OwnerOf(_assetId *big.Int) (common.Address, error) {
	return _DebtEngine.Contract.OwnerOf(&_DebtEngine.CallOpts, _assetId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DebtEngine *DebtEngineCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DebtEngine *DebtEngineSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DebtEngine.Contract.SupportsInterface(&_DebtEngine.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DebtEngine *DebtEngineCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DebtEngine.Contract.SupportsInterface(&_DebtEngine.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DebtEngine *DebtEngineCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DebtEngine *DebtEngineSession) Symbol() (string, error) {
	return _DebtEngine.Contract.Symbol(&_DebtEngine.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DebtEngine *DebtEngineCallerSession) Symbol() (string, error) {
	return _DebtEngine.Contract.Symbol(&_DebtEngine.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_DebtEngine *DebtEngineCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_DebtEngine *DebtEngineSession) Token() (common.Address, error) {
	return _DebtEngine.Contract.Token(&_DebtEngine.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_DebtEngine *DebtEngineCallerSession) Token() (common.Address, error) {
	return _DebtEngine.Contract.Token(&_DebtEngine.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_DebtEngine *DebtEngineCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_DebtEngine *DebtEngineSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _DebtEngine.Contract.TokenByIndex(&_DebtEngine.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) constant returns(uint256)
func (_DebtEngine *DebtEngineCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _DebtEngine.Contract.TokenByIndex(&_DebtEngine.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_DebtEngine *DebtEngineCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_DebtEngine *DebtEngineSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _DebtEngine.Contract.TokenOfOwnerByIndex(&_DebtEngine.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256)
func (_DebtEngine *DebtEngineCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _DebtEngine.Contract.TokenOfOwnerByIndex(&_DebtEngine.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_DebtEngine *DebtEngineCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_DebtEngine *DebtEngineSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DebtEngine.Contract.TokenURI(&_DebtEngine.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) constant returns(string)
func (_DebtEngine *DebtEngineCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DebtEngine.Contract.TokenURI(&_DebtEngine.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DebtEngine *DebtEngineCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DebtEngine.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DebtEngine *DebtEngineSession) TotalSupply() (*big.Int, error) {
	return _DebtEngine.Contract.TotalSupply(&_DebtEngine.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DebtEngine *DebtEngineCallerSession) TotalSupply() (*big.Int, error) {
	return _DebtEngine.Contract.TotalSupply(&_DebtEngine.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _operator, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineTransactor) Approve(opts *bind.TransactOpts, _operator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "approve", _operator, _assetId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _operator, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineSession) Approve(_operator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.Approve(&_DebtEngine.TransactOpts, _operator, _assetId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _operator, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineTransactorSession) Approve(_operator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.Approve(&_DebtEngine.TransactOpts, _operator, _assetId)
}

// Create is a paid mutator transaction binding the contract method 0x649b2b95.
//
// Solidity: function create(address _model, address _owner, address _oracle, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineTransactor) Create(opts *bind.TransactOpts, _model common.Address, _owner common.Address, _oracle common.Address, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "create", _model, _owner, _oracle, _data)
}

// Create is a paid mutator transaction binding the contract method 0x649b2b95.
//
// Solidity: function create(address _model, address _owner, address _oracle, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineSession) Create(_model common.Address, _owner common.Address, _oracle common.Address, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Create(&_DebtEngine.TransactOpts, _model, _owner, _oracle, _data)
}

// Create is a paid mutator transaction binding the contract method 0x649b2b95.
//
// Solidity: function create(address _model, address _owner, address _oracle, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineTransactorSession) Create(_model common.Address, _owner common.Address, _oracle common.Address, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Create(&_DebtEngine.TransactOpts, _model, _owner, _oracle, _data)
}

// Create2 is a paid mutator transaction binding the contract method 0x2ed04265.
//
// Solidity: function create2(address _model, address _owner, address _oracle, uint256 _salt, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineTransactor) Create2(opts *bind.TransactOpts, _model common.Address, _owner common.Address, _oracle common.Address, _salt *big.Int, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "create2", _model, _owner, _oracle, _salt, _data)
}

// Create2 is a paid mutator transaction binding the contract method 0x2ed04265.
//
// Solidity: function create2(address _model, address _owner, address _oracle, uint256 _salt, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineSession) Create2(_model common.Address, _owner common.Address, _oracle common.Address, _salt *big.Int, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Create2(&_DebtEngine.TransactOpts, _model, _owner, _oracle, _salt, _data)
}

// Create2 is a paid mutator transaction binding the contract method 0x2ed04265.
//
// Solidity: function create2(address _model, address _owner, address _oracle, uint256 _salt, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineTransactorSession) Create2(_model common.Address, _owner common.Address, _oracle common.Address, _salt *big.Int, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Create2(&_DebtEngine.TransactOpts, _model, _owner, _oracle, _salt, _data)
}

// Create3 is a paid mutator transaction binding the contract method 0xadc51d42.
//
// Solidity: function create3(address _model, address _owner, address _oracle, uint256 _salt, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineTransactor) Create3(opts *bind.TransactOpts, _model common.Address, _owner common.Address, _oracle common.Address, _salt *big.Int, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "create3", _model, _owner, _oracle, _salt, _data)
}

// Create3 is a paid mutator transaction binding the contract method 0xadc51d42.
//
// Solidity: function create3(address _model, address _owner, address _oracle, uint256 _salt, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineSession) Create3(_model common.Address, _owner common.Address, _oracle common.Address, _salt *big.Int, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Create3(&_DebtEngine.TransactOpts, _model, _owner, _oracle, _salt, _data)
}

// Create3 is a paid mutator transaction binding the contract method 0xadc51d42.
//
// Solidity: function create3(address _model, address _owner, address _oracle, uint256 _salt, bytes _data) returns(bytes32 id)
func (_DebtEngine *DebtEngineTransactorSession) Create3(_model common.Address, _owner common.Address, _oracle common.Address, _salt *big.Int, _data []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Create3(&_DebtEngine.TransactOpts, _model, _owner, _oracle, _salt, _data)
}

// Pay is a paid mutator transaction binding the contract method 0x114f9511.
//
// Solidity: function pay(bytes32 _id, uint256 _amount, address _origin, bytes _oracleData) returns(uint256 paid, uint256 paidToken)
func (_DebtEngine *DebtEngineTransactor) Pay(opts *bind.TransactOpts, _id [32]byte, _amount *big.Int, _origin common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "pay", _id, _amount, _origin, _oracleData)
}

// Pay is a paid mutator transaction binding the contract method 0x114f9511.
//
// Solidity: function pay(bytes32 _id, uint256 _amount, address _origin, bytes _oracleData) returns(uint256 paid, uint256 paidToken)
func (_DebtEngine *DebtEngineSession) Pay(_id [32]byte, _amount *big.Int, _origin common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Pay(&_DebtEngine.TransactOpts, _id, _amount, _origin, _oracleData)
}

// Pay is a paid mutator transaction binding the contract method 0x114f9511.
//
// Solidity: function pay(bytes32 _id, uint256 _amount, address _origin, bytes _oracleData) returns(uint256 paid, uint256 paidToken)
func (_DebtEngine *DebtEngineTransactorSession) Pay(_id [32]byte, _amount *big.Int, _origin common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Pay(&_DebtEngine.TransactOpts, _id, _amount, _origin, _oracleData)
}

// PayBatch is a paid mutator transaction binding the contract method 0xd2d791c0.
//
// Solidity: function payBatch(bytes32[] _ids, uint256[] _amounts, address _origin, address _oracle, bytes _oracleData) returns(uint256[] paid, uint256[] paidTokens)
func (_DebtEngine *DebtEngineTransactor) PayBatch(opts *bind.TransactOpts, _ids [][32]byte, _amounts []*big.Int, _origin common.Address, _oracle common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "payBatch", _ids, _amounts, _origin, _oracle, _oracleData)
}

// PayBatch is a paid mutator transaction binding the contract method 0xd2d791c0.
//
// Solidity: function payBatch(bytes32[] _ids, uint256[] _amounts, address _origin, address _oracle, bytes _oracleData) returns(uint256[] paid, uint256[] paidTokens)
func (_DebtEngine *DebtEngineSession) PayBatch(_ids [][32]byte, _amounts []*big.Int, _origin common.Address, _oracle common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.PayBatch(&_DebtEngine.TransactOpts, _ids, _amounts, _origin, _oracle, _oracleData)
}

// PayBatch is a paid mutator transaction binding the contract method 0xd2d791c0.
//
// Solidity: function payBatch(bytes32[] _ids, uint256[] _amounts, address _origin, address _oracle, bytes _oracleData) returns(uint256[] paid, uint256[] paidTokens)
func (_DebtEngine *DebtEngineTransactorSession) PayBatch(_ids [][32]byte, _amounts []*big.Int, _origin common.Address, _oracle common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.PayBatch(&_DebtEngine.TransactOpts, _ids, _amounts, _origin, _oracle, _oracleData)
}

// PayToken is a paid mutator transaction binding the contract method 0x7c1716af.
//
// Solidity: function payToken(bytes32 id, uint256 amount, address origin, bytes oracleData) returns(uint256 paid, uint256 paidToken)
func (_DebtEngine *DebtEngineTransactor) PayToken(opts *bind.TransactOpts, id [32]byte, amount *big.Int, origin common.Address, oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "payToken", id, amount, origin, oracleData)
}

// PayToken is a paid mutator transaction binding the contract method 0x7c1716af.
//
// Solidity: function payToken(bytes32 id, uint256 amount, address origin, bytes oracleData) returns(uint256 paid, uint256 paidToken)
func (_DebtEngine *DebtEngineSession) PayToken(id [32]byte, amount *big.Int, origin common.Address, oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.PayToken(&_DebtEngine.TransactOpts, id, amount, origin, oracleData)
}

// PayToken is a paid mutator transaction binding the contract method 0x7c1716af.
//
// Solidity: function payToken(bytes32 id, uint256 amount, address origin, bytes oracleData) returns(uint256 paid, uint256 paidToken)
func (_DebtEngine *DebtEngineTransactorSession) PayToken(id [32]byte, amount *big.Int, origin common.Address, oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.PayToken(&_DebtEngine.TransactOpts, id, amount, origin, oracleData)
}

// PayTokenBatch is a paid mutator transaction binding the contract method 0x4d9c1bf3.
//
// Solidity: function payTokenBatch(bytes32[] _ids, uint256[] _tokenAmounts, address _origin, address _oracle, bytes _oracleData) returns(uint256[] paid, uint256[] paidTokens)
func (_DebtEngine *DebtEngineTransactor) PayTokenBatch(opts *bind.TransactOpts, _ids [][32]byte, _tokenAmounts []*big.Int, _origin common.Address, _oracle common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "payTokenBatch", _ids, _tokenAmounts, _origin, _oracle, _oracleData)
}

// PayTokenBatch is a paid mutator transaction binding the contract method 0x4d9c1bf3.
//
// Solidity: function payTokenBatch(bytes32[] _ids, uint256[] _tokenAmounts, address _origin, address _oracle, bytes _oracleData) returns(uint256[] paid, uint256[] paidTokens)
func (_DebtEngine *DebtEngineSession) PayTokenBatch(_ids [][32]byte, _tokenAmounts []*big.Int, _origin common.Address, _oracle common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.PayTokenBatch(&_DebtEngine.TransactOpts, _ids, _tokenAmounts, _origin, _oracle, _oracleData)
}

// PayTokenBatch is a paid mutator transaction binding the contract method 0x4d9c1bf3.
//
// Solidity: function payTokenBatch(bytes32[] _ids, uint256[] _tokenAmounts, address _origin, address _oracle, bytes _oracleData) returns(uint256[] paid, uint256[] paidTokens)
func (_DebtEngine *DebtEngineTransactorSession) PayTokenBatch(_ids [][32]byte, _tokenAmounts []*big.Int, _origin common.Address, _oracle common.Address, _oracleData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.PayTokenBatch(&_DebtEngine.TransactOpts, _ids, _tokenAmounts, _origin, _oracle, _oracleData)
}

// Run is a paid mutator transaction binding the contract method 0xef6ac0f0.
//
// Solidity: function run(bytes32 _id) returns(bool)
func (_DebtEngine *DebtEngineTransactor) Run(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "run", _id)
}

// Run is a paid mutator transaction binding the contract method 0xef6ac0f0.
//
// Solidity: function run(bytes32 _id) returns(bool)
func (_DebtEngine *DebtEngineSession) Run(_id [32]byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Run(&_DebtEngine.TransactOpts, _id)
}

// Run is a paid mutator transaction binding the contract method 0xef6ac0f0.
//
// Solidity: function run(bytes32 _id) returns(bool)
func (_DebtEngine *DebtEngineTransactorSession) Run(_id [32]byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.Run(&_DebtEngine.TransactOpts, _id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "safeTransferFrom", _from, _to, _assetId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineSession) SafeTransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.SafeTransferFrom(&_DebtEngine.TransactOpts, _from, _to, _assetId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.SafeTransferFrom(&_DebtEngine.TransactOpts, _from, _to, _assetId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId, bytes _userData) returns()
func (_DebtEngine *DebtEngineTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _assetId *big.Int, _userData []byte) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "safeTransferFrom0", _from, _to, _assetId, _userData)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId, bytes _userData) returns()
func (_DebtEngine *DebtEngineSession) SafeTransferFrom0(_from common.Address, _to common.Address, _assetId *big.Int, _userData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.SafeTransferFrom0(&_DebtEngine.TransactOpts, _from, _to, _assetId, _userData)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _assetId, bytes _userData) returns()
func (_DebtEngine *DebtEngineTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _assetId *big.Int, _userData []byte) (*types.Transaction, error) {
	return _DebtEngine.Contract.SafeTransferFrom0(&_DebtEngine.TransactOpts, _from, _to, _assetId, _userData)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _authorized) returns()
func (_DebtEngine *DebtEngineTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _authorized bool) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "setApprovalForAll", _operator, _authorized)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _authorized) returns()
func (_DebtEngine *DebtEngineSession) SetApprovalForAll(_operator common.Address, _authorized bool) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetApprovalForAll(&_DebtEngine.TransactOpts, _operator, _authorized)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _authorized) returns()
func (_DebtEngine *DebtEngineTransactorSession) SetApprovalForAll(_operator common.Address, _authorized bool) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetApprovalForAll(&_DebtEngine.TransactOpts, _operator, _authorized)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_DebtEngine *DebtEngineTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_DebtEngine *DebtEngineSession) SetName(_name string) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetName(&_DebtEngine.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_DebtEngine *DebtEngineTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetName(&_DebtEngine.TransactOpts, _name)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_DebtEngine *DebtEngineTransactor) SetSymbol(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "setSymbol", _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_DebtEngine *DebtEngineSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetSymbol(&_DebtEngine.TransactOpts, _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_DebtEngine *DebtEngineTransactorSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetSymbol(&_DebtEngine.TransactOpts, _symbol)
}

// SetURIProvider is a paid mutator transaction binding the contract method 0xdb2c5518.
//
// Solidity: function setURIProvider(address _provider) returns()
func (_DebtEngine *DebtEngineTransactor) SetURIProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "setURIProvider", _provider)
}

// SetURIProvider is a paid mutator transaction binding the contract method 0xdb2c5518.
//
// Solidity: function setURIProvider(address _provider) returns()
func (_DebtEngine *DebtEngineSession) SetURIProvider(_provider common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetURIProvider(&_DebtEngine.TransactOpts, _provider)
}

// SetURIProvider is a paid mutator transaction binding the contract method 0xdb2c5518.
//
// Solidity: function setURIProvider(address _provider) returns()
func (_DebtEngine *DebtEngineTransactorSession) SetURIProvider(_provider common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.SetURIProvider(&_DebtEngine.TransactOpts, _provider)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "transferFrom", _from, _to, _assetId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineSession) TransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.TransferFrom(&_DebtEngine.TransactOpts, _from, _to, _assetId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _assetId) returns()
func (_DebtEngine *DebtEngineTransactorSession) TransferFrom(_from common.Address, _to common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.TransferFrom(&_DebtEngine.TransactOpts, _from, _to, _assetId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_DebtEngine *DebtEngineTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_DebtEngine *DebtEngineSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.TransferOwnership(&_DebtEngine.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_DebtEngine *DebtEngineTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.TransferOwnership(&_DebtEngine.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1b258d50.
//
// Solidity: function withdraw(bytes32 _id, address _to) returns(uint256 amount)
func (_DebtEngine *DebtEngineTransactor) Withdraw(opts *bind.TransactOpts, _id [32]byte, _to common.Address) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "withdraw", _id, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1b258d50.
//
// Solidity: function withdraw(bytes32 _id, address _to) returns(uint256 amount)
func (_DebtEngine *DebtEngineSession) Withdraw(_id [32]byte, _to common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.Withdraw(&_DebtEngine.TransactOpts, _id, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1b258d50.
//
// Solidity: function withdraw(bytes32 _id, address _to) returns(uint256 amount)
func (_DebtEngine *DebtEngineTransactorSession) Withdraw(_id [32]byte, _to common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.Withdraw(&_DebtEngine.TransactOpts, _id, _to)
}

// WithdrawBatch is a paid mutator transaction binding the contract method 0xf1637630.
//
// Solidity: function withdrawBatch(bytes32[] _ids, address _to) returns(uint256 total)
func (_DebtEngine *DebtEngineTransactor) WithdrawBatch(opts *bind.TransactOpts, _ids [][32]byte, _to common.Address) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "withdrawBatch", _ids, _to)
}

// WithdrawBatch is a paid mutator transaction binding the contract method 0xf1637630.
//
// Solidity: function withdrawBatch(bytes32[] _ids, address _to) returns(uint256 total)
func (_DebtEngine *DebtEngineSession) WithdrawBatch(_ids [][32]byte, _to common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.WithdrawBatch(&_DebtEngine.TransactOpts, _ids, _to)
}

// WithdrawBatch is a paid mutator transaction binding the contract method 0xf1637630.
//
// Solidity: function withdrawBatch(bytes32[] _ids, address _to) returns(uint256 total)
func (_DebtEngine *DebtEngineTransactorSession) WithdrawBatch(_ids [][32]byte, _to common.Address) (*types.Transaction, error) {
	return _DebtEngine.Contract.WithdrawBatch(&_DebtEngine.TransactOpts, _ids, _to)
}

// WithdrawPartial is a paid mutator transaction binding the contract method 0xf12e5eaa.
//
// Solidity: function withdrawPartial(bytes32 _id, address _to, uint256 _amount) returns(bool success)
func (_DebtEngine *DebtEngineTransactor) WithdrawPartial(opts *bind.TransactOpts, _id [32]byte, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtEngine.contract.Transact(opts, "withdrawPartial", _id, _to, _amount)
}

// WithdrawPartial is a paid mutator transaction binding the contract method 0xf12e5eaa.
//
// Solidity: function withdrawPartial(bytes32 _id, address _to, uint256 _amount) returns(bool success)
func (_DebtEngine *DebtEngineSession) WithdrawPartial(_id [32]byte, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.WithdrawPartial(&_DebtEngine.TransactOpts, _id, _to, _amount)
}

// WithdrawPartial is a paid mutator transaction binding the contract method 0xf12e5eaa.
//
// Solidity: function withdrawPartial(bytes32 _id, address _to, uint256 _amount) returns(bool success)
func (_DebtEngine *DebtEngineTransactorSession) WithdrawPartial(_id [32]byte, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtEngine.Contract.WithdrawPartial(&_DebtEngine.TransactOpts, _id, _to, _amount)
}

// DebtEngineApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DebtEngine contract.
type DebtEngineApprovalIterator struct {
	Event *DebtEngineApproval // Event containing the contract specifics and raw log

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
func (it *DebtEngineApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineApproval)
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
		it.Event = new(DebtEngineApproval)
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
func (it *DebtEngineApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineApproval represents a Approval event raised by the DebtEngine contract.
type DebtEngineApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_DebtEngine *DebtEngineFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*DebtEngineApprovalIterator, error) {

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

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineApprovalIterator{contract: _DebtEngine.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_DebtEngine *DebtEngineFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DebtEngineApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineApproval)
				if err := _DebtEngine.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DebtEngine *DebtEngineFilterer) ParseApproval(log types.Log) (*DebtEngineApproval, error) {
	event := new(DebtEngineApproval)
	if err := _DebtEngine.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DebtEngine contract.
type DebtEngineApprovalForAllIterator struct {
	Event *DebtEngineApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DebtEngineApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineApprovalForAll)
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
		it.Event = new(DebtEngineApprovalForAll)
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
func (it *DebtEngineApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineApprovalForAll represents a ApprovalForAll event raised by the DebtEngine contract.
type DebtEngineApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DebtEngine *DebtEngineFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*DebtEngineApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineApprovalForAllIterator{contract: _DebtEngine.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DebtEngine *DebtEngineFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DebtEngineApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineApprovalForAll)
				if err := _DebtEngine.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_DebtEngine *DebtEngineFilterer) ParseApprovalForAll(log types.Log) (*DebtEngineApprovalForAll, error) {
	event := new(DebtEngineApprovalForAll)
	if err := _DebtEngine.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the DebtEngine contract.
type DebtEngineCreatedIterator struct {
	Event *DebtEngineCreated // Event containing the contract specifics and raw log

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
func (it *DebtEngineCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineCreated)
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
		it.Event = new(DebtEngineCreated)
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
func (it *DebtEngineCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineCreated represents a Created event raised by the DebtEngine contract.
type DebtEngineCreated struct {
	Id    [32]byte
	Nonce *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0xc9310bea14cefb8c6d9a8fce4ead7f6825a31223af9d0bde2a30b2d0c71ccbf3.
//
// Solidity: event Created(bytes32 indexed _id, uint256 _nonce, bytes _data)
func (_DebtEngine *DebtEngineFilterer) FilterCreated(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineCreatedIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Created", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineCreatedIterator{contract: _DebtEngine.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0xc9310bea14cefb8c6d9a8fce4ead7f6825a31223af9d0bde2a30b2d0c71ccbf3.
//
// Solidity: event Created(bytes32 indexed _id, uint256 _nonce, bytes _data)
func (_DebtEngine *DebtEngineFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *DebtEngineCreated, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Created", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineCreated)
				if err := _DebtEngine.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0xc9310bea14cefb8c6d9a8fce4ead7f6825a31223af9d0bde2a30b2d0c71ccbf3.
//
// Solidity: event Created(bytes32 indexed _id, uint256 _nonce, bytes _data)
func (_DebtEngine *DebtEngineFilterer) ParseCreated(log types.Log) (*DebtEngineCreated, error) {
	event := new(DebtEngineCreated)
	if err := _DebtEngine.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineCreated2Iterator is returned from FilterCreated2 and is used to iterate over the raw logs and unpacked data for Created2 events raised by the DebtEngine contract.
type DebtEngineCreated2Iterator struct {
	Event *DebtEngineCreated2 // Event containing the contract specifics and raw log

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
func (it *DebtEngineCreated2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineCreated2)
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
		it.Event = new(DebtEngineCreated2)
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
func (it *DebtEngineCreated2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineCreated2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineCreated2 represents a Created2 event raised by the DebtEngine contract.
type DebtEngineCreated2 struct {
	Id   [32]byte
	Salt *big.Int
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCreated2 is a free log retrieval operation binding the contract event 0x4ee48d90c05f58fc51a05728e403947a51e36cc7a1bf4d82260473b433538408.
//
// Solidity: event Created2(bytes32 indexed _id, uint256 _salt, bytes _data)
func (_DebtEngine *DebtEngineFilterer) FilterCreated2(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineCreated2Iterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Created2", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineCreated2Iterator{contract: _DebtEngine.contract, event: "Created2", logs: logs, sub: sub}, nil
}

// WatchCreated2 is a free log subscription operation binding the contract event 0x4ee48d90c05f58fc51a05728e403947a51e36cc7a1bf4d82260473b433538408.
//
// Solidity: event Created2(bytes32 indexed _id, uint256 _salt, bytes _data)
func (_DebtEngine *DebtEngineFilterer) WatchCreated2(opts *bind.WatchOpts, sink chan<- *DebtEngineCreated2, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Created2", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineCreated2)
				if err := _DebtEngine.contract.UnpackLog(event, "Created2", log); err != nil {
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

// ParseCreated2 is a log parse operation binding the contract event 0x4ee48d90c05f58fc51a05728e403947a51e36cc7a1bf4d82260473b433538408.
//
// Solidity: event Created2(bytes32 indexed _id, uint256 _salt, bytes _data)
func (_DebtEngine *DebtEngineFilterer) ParseCreated2(log types.Log) (*DebtEngineCreated2, error) {
	event := new(DebtEngineCreated2)
	if err := _DebtEngine.contract.UnpackLog(event, "Created2", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineCreated3Iterator is returned from FilterCreated3 and is used to iterate over the raw logs and unpacked data for Created3 events raised by the DebtEngine contract.
type DebtEngineCreated3Iterator struct {
	Event *DebtEngineCreated3 // Event containing the contract specifics and raw log

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
func (it *DebtEngineCreated3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineCreated3)
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
		it.Event = new(DebtEngineCreated3)
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
func (it *DebtEngineCreated3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineCreated3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineCreated3 represents a Created3 event raised by the DebtEngine contract.
type DebtEngineCreated3 struct {
	Id   [32]byte
	Salt *big.Int
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCreated3 is a free log retrieval operation binding the contract event 0xc46d3d245e649272659dd5731879c1b019caccb00acb3a94488a28c82ef924a6.
//
// Solidity: event Created3(bytes32 indexed _id, uint256 _salt, bytes _data)
func (_DebtEngine *DebtEngineFilterer) FilterCreated3(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineCreated3Iterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Created3", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineCreated3Iterator{contract: _DebtEngine.contract, event: "Created3", logs: logs, sub: sub}, nil
}

// WatchCreated3 is a free log subscription operation binding the contract event 0xc46d3d245e649272659dd5731879c1b019caccb00acb3a94488a28c82ef924a6.
//
// Solidity: event Created3(bytes32 indexed _id, uint256 _salt, bytes _data)
func (_DebtEngine *DebtEngineFilterer) WatchCreated3(opts *bind.WatchOpts, sink chan<- *DebtEngineCreated3, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Created3", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineCreated3)
				if err := _DebtEngine.contract.UnpackLog(event, "Created3", log); err != nil {
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

// ParseCreated3 is a log parse operation binding the contract event 0xc46d3d245e649272659dd5731879c1b019caccb00acb3a94488a28c82ef924a6.
//
// Solidity: event Created3(bytes32 indexed _id, uint256 _salt, bytes _data)
func (_DebtEngine *DebtEngineFilterer) ParseCreated3(log types.Log) (*DebtEngineCreated3, error) {
	event := new(DebtEngineCreated3)
	if err := _DebtEngine.contract.UnpackLog(event, "Created3", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the DebtEngine contract.
type DebtEngineErrorIterator struct {
	Event *DebtEngineError // Event containing the contract specifics and raw log

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
func (it *DebtEngineErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineError)
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
		it.Event = new(DebtEngineError)
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
func (it *DebtEngineErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineError represents a Error event raised by the DebtEngine contract.
type DebtEngineError struct {
	Id       [32]byte
	Sender   common.Address
	Value    *big.Int
	GasLeft  *big.Int
	GasLimit *big.Int
	CallData []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0xe000ee8a04c9b1ff738cb5c90d6a362bf5be3c83b1eb221377df6324747d4c07.
//
// Solidity: event Error(bytes32 indexed _id, address _sender, uint256 _value, uint256 _gasLeft, uint256 _gasLimit, bytes _callData)
func (_DebtEngine *DebtEngineFilterer) FilterError(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineErrorIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Error", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineErrorIterator{contract: _DebtEngine.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0xe000ee8a04c9b1ff738cb5c90d6a362bf5be3c83b1eb221377df6324747d4c07.
//
// Solidity: event Error(bytes32 indexed _id, address _sender, uint256 _value, uint256 _gasLeft, uint256 _gasLimit, bytes _callData)
func (_DebtEngine *DebtEngineFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *DebtEngineError, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Error", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineError)
				if err := _DebtEngine.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0xe000ee8a04c9b1ff738cb5c90d6a362bf5be3c83b1eb221377df6324747d4c07.
//
// Solidity: event Error(bytes32 indexed _id, address _sender, uint256 _value, uint256 _gasLeft, uint256 _gasLimit, bytes _callData)
func (_DebtEngine *DebtEngineFilterer) ParseError(log types.Log) (*DebtEngineError, error) {
	event := new(DebtEngineError)
	if err := _DebtEngine.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineErrorRecoverIterator is returned from FilterErrorRecover and is used to iterate over the raw logs and unpacked data for ErrorRecover events raised by the DebtEngine contract.
type DebtEngineErrorRecoverIterator struct {
	Event *DebtEngineErrorRecover // Event containing the contract specifics and raw log

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
func (it *DebtEngineErrorRecoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineErrorRecover)
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
		it.Event = new(DebtEngineErrorRecover)
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
func (it *DebtEngineErrorRecoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineErrorRecoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineErrorRecover represents a ErrorRecover event raised by the DebtEngine contract.
type DebtEngineErrorRecover struct {
	Id       [32]byte
	Sender   common.Address
	Value    *big.Int
	GasLeft  *big.Int
	GasLimit *big.Int
	Result   [32]byte
	CallData []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterErrorRecover is a free log retrieval operation binding the contract event 0xa43b78c6d135806f3e817a90a7b690c0d2d94b7d2e356049fa5ba357c0fab9c3.
//
// Solidity: event ErrorRecover(bytes32 indexed _id, address _sender, uint256 _value, uint256 _gasLeft, uint256 _gasLimit, bytes32 _result, bytes _callData)
func (_DebtEngine *DebtEngineFilterer) FilterErrorRecover(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineErrorRecoverIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "ErrorRecover", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineErrorRecoverIterator{contract: _DebtEngine.contract, event: "ErrorRecover", logs: logs, sub: sub}, nil
}

// WatchErrorRecover is a free log subscription operation binding the contract event 0xa43b78c6d135806f3e817a90a7b690c0d2d94b7d2e356049fa5ba357c0fab9c3.
//
// Solidity: event ErrorRecover(bytes32 indexed _id, address _sender, uint256 _value, uint256 _gasLeft, uint256 _gasLimit, bytes32 _result, bytes _callData)
func (_DebtEngine *DebtEngineFilterer) WatchErrorRecover(opts *bind.WatchOpts, sink chan<- *DebtEngineErrorRecover, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "ErrorRecover", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineErrorRecover)
				if err := _DebtEngine.contract.UnpackLog(event, "ErrorRecover", log); err != nil {
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

// ParseErrorRecover is a log parse operation binding the contract event 0xa43b78c6d135806f3e817a90a7b690c0d2d94b7d2e356049fa5ba357c0fab9c3.
//
// Solidity: event ErrorRecover(bytes32 indexed _id, address _sender, uint256 _value, uint256 _gasLeft, uint256 _gasLimit, bytes32 _result, bytes _callData)
func (_DebtEngine *DebtEngineFilterer) ParseErrorRecover(log types.Log) (*DebtEngineErrorRecover, error) {
	event := new(DebtEngineErrorRecover)
	if err := _DebtEngine.contract.UnpackLog(event, "ErrorRecover", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DebtEngine contract.
type DebtEngineOwnershipTransferredIterator struct {
	Event *DebtEngineOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DebtEngineOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineOwnershipTransferred)
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
		it.Event = new(DebtEngineOwnershipTransferred)
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
func (it *DebtEngineOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineOwnershipTransferred represents a OwnershipTransferred event raised by the DebtEngine contract.
type DebtEngineOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_DebtEngine *DebtEngineFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*DebtEngineOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineOwnershipTransferredIterator{contract: _DebtEngine.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_DebtEngine *DebtEngineFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DebtEngineOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineOwnershipTransferred)
				if err := _DebtEngine.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DebtEngine *DebtEngineFilterer) ParseOwnershipTransferred(log types.Log) (*DebtEngineOwnershipTransferred, error) {
	event := new(DebtEngineOwnershipTransferred)
	if err := _DebtEngine.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEnginePaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the DebtEngine contract.
type DebtEnginePaidIterator struct {
	Event *DebtEnginePaid // Event containing the contract specifics and raw log

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
func (it *DebtEnginePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEnginePaid)
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
		it.Event = new(DebtEnginePaid)
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
func (it *DebtEnginePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEnginePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEnginePaid represents a Paid event raised by the DebtEngine contract.
type DebtEnginePaid struct {
	Id              [32]byte
	Sender          common.Address
	Origin          common.Address
	Requested       *big.Int
	RequestedTokens *big.Int
	Paid            *big.Int
	Tokens          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x25b52320bc27b845f37eae2240cc285c7b6e5643fc2995e6d22afa10e2f657d2.
//
// Solidity: event Paid(bytes32 indexed _id, address _sender, address _origin, uint256 _requested, uint256 _requestedTokens, uint256 _paid, uint256 _tokens)
func (_DebtEngine *DebtEngineFilterer) FilterPaid(opts *bind.FilterOpts, _id [][32]byte) (*DebtEnginePaidIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Paid", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEnginePaidIterator{contract: _DebtEngine.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x25b52320bc27b845f37eae2240cc285c7b6e5643fc2995e6d22afa10e2f657d2.
//
// Solidity: event Paid(bytes32 indexed _id, address _sender, address _origin, uint256 _requested, uint256 _requestedTokens, uint256 _paid, uint256 _tokens)
func (_DebtEngine *DebtEngineFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *DebtEnginePaid, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Paid", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEnginePaid)
				if err := _DebtEngine.contract.UnpackLog(event, "Paid", log); err != nil {
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

// ParsePaid is a log parse operation binding the contract event 0x25b52320bc27b845f37eae2240cc285c7b6e5643fc2995e6d22afa10e2f657d2.
//
// Solidity: event Paid(bytes32 indexed _id, address _sender, address _origin, uint256 _requested, uint256 _requestedTokens, uint256 _paid, uint256 _tokens)
func (_DebtEngine *DebtEngineFilterer) ParsePaid(log types.Log) (*DebtEnginePaid, error) {
	event := new(DebtEnginePaid)
	if err := _DebtEngine.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEnginePayBatchErrorIterator is returned from FilterPayBatchError and is used to iterate over the raw logs and unpacked data for PayBatchError events raised by the DebtEngine contract.
type DebtEnginePayBatchErrorIterator struct {
	Event *DebtEnginePayBatchError // Event containing the contract specifics and raw log

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
func (it *DebtEnginePayBatchErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEnginePayBatchError)
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
		it.Event = new(DebtEnginePayBatchError)
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
func (it *DebtEnginePayBatchErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEnginePayBatchErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEnginePayBatchError represents a PayBatchError event raised by the DebtEngine contract.
type DebtEnginePayBatchError struct {
	Id           [32]byte
	TargetOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPayBatchError is a free log retrieval operation binding the contract event 0x8a1ca4383d0682548257effd0249245baaabe2c4672c83c48397a88faa0831f0.
//
// Solidity: event PayBatchError(bytes32 indexed _id, address _targetOracle)
func (_DebtEngine *DebtEngineFilterer) FilterPayBatchError(opts *bind.FilterOpts, _id [][32]byte) (*DebtEnginePayBatchErrorIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "PayBatchError", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEnginePayBatchErrorIterator{contract: _DebtEngine.contract, event: "PayBatchError", logs: logs, sub: sub}, nil
}

// WatchPayBatchError is a free log subscription operation binding the contract event 0x8a1ca4383d0682548257effd0249245baaabe2c4672c83c48397a88faa0831f0.
//
// Solidity: event PayBatchError(bytes32 indexed _id, address _targetOracle)
func (_DebtEngine *DebtEngineFilterer) WatchPayBatchError(opts *bind.WatchOpts, sink chan<- *DebtEnginePayBatchError, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "PayBatchError", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEnginePayBatchError)
				if err := _DebtEngine.contract.UnpackLog(event, "PayBatchError", log); err != nil {
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

// ParsePayBatchError is a log parse operation binding the contract event 0x8a1ca4383d0682548257effd0249245baaabe2c4672c83c48397a88faa0831f0.
//
// Solidity: event PayBatchError(bytes32 indexed _id, address _targetOracle)
func (_DebtEngine *DebtEngineFilterer) ParsePayBatchError(log types.Log) (*DebtEnginePayBatchError, error) {
	event := new(DebtEnginePayBatchError)
	if err := _DebtEngine.contract.UnpackLog(event, "PayBatchError", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineReadedOracleIterator is returned from FilterReadedOracle and is used to iterate over the raw logs and unpacked data for ReadedOracle events raised by the DebtEngine contract.
type DebtEngineReadedOracleIterator struct {
	Event *DebtEngineReadedOracle // Event containing the contract specifics and raw log

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
func (it *DebtEngineReadedOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineReadedOracle)
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
		it.Event = new(DebtEngineReadedOracle)
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
func (it *DebtEngineReadedOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineReadedOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineReadedOracle represents a ReadedOracle event raised by the DebtEngine contract.
type DebtEngineReadedOracle struct {
	Id         [32]byte
	Tokens     *big.Int
	Equivalent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReadedOracle is a free log retrieval operation binding the contract event 0x686ef5859f39efd8bf5a57659ca5b60fe18f984d435a0b7418f18bcbe2d349d6.
//
// Solidity: event ReadedOracle(bytes32 indexed _id, uint256 _tokens, uint256 _equivalent)
func (_DebtEngine *DebtEngineFilterer) FilterReadedOracle(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineReadedOracleIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "ReadedOracle", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineReadedOracleIterator{contract: _DebtEngine.contract, event: "ReadedOracle", logs: logs, sub: sub}, nil
}

// WatchReadedOracle is a free log subscription operation binding the contract event 0x686ef5859f39efd8bf5a57659ca5b60fe18f984d435a0b7418f18bcbe2d349d6.
//
// Solidity: event ReadedOracle(bytes32 indexed _id, uint256 _tokens, uint256 _equivalent)
func (_DebtEngine *DebtEngineFilterer) WatchReadedOracle(opts *bind.WatchOpts, sink chan<- *DebtEngineReadedOracle, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "ReadedOracle", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineReadedOracle)
				if err := _DebtEngine.contract.UnpackLog(event, "ReadedOracle", log); err != nil {
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

// ParseReadedOracle is a log parse operation binding the contract event 0x686ef5859f39efd8bf5a57659ca5b60fe18f984d435a0b7418f18bcbe2d349d6.
//
// Solidity: event ReadedOracle(bytes32 indexed _id, uint256 _tokens, uint256 _equivalent)
func (_DebtEngine *DebtEngineFilterer) ParseReadedOracle(log types.Log) (*DebtEngineReadedOracle, error) {
	event := new(DebtEngineReadedOracle)
	if err := _DebtEngine.contract.UnpackLog(event, "ReadedOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineReadedOracleBatchIterator is returned from FilterReadedOracleBatch and is used to iterate over the raw logs and unpacked data for ReadedOracleBatch events raised by the DebtEngine contract.
type DebtEngineReadedOracleBatchIterator struct {
	Event *DebtEngineReadedOracleBatch // Event containing the contract specifics and raw log

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
func (it *DebtEngineReadedOracleBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineReadedOracleBatch)
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
		it.Event = new(DebtEngineReadedOracleBatch)
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
func (it *DebtEngineReadedOracleBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineReadedOracleBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineReadedOracleBatch represents a ReadedOracleBatch event raised by the DebtEngine contract.
type DebtEngineReadedOracleBatch struct {
	Oracle     common.Address
	Count      *big.Int
	Tokens     *big.Int
	Equivalent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReadedOracleBatch is a free log retrieval operation binding the contract event 0x79601cc1ca8f1ff7c3e7f7b522c2f1377cdb4c318a131afaf426cb2c976ef4be.
//
// Solidity: event ReadedOracleBatch(address _oracle, uint256 _count, uint256 _tokens, uint256 _equivalent)
func (_DebtEngine *DebtEngineFilterer) FilterReadedOracleBatch(opts *bind.FilterOpts) (*DebtEngineReadedOracleBatchIterator, error) {

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "ReadedOracleBatch")
	if err != nil {
		return nil, err
	}
	return &DebtEngineReadedOracleBatchIterator{contract: _DebtEngine.contract, event: "ReadedOracleBatch", logs: logs, sub: sub}, nil
}

// WatchReadedOracleBatch is a free log subscription operation binding the contract event 0x79601cc1ca8f1ff7c3e7f7b522c2f1377cdb4c318a131afaf426cb2c976ef4be.
//
// Solidity: event ReadedOracleBatch(address _oracle, uint256 _count, uint256 _tokens, uint256 _equivalent)
func (_DebtEngine *DebtEngineFilterer) WatchReadedOracleBatch(opts *bind.WatchOpts, sink chan<- *DebtEngineReadedOracleBatch) (event.Subscription, error) {

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "ReadedOracleBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineReadedOracleBatch)
				if err := _DebtEngine.contract.UnpackLog(event, "ReadedOracleBatch", log); err != nil {
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

// ParseReadedOracleBatch is a log parse operation binding the contract event 0x79601cc1ca8f1ff7c3e7f7b522c2f1377cdb4c318a131afaf426cb2c976ef4be.
//
// Solidity: event ReadedOracleBatch(address _oracle, uint256 _count, uint256 _tokens, uint256 _equivalent)
func (_DebtEngine *DebtEngineFilterer) ParseReadedOracleBatch(log types.Log) (*DebtEngineReadedOracleBatch, error) {
	event := new(DebtEngineReadedOracleBatch)
	if err := _DebtEngine.contract.UnpackLog(event, "ReadedOracleBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineSetNameIterator is returned from FilterSetName and is used to iterate over the raw logs and unpacked data for SetName events raised by the DebtEngine contract.
type DebtEngineSetNameIterator struct {
	Event *DebtEngineSetName // Event containing the contract specifics and raw log

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
func (it *DebtEngineSetNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineSetName)
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
		it.Event = new(DebtEngineSetName)
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
func (it *DebtEngineSetNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineSetNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineSetName represents a SetName event raised by the DebtEngine contract.
type DebtEngineSetName struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetName is a free log retrieval operation binding the contract event 0x4df9dcd34ae35f40f2c756fd8ac83210ed0b76d065543ee73d868aec7c7fcf02.
//
// Solidity: event SetName(string _name)
func (_DebtEngine *DebtEngineFilterer) FilterSetName(opts *bind.FilterOpts) (*DebtEngineSetNameIterator, error) {

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "SetName")
	if err != nil {
		return nil, err
	}
	return &DebtEngineSetNameIterator{contract: _DebtEngine.contract, event: "SetName", logs: logs, sub: sub}, nil
}

// WatchSetName is a free log subscription operation binding the contract event 0x4df9dcd34ae35f40f2c756fd8ac83210ed0b76d065543ee73d868aec7c7fcf02.
//
// Solidity: event SetName(string _name)
func (_DebtEngine *DebtEngineFilterer) WatchSetName(opts *bind.WatchOpts, sink chan<- *DebtEngineSetName) (event.Subscription, error) {

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "SetName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineSetName)
				if err := _DebtEngine.contract.UnpackLog(event, "SetName", log); err != nil {
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

// ParseSetName is a log parse operation binding the contract event 0x4df9dcd34ae35f40f2c756fd8ac83210ed0b76d065543ee73d868aec7c7fcf02.
//
// Solidity: event SetName(string _name)
func (_DebtEngine *DebtEngineFilterer) ParseSetName(log types.Log) (*DebtEngineSetName, error) {
	event := new(DebtEngineSetName)
	if err := _DebtEngine.contract.UnpackLog(event, "SetName", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineSetSymbolIterator is returned from FilterSetSymbol and is used to iterate over the raw logs and unpacked data for SetSymbol events raised by the DebtEngine contract.
type DebtEngineSetSymbolIterator struct {
	Event *DebtEngineSetSymbol // Event containing the contract specifics and raw log

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
func (it *DebtEngineSetSymbolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineSetSymbol)
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
		it.Event = new(DebtEngineSetSymbol)
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
func (it *DebtEngineSetSymbolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineSetSymbolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineSetSymbol represents a SetSymbol event raised by the DebtEngine contract.
type DebtEngineSetSymbol struct {
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSymbol is a free log retrieval operation binding the contract event 0xadf3ae8bd543b3007d464f15cb8ea1db3f44e84d41d203164f40b95e27558ac6.
//
// Solidity: event SetSymbol(string _symbol)
func (_DebtEngine *DebtEngineFilterer) FilterSetSymbol(opts *bind.FilterOpts) (*DebtEngineSetSymbolIterator, error) {

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "SetSymbol")
	if err != nil {
		return nil, err
	}
	return &DebtEngineSetSymbolIterator{contract: _DebtEngine.contract, event: "SetSymbol", logs: logs, sub: sub}, nil
}

// WatchSetSymbol is a free log subscription operation binding the contract event 0xadf3ae8bd543b3007d464f15cb8ea1db3f44e84d41d203164f40b95e27558ac6.
//
// Solidity: event SetSymbol(string _symbol)
func (_DebtEngine *DebtEngineFilterer) WatchSetSymbol(opts *bind.WatchOpts, sink chan<- *DebtEngineSetSymbol) (event.Subscription, error) {

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "SetSymbol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineSetSymbol)
				if err := _DebtEngine.contract.UnpackLog(event, "SetSymbol", log); err != nil {
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

// ParseSetSymbol is a log parse operation binding the contract event 0xadf3ae8bd543b3007d464f15cb8ea1db3f44e84d41d203164f40b95e27558ac6.
//
// Solidity: event SetSymbol(string _symbol)
func (_DebtEngine *DebtEngineFilterer) ParseSetSymbol(log types.Log) (*DebtEngineSetSymbol, error) {
	event := new(DebtEngineSetSymbol)
	if err := _DebtEngine.contract.UnpackLog(event, "SetSymbol", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineSetURIProviderIterator is returned from FilterSetURIProvider and is used to iterate over the raw logs and unpacked data for SetURIProvider events raised by the DebtEngine contract.
type DebtEngineSetURIProviderIterator struct {
	Event *DebtEngineSetURIProvider // Event containing the contract specifics and raw log

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
func (it *DebtEngineSetURIProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineSetURIProvider)
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
		it.Event = new(DebtEngineSetURIProvider)
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
func (it *DebtEngineSetURIProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineSetURIProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineSetURIProvider represents a SetURIProvider event raised by the DebtEngine contract.
type DebtEngineSetURIProvider struct {
	UriProvider common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetURIProvider is a free log retrieval operation binding the contract event 0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec.
//
// Solidity: event SetURIProvider(address _uriProvider)
func (_DebtEngine *DebtEngineFilterer) FilterSetURIProvider(opts *bind.FilterOpts) (*DebtEngineSetURIProviderIterator, error) {

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "SetURIProvider")
	if err != nil {
		return nil, err
	}
	return &DebtEngineSetURIProviderIterator{contract: _DebtEngine.contract, event: "SetURIProvider", logs: logs, sub: sub}, nil
}

// WatchSetURIProvider is a free log subscription operation binding the contract event 0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec.
//
// Solidity: event SetURIProvider(address _uriProvider)
func (_DebtEngine *DebtEngineFilterer) WatchSetURIProvider(opts *bind.WatchOpts, sink chan<- *DebtEngineSetURIProvider) (event.Subscription, error) {

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "SetURIProvider")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineSetURIProvider)
				if err := _DebtEngine.contract.UnpackLog(event, "SetURIProvider", log); err != nil {
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
func (_DebtEngine *DebtEngineFilterer) ParseSetURIProvider(log types.Log) (*DebtEngineSetURIProvider, error) {
	event := new(DebtEngineSetURIProvider)
	if err := _DebtEngine.contract.UnpackLog(event, "SetURIProvider", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DebtEngine contract.
type DebtEngineTransferIterator struct {
	Event *DebtEngineTransfer // Event containing the contract specifics and raw log

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
func (it *DebtEngineTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineTransfer)
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
		it.Event = new(DebtEngineTransfer)
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
func (it *DebtEngineTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineTransfer represents a Transfer event raised by the DebtEngine contract.
type DebtEngineTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_DebtEngine *DebtEngineFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*DebtEngineTransferIterator, error) {

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

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineTransferIterator{contract: _DebtEngine.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_DebtEngine *DebtEngineFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DebtEngineTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineTransfer)
				if err := _DebtEngine.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DebtEngine *DebtEngineFilterer) ParseTransfer(log types.Log) (*DebtEngineTransfer, error) {
	event := new(DebtEngineTransfer)
	if err := _DebtEngine.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebtEngineWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the DebtEngine contract.
type DebtEngineWithdrawnIterator struct {
	Event *DebtEngineWithdrawn // Event containing the contract specifics and raw log

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
func (it *DebtEngineWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtEngineWithdrawn)
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
		it.Event = new(DebtEngineWithdrawn)
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
func (it *DebtEngineWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtEngineWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtEngineWithdrawn represents a Withdrawn event raised by the DebtEngine contract.
type DebtEngineWithdrawn struct {
	Id     [32]byte
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xa6786aab7dbbc48b4b0387488b407bd81448030ab207b50bea7dbb5fbc1cd9eb.
//
// Solidity: event Withdrawn(bytes32 indexed _id, address _sender, address _to, uint256 _amount)
func (_DebtEngine *DebtEngineFilterer) FilterWithdrawn(opts *bind.FilterOpts, _id [][32]byte) (*DebtEngineWithdrawnIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.FilterLogs(opts, "Withdrawn", _idRule)
	if err != nil {
		return nil, err
	}
	return &DebtEngineWithdrawnIterator{contract: _DebtEngine.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xa6786aab7dbbc48b4b0387488b407bd81448030ab207b50bea7dbb5fbc1cd9eb.
//
// Solidity: event Withdrawn(bytes32 indexed _id, address _sender, address _to, uint256 _amount)
func (_DebtEngine *DebtEngineFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *DebtEngineWithdrawn, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _DebtEngine.contract.WatchLogs(opts, "Withdrawn", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtEngineWithdrawn)
				if err := _DebtEngine.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xa6786aab7dbbc48b4b0387488b407bd81448030ab207b50bea7dbb5fbc1cd9eb.
//
// Solidity: event Withdrawn(bytes32 indexed _id, address _sender, address _to, uint256 _amount)
func (_DebtEngine *DebtEngineFilterer) ParseWithdrawn(log types.Log) (*DebtEngineWithdrawn, error) {
	event := new(DebtEngineWithdrawn)
	if err := _DebtEngine.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
