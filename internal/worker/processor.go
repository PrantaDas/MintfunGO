package worker

import (
	"context"
	"log"
	"mintfun/internal/api"
	"mintfun/internal/helpers"
)

func TaskProcessor(ctx context.Context, dataChan <-chan []api.Collection, txChan chan<- helpers.ProcessedData) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Task Processor stopped")
			return
		case data := <-dataChan:
			helpers.BroadCast(ctx, data, txChan)
		}
	}
}
