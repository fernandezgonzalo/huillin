package LoanManager

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"sarlanga/contracts/Installments"
	"sarlanga/contracts/Oracle"
	"sarlanga/db"
	"sarlanga/notification"
	"strings"
)

var Address = os.Getenv("LOAN_MANAGER_ADDRESS")

const ApprovedHash string = "0x6fb7fd1eda743aa3eb32c69f3b8cf14a5aeadf26db51057a7c5c78ba10eac8a4"
const ApprovedByCallBackHash string = "0xbcd1a053efeeac2159254030853620c82ddeffb174f76982dc0dd1083648ade3"
const ApprovedBySignatureHash string = "0xac4faa318b3929891fe5ee57344cce181e72bced7246f5aa4a788ab61e7bb5d2"
const ApprovedErrorHash string = "0x3e35d8182867e897023d9da65f18a2d16b1874c35cf8dd2333276ea1ee5f54e9"
const ApprovedRejectedHash string = "0x7b550df26fb29cebc893b0071262bded4c36eb15da43317bf72c3f9f5aeb554a"
const BorrowerByCallbackHash string = "0x168544dde877630d31a353bc62115efd38d8e072b48761461075d8421620e772"
const BorrowerBySignatureHash string = "0x04acecab035917c895d49033509a027f3fbb1913e978eec824926f605370653b"
const CanceledHash string = "0x5f32c817b1de98e681044c3d78ca791e6018e4c3b4f230e4d85bbcc3815989f3"
const CosignedHash string = "0x8a4448add099432884d9833a6b44c465104441b7659c2c00b037da123a54f9d2"
const CreatorByCallbackHash string = "0x025f21426e185fab3bf170a9374c750b17a510c3ed51549fa4ff068dbca66057"
const CreatorBySignatureHash string = "0x3971d1c591487c96ba2561ab5f74f00001a6134408897ab8b0124c1673055dd2"
const LentHash string = "0xf701a0be7203a591a078aa39b506b74dcf910d7a73b1f73a4dff4d0df3968361"
const ReadedOracleHash string = "0x0b53fa0f90e4773f861d84b5da681ae7b3dd6b29358fdb4a38ba8dec13fc80f3"
const RequestedHash string = "0x200404915876dec9fd44568ed7f5358a7fc4b212c1cdfc404ab17be21a1cf4af"
const SettledCancelHash string = "0x029e3e789ef1a202e83e668d22e38205106d93105847601f49983c21efb8cfbe"
const SettledLendHash string = "0xb1d906969e9527a5ad3ac2ca84eb415890ddf44581528edf18d0e7154d7f126d"

type Approved struct {
	id string
}

type Lent struct {
	id string
	Lender common.Address `json:_lender`
	Tokens *big.Int `json:_tokens`
}

type Requested struct {
	id string
	Amount *big.Int `json:_amount`
	Model common.Address `json:_model`
	Creator common.Address `json:_creator`
	Oracle common.Address `json:_oracle`
	Borrower common.Address `json:_borrower`
	Callback common.Address `json:_callback`
	Salt *big.Int `json:_salt`
	LoanData []byte `json:_loanData`
	Expiration *big.Int `json:_expiration`
}

type Cosigned struct {
	id string
	Cosigner common.Address `json:_cosigner`
	Cost *big.Int `json:_cost`
}

type Canceled struct {
	id string
	Canceler common.Address `json:_canceler`
}

type ReadedOracle struct {
	id string
	Oracle common.Address `json:_oracle`
}

type ApprovedRejected struct {
	id string
	Response []byte `json:_response`
}
type ApprovedError struct {
	id string
	Response []byte `json:_response`
}

type ApprovedByCallback struct {
	id string
}
type ApprovedBySignature struct {
	id string
}
type CreatorByCallback struct {
	id string
}
type BorrowerByCallback struct {
	id string
}
type CreatorBySignature struct {
	id string
}
type BorrowerBySignature struct {
	id string
}
type SettledLend struct {
	id string
	Lender common.Address `json:_lender`
	Tokens *big.Int `json:_tokens`
}
type SettledCancel struct {
	id string
	Canceler common.Address `json:_canceler`
}

type LoanManagerContract struct {
	ABI abi.ABI
	ethClient *ethclient.Client
	Address string
	DAL *db.MongoDAL
	Contract *LoanManager
	pusher notification.Pusher
}

