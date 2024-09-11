package worker

import (
	"context"
	"log"
	"time"

	"mintfun/internal/api"
)

func FetchWorker(ctx context.Context, dataChan chan<- []api.Collection) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Collection fetcher stopped")
			return
		case <-ticker.C:
			data, err := api.FetchCollection(ctx)
			if err != nil {
				log.Printf("Error fetching collection data: %v", err)
				continue
			}

			select {
			case dataChan <- data:

			case <-ctx.Done():
				log.Println("Data channel closed, stopping fetcher")
				return
			}
		}
	}
}
