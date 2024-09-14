package worker

import (
	"context"
	"log"
	"mintfun/internal/db"
	"mintfun/internal/helpers"
	"mintfun/internal/web3"
)

func Minter(ctx context.Context, db *db.MongoDBPersister, txChan <-chan helpers.ProcessedData, wallet *web3.Wallet) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Minter stopped")
			return
		case tx := <-txChan:
			helpers.Transaction(ctx, tx, db, wallet)
		}
	}
}
