package helpers

import (
	"context"
	"log"
	"mintfun/internal/api"
)

func BroadCast(ctx context.Context, colls []api.Collection, txChan chan<- ProcessedData) {
	for _, col := range colls {
		transactions, err := api.GetTransaction(ctx, col)
		if err != nil {
			log.Println("Error getting transaction")
			continue
		}
		data, err := ProcessData(ctx, col, transactions)
		if err != nil {
			log.Println("Error processing transaction data")
			continue
		}
		txChan <- data
	}
}
