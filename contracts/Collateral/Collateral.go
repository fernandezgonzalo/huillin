package Collateral

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
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

var Address = os.Getenv("COLLATERAL_ADDRESS")

const ApprovalForAllHash string = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
const ApprovalHash string = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
const BorrowCollateralHash string = "0x9b33a4c997038b93eec9f3b361e2d2ff5b830ce8a15d2adbb472e1a98f3d2fdf"
const ClaimedExpiredHash string = "0xb8d56805081871105665f00274333d930015d8cff88cf02a2d8e4b99227ddb4e"
const ClaimedLiquidationHash string = "0x4df462f6bb661fc317efa25950d47fbd7f4ad62c6b69664c01bab0f119f02d32"
const ClosedAuctionHash string = "0xeff56510b465b07772acd0d1a3723fe4a7cce582a26efe952f226872fbc3595c"
const CreatedHash string = "0xfe5bb50bc3d4e5eea6293acf0cb8bb5ff31625bb76a447fa888dd393b980b52a"
const DepositedHash string = "0x06da3309189fa49284f335d2c2bcb4cb0b8ad2a59ad92a9bdebeeb8f1ceba511"
const OwnershipTransferredHash string = "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"
const RedeemHash string = "0xc5fcfb68332ef11d542d4ab7e75045a5e4d66eb2d8f846f13a48356b32e534db"
const SetURIProviderHash string = "0x8830bfff0a198778822a37d97bfba3d9d6e08bcd080eb82f2a76f2060a7494ec"
const SetURLHash string = "0x93dbd152ec966d01e4950135dfb03d66ff1a9d38d1dbf031f2566162532d06db"
const Started string = "0x006e0c97de781a7389d44ba8fd35d1467cabb17ed04d038d166d34ab819213f3"
const TransferHash string = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
const WithdrawHash string = "0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf"

type CollateralContract struct {
	ABI abi.ABI
	ethClient *ethclient.Client
	Address string
	DAL *db.MongoDAL
	Contract *Collateral
	pusher notification.Pusher
}

func NewCollateralContract (ethClient *ethclient.Client, dal *db.MongoDAL) CollateralContract {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	cc := &CollateralContract{}

	collateralABI, err := abi.JSON(strings.NewReader(CollateralABI))
	if err != nil {
		log.Fatal(err)
	}

	contract, err := NewCollateral(common.HexToAddress(Address), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	cc.ABI = collateralABI
	cc.ethClient = ethClient
	cc.Address = Address
	cc.DAL = dal
	cc.Contract = contract
	cc.pusher = notification.NewPusher()

	return *cc
}

func (self CollateralContract) HandleEvent(event types.Log, block types.Block) {
	topic0 := event.Topics[0].Hex()
	switch topic0 {
	case ApprovalForAllHash:
		log.Info("ApprovalForAll")
	case ApprovalHash:
		log.Info("Approval")
	case BorrowCollateralHash:
		parsedEvent, _ := self.Contract.ParseBorrowCollateral(event)
		log.Info("BorrowCollateral: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())
		collateral.Status = "4"
		collateral.Amount = parsedEvent.NewAmount.String()

		err := self.DAL.UpdateCollateral(collateral)
		if err != nil {
			log.Fatal(err)
		}

	case ClaimedExpiredHash:
		parsedEvent, _ := self.Contract.ParseClaimedExpired(event)
		log.Info("ClaimedExpirad: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())
		collateral.Status = "3"

		err := self.DAL.UpdateCollateral(collateral)

		if err != nil {
			log.Fatal(err)
		}

	case ClaimedLiquidationHash:
		parsedEvent, _ := self.Contract.ParseClaimedLiquidation(event)
		log.Info("ClaimedLiquidation: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())
		collateral.Status = "3"

		err := self.DAL.UpdateCollateral(collateral)

		if err != nil {
			log.Fatal(err)
		}

	case ClosedAuctionHash:
		parsedEvent, _ := self.Contract.ParseClosedAuction(event)
		log.Info("ClosedAuction: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())
		loan, _ := self.DAL.GetLoan(collateral.DebtID)

		if loan.Status == "2" {
			collateral.Status = "4"
		} else {
			collateral.Status = "2"
		}
		collateral.Amount = parsedEvent.Leftover.String()

		err := self.DAL.UpdateCollateral(collateral)

		if err != nil {
			log.Fatal(err)
		}

	case CreatedHash:
		parsedEvent, _ := self.Contract.ParseCreated(event)
		log.Info("Created: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())

		collateral.DebtID = common.BytesToHash(parsedEvent.DebtId[:]).Hex()
		collateral.Oracle = parsedEvent.Oracle.Hex()
		collateral.Token = parsedEvent.Token.Hex()
		collateral.LiquidationRatio = parsedEvent.LiquidationRatio.String()
		collateral.BalanceRatio = parsedEvent.BalanceRatio.String()
		collateral.Amount = parsedEvent.Amount.String()
		collateral.Status = "1"

		err := self.DAL.UpdateCollateral(collateral)

		if err != nil {
			log.Fatal(err)
		}

	case DepositedHash:
		parsedEvent, _ := self.Contract.ParseDeposited(event)
		log.Info("Deposited: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())
		if collateral.Status == "5" && parsedEvent.Amount.String() != "0" {
			collateral.Status = "4"
		}

		var newAmount big.Int
		oldAmount, _ := new(big.Int).SetString(collateral.Amount, 10)
		newAmount.Add(oldAmount, parsedEvent.Amount)

		collateral.Amount = newAmount.String()

		err := self.DAL.UpdateCollateral(collateral)
		if err != nil {
			log.Fatal(err)
		}


	case OwnershipTransferredHash:
		log.Info("OwnershipTransferred")
	case RedeemHash:
		log.Info("Redeem")
	case SetURIProviderHash:
		log.Info("SetURIProvider")
	case SetURLHash:
		log.Info("SetURL")
	case Started:
		parsedEvent, _ := self.Contract.ParseStarted(event)
		log.Info("Started: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())
		collateral.Status = "2"

		err := self.DAL.UpdateCollateral(collateral)
		if err != nil {
			log.Fatal(err)
		}

		collaterals, _ := self.DAL.GetCollaterals(collateral.DebtID)
		for _, collateral_ := range collaterals {
			if collateral_.ID != collateral.ID {
				collateral_.Status = "4"
				self.DAL.UpdateCollateral(&collateral_)
			}
		}

	case TransferHash:
		parsedEvent, _ := self.Contract.ParseTransfer(event)
		log.Info("Transfer: " + parsedEvent.TokenId.String())

		collateral := db.DBCollateral{}
		collateral.ID = parsedEvent.TokenId.String()
		collateral.Owner = parsedEvent.To.Hex()

		err := self.DAL.SaveCollateral(&collateral)
		if err != nil {
			log.Fatal(err)
		}


	case WithdrawHash:
		parsedEvent, _ := self.Contract.ParseWithdraw(event)
		log.Info("Withdraw: " + parsedEvent.EntryId.String())

		collateral, _ := self.DAL.GetCollateral(parsedEvent.EntryId.String())

		var newAmount big.Int
		oldAmount, _ := new(big.Int).SetString(collateral.Amount, 10)
		newAmount.Sub(oldAmount, parsedEvent.Amount)
		collateral.Amount = newAmount.String()

		if newAmount.String() == "0" {
			collateral.Status = "5"
		}

		err := self.DAL.UpdateCollateral(collateral)
		if err != nil {
			log.Fatal(err)
		}

	}
}

