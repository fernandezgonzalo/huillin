package logPuller

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type LogPuller struct {
	ethClient *ethclient.Client
	addresses []common.Address
}

func NewLogPuller(ethClient *ethclient.Client, contractAddresses []common.Address) LogPuller {
	papa := &LogPuller{}

	papa.ethClient = ethClient
	papa.addresses = contractAddresses

	return *papa
}

func (self LogPuller) PullLogs(fromBlock string, toBlock string) []types.Log {
	log.Info("Pulling logs from:" + fromBlock + " to:" + toBlock)
	from_block := new(big.Int)
	from_block, _ = from_block.SetString(fromBlock, 10)
	var query ethereum.FilterQuery

	if toBlock == "latest" {
		query = ethereum.FilterQuery{
			FromBlock: from_block,
			ToBlock:   nil,
			Addresses: self.addresses,
		}
	} else {
		to_block := new(big.Int)
		to_block, _ = to_block.SetString(toBlock, 10)
		query = ethereum.FilterQuery{
			FromBlock: from_block,
			ToBlock:   to_block,
			Addresses: self.addresses,
		}
	}

	logs, err := self.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	return logs
}