func NewLoanManagerContract(ethClient *ethclient.Client, dal *db.MongoDAL) LoanManagerContract {
	lm := &LoanManagerContract{}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	loanManagerABI, err := abi.JSON(strings.NewReader(LoanManagerABI))
	if err != nil {
		log.Fatal(err)
	}

	contract, err := NewLoanManager(common.HexToAddress(Address), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	lm.ABI = loanManagerABI
	lm.ethClient = ethClient
	lm.Address = Address
	lm.DAL = dal
	lm.Contract = contract
	lm.pusher = notification.NewPusher()

	return *lm
}

func (self LoanManagerContract) HandleEvent(event types.Log, block types.Block) {
	switch event.Topics[0].Hex() {
	case ApprovedHash:
		id := event.Topics[1].Hex()
		log.Info("Log Approved: ", id)
		loan, err := self.DAL.GetLoan(id)
		if err != nil {
			log.Fatal(err)
		}
		loan.Approved = true

		err = self.DAL.UpdateLoan(loan)
		if err != nil {
			log.Fatal(err)
		}

	case ApprovedByCallBackHash:
		log.Info("Log ApprovedByCallback")
	case ApprovedBySignatureHash:
		log.Info("Log ApprovedBySignature")
	case ApprovedErrorHash:
		log.Info("Log ApprovedError")
	case ApprovedRejectedHash:
		log.Info("Log ApprovedRejected")
	case BorrowerByCallbackHash:
		log.Info("Log BorrowerByCallback")
	case BorrowerBySignatureHash:
		log.Info("Log BorrowerBySignature")
	case CanceledHash:
		id := event.Topics[1].Hex()
		log.Info("Log Canceled: ", id)

		loan, err := self.DAL.GetLoan(id)
		if err != nil {
			log.Fatal(err)
		}
		loan.Canceled = true

		err = self.DAL.UpdateLoan(loan)
		if err != nil {
			log.Fatal(err)
		}
	case CosignedHash:
		log.Info("Log Cosigned")
	case CreatorByCallbackHash:
		log.Info("Log CreatorByCallback")
	case CreatorBySignatureHash:
		log.Info("Log CreatorBySignature")
	case LentHash:
		id := event.Topics[1].Hex()
		log.Info("Log Lent: ", id)
		var lent Lent

		err := self.ABI.Unpack(&lent, "Lent", event.Data)
		if err != nil {
			log.Fatal(err)
		}

		lent.id = id
		loan, err := self.DAL.GetLoan(lent.id)
		if err != nil {
			log.Fatal(loan)
		}

		loan.Open = false
		loan.Status = "1"

		err = self.DAL.UpdateLoan(loan)
		if err != nil {
			log.Fatal(err)
		}

		collaterals, _ := self.DAL.GetCollaterals(loan.ID)

		for _, collateral := range collaterals {
			collateral.Status = "4"
			self.DAL.SaveCollateral(&collateral)
		}

		self.pusher.Push("loan_manager:lent", lent.id)
	case ReadedOracleHash:
		log.Info("Log ReadedOracle")
	case RequestedHash:
		id := event.Topics[1].Hex()
		log.Info("Log Requested: ", id)
		var requested Requested
		err := self.ABI.Unpack(&requested, "Requested", event.Data)
		if err != nil {
			log.Fatal(err)
		}
		requested.id = id
		loan := db.DBLoan{}

		// unpack from loanManagerContract

		parsedRequested, _ := self.Contract.ParseRequested(event)

		var currency [32]byte
		if parsedRequested.Oracle.Hex() != "0x0000000000000000000000000000000000000000" {
			oracleContract, _ := Oracle.NewOracle(parsedRequested.Oracle, self.ethClient)
			currency, _ = oracleContract.Currency(nil)
		} else {
			currency, err = self.Contract.GetCurrency(nil, event.Topics[1].Big())
		}
		loan.Currency = common.BytesToHash(currency[:]).Hex()

		loan.ID = requested.id
		loan.Open = true
		loan.Approved = requested.Creator == requested.Borrower
		loan.Expiration = requested.Expiration.String()
		loan.Amount = requested.Amount.String()
		loan.Cosigner = "0x0000000000000000000000000000000000000000"
		loan.Model = requested.Model.String()
		loan.Creator = requested.Creator.String()
		loan.Oracle = requested.Oracle.String()
		loan.Borrower = requested.Borrower.String()
		loan.LoanData = common.BytesToHash(parsedRequested.LoanData).Hex()
		loan.Created = requested.Creator.String()
		loan.Status = "0"
		loan.Canceled = false


		err = self.DAL.SaveLoan(&loan)
		if err != nil {
			log.Fatal(err)
		}
		go SaveDescriptor(*self.ethClient, *parsedRequested, self.DAL)

		self.pusher.Push("loan_manager:requested", requested.id)
	case SettledCancelHash:
		log.Info("Log SettledCancel")
	case SettledLendHash:
		log.Info("Log SettledLend")

	}
}

func SaveDescriptor(ethClient ethclient.Client, parsedRequest LoanManagerRequested, DAL *db.MongoDAL)  {
	ModelContract, _ := Installments.NewInstallments(parsedRequest.Model, &ethClient)

	validOracle, _ := ModelContract.Validate(nil, parsedRequest.LoanData)
	d := &db.DBDescriptor{}
	if validOracle {

		installments, _ := ModelContract.SimInstallments(nil, parsedRequest.LoanData)
		firstObligation, _ := ModelContract.SimFirstObligation(nil, parsedRequest.LoanData)
		totalObligation, _ := ModelContract.SimTotalObligation(nil, parsedRequest.LoanData)
		duration, _ := ModelContract.SimDuration(nil, parsedRequest.LoanData)
		punitiveInterestRate, _ := ModelContract.SimPunitiveInterestRate(nil, parsedRequest.LoanData)
		frecuency, _ := ModelContract.SimFrequency(nil, parsedRequest.LoanData)

		d.ID = common.BytesToHash(parsedRequest.Id[:]).Hex()
		d.Installments = installments.String()
		d.Duration = duration.String()
		d.FirstObligation = firstObligation.Amount.String()
		d.Frequency = frecuency.String()
		d.PunitiveInterestRate = punitiveInterestRate.String()
		d.TotalObligation = totalObligation.String()

		d.InterestRate = "0"
	}
	DAL.SaveDescriptor(d)
}
