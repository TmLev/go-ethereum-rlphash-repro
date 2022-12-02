package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelCtx()

	{
		bscClient, err := ethclient.Dial("https://bsc-dataseed.binance.org")
		panicOnError(err, "could not dial bsc node")

		bscTxHash := common.HexToHash("0xb72a06e15ccf3af1283d71f66abff02cdc4c023f673203f2c319538d96bdf50c")
		bscTx, isPending, err := bscClient.TransactionByHash(ctx, bscTxHash)
		panicOnError(err, "could not get bsc tx")
		assert(!isPending, "bsc tx is pending")

		log.Print("BSC:")
		log.Printf("original tx hash:  %s", bscTxHash.String())
		log.Printf("tx hash from node: %s", bscTx.Hash().String())
	}

	{
		meterClient, err := ethclient.Dial("https://rpc.meter.io")
		panicOnError(err, "could not dial meter node")

		meterTxHash := common.HexToHash("0x37656fc5eb232510cb47b2db7b90fe0763d6ebbc2ef0578409a6aac3ae24dcf5")
		meterTx, isPending, err := meterClient.TransactionByHash(ctx, meterTxHash)
		panicOnError(err, "could not get meter tx")
		assert(!isPending, "meter tx is pending")

		log.Print("Meter:")
		log.Printf("original tx hash:  %s", meterTxHash.String())
		log.Printf("tx hash from node: %s", meterTx.Hash().String())
	}
}

func panicOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %+v", msg, err)
	}
}

func assert(cond bool, msg string) {
	if !cond {
		log.Fatalf("assertion failed: %s", msg)
	}
}
