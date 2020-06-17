package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"sarlanga/contracts/Collateral"
	"sarlanga/contracts/DebtEngine"
	"sarlanga/contracts/Installments"
	"sarlanga/contracts/LoanManager"
	"sarlanga/db"
	"sarlanga/logPuller"
	"sarlanga/utils"
	"time"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v7"
)

func MinBlock(block1 string, block2 string) string {
	if block1 > block2 {
		return block2
	} else if block1 < block2 {
		return block1
	} else {
		return block1
	}
}

func NextNumber(number big.Int) *big.Int {
	return big.NewInt(0).Add(&number, big.NewInt(1))
}

func StepNumber(number big.Int, step big.Int) *big.Int {
	return big.NewInt(0).Add(&number, &step)
}

func StrToBigInt(number string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(number, 10)
	return n
}

func GetLastBlock(ethClient *ethclient.Client) *types.Header {
	header, _ := ethClient.HeaderByNumber(context.Background(), nil)
	return header
}

func watcher() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client, err := ethclient.Dial(os.Getenv("NODE_URL_WS"))  // "wss://ropsten.infura.io/ws/v3/fb5d21a911c1487ba9835be28b370b12"
	if err != nil {
		log.Fatal(err)
	}

	lastBlock := GetLastBlock(client)
	redisClient.Set("lastBlock", lastBlock.Number.String(), 0)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			log.Info("new block: " + header.Number.String())
			err := redisClient.Set("lastBlock", header.Number.String(), 0).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	go watcher()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	initialNumber := StrToBigInt(os.Getenv("INITIAL_NUMBER"))

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")))
	MongoClient, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}


	dal := db.NewMongoDAL(MongoClient)

	ethClient, err := ethclient.Dial(os.Getenv("NODE_URL_HTTP"))  // "https://ropsten.infura.io/v3/fb5d21a911c1487ba9835be28b370b12"
	if err != nil {
		log.Fatal(err)
	}

	loanManager := LoanManager.NewLoanManagerContract(ethClient, &dal)
	debtEngine := DebtEngine.NewDebtEngineContract(ethClient, &dal)
	installments := Installments.NewInstallmentsContract(ethClient, &dal)
	collateral := Collateral.NewCollateralContract(ethClient, &dal)

	Block, err := ethClient.BlockByNumber(context.Background(), initialNumber)
	if err != nil {
		log.Fatal(err)
	}

	currentBlock := initialNumber
	addresses := []common.Address{
		common.HexToAddress(LoanManager.Address),
		common.HexToAddress(DebtEngine.Address),
		common.HexToAddress(Installments.Address),
		common.HexToAddress(collateral.Address),
	}

	papa := logPuller.NewLogPuller(ethClient, addresses)
	steps := StrToBigInt(os.Getenv("STEPS_NUMBER_BLOCKS"))  //big.NewInt(int64(1000000))

	time.Sleep(1 * time.Second)
	for {
		lastBlock, _ := redisClient.Get("lastBlock").Result()

		if lastBlock == currentBlock.String() {
			log.Info("Synced sleep 1 sec")
			time.Sleep(1 * time.Second)
		} else {
			toBlock := MinBlock(lastBlock, StepNumber(*currentBlock, *steps).String())
			logs := papa.PullLogs(NextNumber(*currentBlock).String(), toBlock)

			for _, vLog := range logs {
				if vLog.BlockNumber != Block.NumberU64() {
					Block, err = utils.GetBlock(big.NewInt(int64(vLog.BlockNumber)), *ethClient, 5, 1)
					log.Info("Pull blockNumber:", Block.Number())
				}
				switch vLog.Address.Hex() {
				case loanManager.Address:
					loanManager.HandleEvent(vLog, *Block)
				case debtEngine.Address:
					debtEngine.HandleEvent(vLog, *Block)
				case installments.Address:
					installments.HandleEvent(vLog, *Block)
				case collateral.Address:
					collateral.HandleEvent(vLog, *Block)
				}
				// save block on redis
				redisClient.Set("lastBlockProcesed", Block.Number().String(), 0)
			}
			redisClient.Set("lastBlockProcesed", toBlock, 0)

			currentBlock = StrToBigInt(toBlock)
			log.Info("CurrentBlock", currentBlock)
		}
	}
}
