package worker

import (
	"context"
	"fmt"
	"log"
	"mintfun/internal/helpers"
)

func Minter(ctx context.Context, txChan <-chan helpers.ProcessedData) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Minter stopped")
			return
		case tx := <-txChan:
			fmt.Println(tx)
		}
	}
}
