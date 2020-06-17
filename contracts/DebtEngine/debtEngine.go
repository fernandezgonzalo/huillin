package DebtEngine

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"sarlanga/db"
	"sarlanga/notification"
	"strconv"
	"strings"
	"time"
)

var Address = os.Getenv("DEBT_ENGINE_ADDRESS")

const OwnershipTransferredHash string = "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"
const SetURIProviderHash string = "0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec"
const TransferHash string = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
const ApprovalHash string = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
const ApprovalForAllHash string = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
const CreatedHash string = "0xc9310bea14cefb8c6d9a8fce4ead7f6825a31223af9d0bde2a30b2d0c71ccbf3"
const Created2Hash string = "0x4ee48d90c05f58fc51a05728e403947a51e36cc7a1bf4d82260473b433538408"
const Created3Hash string = "0xc46d3d245e649272659dd5731879c1b019caccb00acb3a94488a28c82ef924a6"
const PaidHash string = "0x25b52320bc27b845f37eae2240cc285c7b6e5643fc2995e6d22afa10e2f657d2"
const ReadedOracleBatchHash string = "0x79601cc1ca8f1ff7c3e7f7b522c2f1377cdb4c318a131afaf426cb2c976ef4be"
const ReadedOracleHash string = "0x686ef5859f39efd8bf5a57659ca5b60fe18f984d435a0b7418f18bcbe2d349d6"
const PayBatchErrorHash string = "0x8a1ca4383d0682548257effd0249245baaabe2c4672c83c48397a88faa0831f0"
const WithdrawnHash string = "0xa6786aab7dbbc48b4b0387488b407bd81448030ab207b50bea7dbb5fbc1cd9eb"
const ErrorHash string = "0xe000ee8a04c9b1ff738cb5c90d6a362bf5be3c83b1eb221377df6324747d4c07"
const ErrorRecoverHash string = "0xa43b78c6d135806f3e817a90a7b690c0d2d94b7d2e356049fa5ba357c0fab9c3"

type OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner common.Address
}

type SetURIProvider struct {
	UriProvider common.Address `json:_uriProvider`
}

type Transfer struct {
	ID string
	From common.Address
	To common.Address
}

type Approval struct {
	ID string
	Owner common.Address
	Approved common.Address
}

type ApprovalForAll struct {
	Owner common.Address
	Operator common.Address
	Approved bool
}

type Created struct {
	ID string
	Nonce *big.Int `json:_nonce`
	Data []byte `json:_data`
}

type Created2 struct {
	ID string
	Salt *big.Int `json:_salt`
	Data []byte `json:_data`
}

type Created3 struct {
	ID string
	Salt *big.Int `json:_salt`
	Data []byte `json:_data`
}

type Paid struct {
	ID string
	Sender common.Address `json:_sender`
	Origin common.Address `json:_origin`
	Requested *big.Int `json:_requested`
	RequestedTokens *big.Int `json:_requestedTokens`
	Paid *big.Int `json:_paid`
	Tokens *big.Int `json:_tokens`
}

type ReadedOracleBatch struct {
	Oracle common.Address `json:_oracle`
	Count *big.Int `json:_count`
	Tokens *big.Int `json:_tokens`
	Equivalent *big.Int `json:_equivalent`
}

type ReadedOracle struct {
	ID string
	Tokens *big.Int `json:_tokens`
	Equivalent *big.Int `json:_equivalent`
}
type PayBatchError struct {
	ID string
	TargetOracle common.Address `json:_targetOracle`
}

type Withdrawn struct {
	ID string
	Sender common.Address `json:_sender`
	To common.Address `json:_to`
	Amount *big.Int `json:_amount`
}

type Error struct {
	ID string
	Sender common.Address `json:_sender`
	Value *big.Int `json:_value`
	GasLeft *big.Int `json:_gasLeft`
	GasLimit *big.Int `json:_gasLimit`
	CallData []byte `json:_callData`
}

type ErrorRecover struct {
	ID string
	Sender common.Address `json:_sender`
	Value *big.Int `json:_value`
	GasLeft *big.Int `json:_gasLeft`
	GasLimit *big.Int `json:_gasLimit`
	Result []byte `json:_result`
	CallData []byte `json:_callData`
}

type DebtEngineContract struct {
	ABI abi.ABI
	ethClient *ethclient.Client
	Address string
	DAL *db.MongoDAL
	Contract *DebtEngine
	pusher notification.Pusher
}

