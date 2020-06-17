package Installments

import (
	"fmt"
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
	"strings"
)

var Address = os.Getenv("INSTALLMENTS_ADDRESS")

const OwnershipTransferredHash string = "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"
const AddedDebtHash string = "0x2f39b1101bb100a4485571ad3147db16cd731b8fc83bc513c6b15acf4a24464f"
const AddedPaidHash string = "0xc778ddfdb0d4ccac7447183c180c8559ec30cdb51c3726b6c58b3bd342596093"
const ChangedDueTimeHash string = "0x8108c1693b12a3bca1add9e080d5917b7277a9c7a45f54c13c6cf211565229c9"
const ChangedFinalTimeHash string = "0x60e093518baf1919f536c54e83aada895ba93b5f32d096d8584e186ed3e317f1"
const ChangedFrecuencyHash string = "0x1ed1e38de3701289d91ec072a7b6a4e9ec4b4e290b8609afde370fef784737e0"
const ChangedObligationHash string = "0x35a463d00829418b627fa74ee01c7981a42d59eb6b813193ca2b7ecb28c88d6b"
const ChangedStatusHash string = "0x03202132f86ea915f0496d740ab57083235a3850003b398f13d3328f966e5e9e"
const CreatedHash string = "0x102d25c49d33fcdb8976a3f2744e0785c98d9e43b88364859e6aec4ae82eff5c"
const SetClockHash string = "0x5ac4222bff447e4964c68c4e68ba64f779bf5b79d7ba34c88fc1aa49f18eadbf"
const SetDescriptorHash string = "0xfd7a059fb02438ecc7334e2f3b4dbc55b37aeea735bb0391714041593a5beacd"
const SetEngineHash string = "0x0188a486877924f3e0702cef6a0566d0fc77155ab3edb363748f72c6c451ff43"
const SetInterestHash string = "0xfe35a3cf12dd1d691d662ec765e7ad19e6f21c24923847936c5ec4c0cc7186ba"
const SetPaidBaseHash string = "0xbe7c4be01d96a9021826187a698b198967f64a0d27f8a4e27228ae337c2f380b"

type OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner common.Address
}

type AddedDebt struct {
	ID string `json:_id`
	Amount *big.Int `json:_amount`
}

type AddedPaid struct {
	ID string `json:_id`
	Paid *big.Int `json:_paid`
}

type Created struct {
	ID string
}

type ChangedStatus struct {
	ID string
	Timestamp *big.Int `json:_timestamp`
	Status *big.Int `json:_status`
}

type ChangedObligation struct {
	ID string
	Timestamp *big.Int `json:_timestamp`
	Debt *big.Int `json:_debt`
}

type ChangedFrequency struct {
	ID string
	Timestamp *big.Int `json:_timestamp`
	Frecuency *big.Int `json:_frequency`
}

type ChangedDueTime struct {
	ID string
	Timestamp *big.Int `json:_timestamp`
	Status *big.Int `json:_status`
}

type ChangedFinalTime struct {
	ID        string
	Timestamp *big.Int `json:_timestamp`
	DueTime   *big.Int `json:_dueTime`
}

type SetEngine struct {
	Engine common.Address `json:_engine`
}

type SetDescriptor struct {
	Descriptor common.Address `json:_descriptor`
}

type SetClock struct {
	ID [32]byte // `json:_id`
	To uint64 //`json:_to`
}

type SetPaidBase struct {
	ID []byte `json:_id`
	PaidBase *big.Int `json:_paidBase`
}

type SetInterest struct {
	ID []byte `json:_id`
	Interest *big.Int `json:_interest`
}

type InstallmentsContract struct {
	ABI abi.ABI
	ethClient *ethclient.Client
	Address string
	DAL *db.MongoDAL
	Contract *Installments
	pusher notification.Pusher
}

