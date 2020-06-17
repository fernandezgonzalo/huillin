package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"math/big"
	"time"
)

func GetBlock(number *big.Int, client ethclient.Client, retries int, sleep int) (*types.Block, error) {
	for i := 1; i <= retries; i++ {
		Block, err := client.BlockByNumber(context.Background(), number)
		if err == nil {
			return Block, nil
		} else {
			time.Sleep(1)
		}
	}
	return nil, errors.New("No puedo obtener el block desde el nodo")
}