func NewDebtEngineContract(ethClient *ethclient.Client, dal *db.MongoDAL) DebtEngineContract {
	db := &DebtEngineContract{}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	debtEngineABI, err := abi.JSON(strings.NewReader(DebtEngineABI))
	if err != nil {
		log.Fatal(err)
	}

	contract, err := NewDebtEngine(common.HexToAddress(Address), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	db.ABI = debtEngineABI
	db.ethClient = ethClient
	db.Address = Address
	db.DAL = dal
	db.Contract = contract
	db.pusher = notification.NewPusher()

	return *db
}

func (self DebtEngineContract) HandleEvent(event types.Log, block types.Block) {
	switch event.Topics[0].Hex() {
	case OwnershipTransferredHash:
		log.Info("Log OwnershipTransferredHash")
	case SetURIProviderHash:
		log.Info("Log SetURIProviderHash")
	case TransferHash:
		from := event.Topics[1].Hex()
		to := event.Topics[2].Hex()
		id := event.Topics[3].Hex()
		log.Info("Log TransferHash: ", id)

		if from != "0x0000000000000000000000000000000000000000000000000000000000000000" {
			debt, err := self.DAL.GetDebt(id)
			if err == nil {
				debt.Owner = to
				err = self.DAL.UpdateDebt(debt)
				if err != nil {
					log.Fatal(err)
				}
			}

		} else {
			id := event.Topics[3]
			d, err := self.Contract.Debts(&bind.CallOpts{}, id)
			if err != nil {
				log.Fatal(err)
			}

			debtDB := db.DBDebt{}
			debtDB.Creator = d.Creator.String()
			debtDB.Oracle = d.Oracle.String()
			debtDB.Model = d.Model.String()
			debtDB.Error = false
			debtDB.ID = id.Hex()
			debtDB.Created = strconv.FormatInt(time.Now().Unix(), 10)
			debtDB.Balance = "0"

			err = self.DAL.SaveDebt(&debtDB)
			if err != nil {
				log.Fatal(err)
			}
		}
	case ApprovalHash:
		log.Info("Log ApprovalHash")
	case ApprovalForAllHash:
		log.Info("Log ApprovalForAllHash")
	case CreatedHash:
		log.Info("Log CreatedHash")
	case Created2Hash:
		log.Info("Log Created2Hash")
	case Created3Hash:
		log.Info("Log Created3Hash")
	case PaidHash:
		id := event.Topics[1].Hex()
		log.Info("Log PaidHash: ", id)

		var paid Paid
		err := self.ABI.Unpack(&paid, "Paid", event.Data)
		if err != nil {
			log.Fatal(err)
		}

		debt, err := self.DAL.GetDebt(id)
		if err == nil {
			var newBalance big.Int
			oldBalance := new(big.Int)

			oldBalance, _ = oldBalance.SetString(debt.Balance, 10)
			newBalance.Add(oldBalance, paid.Tokens)

			debt.Balance = newBalance.String()

			err = self.DAL.UpdateDebt(debt)
			if err != nil {
				log.Fatal(err)
			}
			self.pusher.Push("debt_engine:paid", debt.ID)
		}


	case ReadedOracleBatchHash:
		log.Info("Log ReadedOracleBatchHash")
	case ReadedOracleHash:
		log.Info("Log ReadedOracleHash")
	case PayBatchErrorHash:
		log.Info("Log PayBatchErrorHash")
	case WithdrawnHash:
		id := event.Topics[1].Hex()
		log.Info("Log WithdrawnHash: ", id)

		var withdrawn Withdrawn
		err := self.ABI.Unpack(&withdrawn, "Withdrawn", event.Data)
		if err != nil {
			log.Fatal(err)
		}

		debt, err := self.DAL.GetDebt(id)
		if err == nil {
			var newBalance big.Int
			oldBalance := new(big.Int)

			oldBalance, ok := oldBalance.SetString(debt.Balance, 10)
			if !ok {
				log.Fatal(err)
			}
			newBalance.Sub(oldBalance, withdrawn.Amount)

			debt.Balance = newBalance.String()

			err = self.DAL.UpdateDebt(debt)
			if err != nil {
				log.Fatal(err)
			}
		}
	case ErrorHash:
		id := event.Topics[1].Hex()
		log.Info("Log ErrorHash: ", id)
		debt, err := self.DAL.GetDebt(id)
		if err == nil {
			debt.Error = true
			err = self.DAL.UpdateDebt(debt)
			if err != nil {
				log.Fatal(err)
			}
		}
	case ErrorRecoverHash:
		id := event.Topics[1].Hex()
		log.Info("Log ErrorRecoverHash: ", id)
		debt, err := self.DAL.GetDebt(id)
		if err != nil {
			log.Fatal(err)
		}

		debt.Error = false

		err = self.DAL.UpdateDebt(debt)
		if err != nil {
			log.Fatal(err)
		}
	}
}