func NewInstallmentsContract(ethClient *ethclient.Client, dal *db.MongoDAL) InstallmentsContract {
	i := &InstallmentsContract{}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	installmentsABI, err := abi.JSON(strings.NewReader(InstallmentsABI))
	if err != nil {
		log.Fatal(err)
	}

	contract, err := NewInstallments(common.HexToAddress(Address), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	i.ABI = installmentsABI
	i.ethClient = ethClient
	i.Address = Address
	i.DAL = dal
	i.Contract = contract
	i.pusher = notification.NewPusher()

	return *i
}

func (self InstallmentsContract) HandleEvent(event types.Log, block types.Block) {
	switch event.Topics[0].Hex() {
	case OwnershipTransferredHash:
		log.Info("Log OwnershipTransferredHash")
	case AddedDebtHash:
		log.Info("Log AddedDebtHash")
	case AddedPaidHash:
		id := event.Topics[1].Hex()
		log.Info("Log AddedPaidHash: ", id)

		installment, err := self.DAL.GetInstallment(id)
		if err != nil {
			log.Fatal(err)
		}

		var addedPaid AddedPaid
		err = self.ABI.Unpack(&addedPaid, "AddedPaid", event.Data)
		if err != nil {
			log.Fatal(err)
		}
		var newPaid big.Int
		oldPaid, _ := new(big.Int).SetString(installment.State.Paid, 10)
		newPaid.Add(oldPaid, addedPaid.Paid)

		installment.State.Paid = newPaid.String()
		installment.State.LastPayment = installment.State.Clock

		err = self.DAL.UpdateInstallment(installment)
		if err != nil {
			log.Fatal(err)
		}
	case ChangedDueTimeHash:
		log.Info("Log ChangedDueTimeHash")
	case ChangedFinalTimeHash:
		log.Info("Log ChangedFinalTimeHash")
	case ChangedFrecuencyHash:
		log.Info("Log ChangedFrecuencyHash")
	case ChangedObligationHash:
		log.Info("Log ChangedObligationHash")
	case ChangedStatusHash:
		id := event.Topics[1].Hex()
		log.Info("Log ChangedStatusHash: ", id)

		installment, err := self.DAL.GetInstallment(id)
		if err != nil {
			log.Fatal(err)
		}

		var changedStatus ChangedStatus
		err = self.ABI.Unpack(&changedStatus, "ChangedStatus", event.Data)
		if err !=  nil {
			log.Fatal(err)
		}
		installment.State.Status = changedStatus.Status.String()

		err = self.DAL.UpdateInstallment(installment)

		loan, err := self.DAL.GetLoan(id)
		if err != nil {
			log.Fatal(err)
		}
		loan.Status = changedStatus.Status.String()

		err = self.DAL.UpdateLoan(loan)
		if err != nil {
			log.Fatal(err)
		}

		collaterals, _ := self.DAL.GetCollaterals(loan.ID)
		for _, collateral := range collaterals {
			if collateral.Status == "2" || collateral.Status == "3" {
				collateral.Status = "4"
				self.DAL.UpdateCollateral(&collateral)
			}
		}
		self.pusher.Push("installments:changed_status", id)
	case CreatedHash:
		id := event.Topics[1]
		log.Info("Log CreatedHash: ", id.Hex())

		c, err := self.Contract.Configs(&bind.CallOpts{}, id)
		if err != nil {
			log.Fatal(err)
		}

		configDB := db.Config{}
		configDB.Installments = c.Installments.String()
		configDB.Cuota = c.Cuota.String()
		configDB.Duration = c.Duration.String()
		configDB.InterestRate = c.InterestRate.String()

		configDB.LentTime = fmt.Sprintf("%v", c.LentTime)
		configDB.TimeUnit = fmt.Sprintf("%v", c.TimeUnit)

		stateDB := db.State{}
		stateDB.Paid = "0"
		stateDB.Status = "0"
		stateDB.Clock = "0"
		stateDB.Interest = "0"
		stateDB.LastPayment = "0"
		stateDB.PaidBase = "0"

		installment := db.DBInstallment{}
		installment.ID = id.Hex()
		installment.State = &stateDB
		installment.Config = &configDB

		err = self.DAL.SaveInstallment(&installment)
		if err != nil {
			log.Fatal(err)
		}
	case SetClockHash:
		id := common.BytesToHash(event.Data[:32]).Hex()
		log.Info("Log SetClockHash: ", id)

		to := common.BytesToHash(event.Data[32:]).Big()

		installment, err := self.DAL.GetInstallment(id)
		if err != nil {
			log.Fatal(err)
		}
		installment.State.Clock = to.String()

		err = self.DAL.UpdateInstallment(installment)
		if err != nil {
			log.Fatal(err)
		}
	case SetDescriptorHash:
		log.Info("Log SetDescriptorHash")
	case SetEngineHash:
		log.Info("Log SetEngineHash")
	case SetInterestHash:
		id := common.BytesToHash(event.Data[:32]).Hex()
		log.Info("Log SetInterestHash: ", id)

		interest := common.BytesToHash(event.Data[32:]).Big()

		installment, err := self.DAL.GetInstallment(id)
		if err != nil {
			log.Fatal(err)
		}
		installment.State.Interest = interest.String()

		err = self.DAL.UpdateInstallment(installment)
		if err != nil {
			log.Fatal(err)
		}

	case SetPaidBaseHash:
		id := common.BytesToHash(event.Data[:32]).Hex()
		log.Info("Log SetPaidBaseHash: ", id)

		paidBase := common.BytesToHash(event.Data[32:]).Big()

		installment, err := self.DAL.GetInstallment(id)
		if err != nil {
			log.Fatal(err)
		}

		installment.State.PaidBase = paidBase.String()

		err = self.DAL.UpdateInstallment(installment)
		if err != nil {
			log.Fatal(err)
		}
	}
}