package main

import (
	"context"
	"log"
	"mintfun/internal/api"
	"mintfun/internal/db"
	"mintfun/internal/helpers"
	"mintfun/internal/worker"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	_, err = db.ConnectToMongoDB(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}

	dataChannel := make(chan []api.Collection, 4)
	defer close(dataChannel)

	txChannel := make(chan helpers.ProcessedData, 4)
	defer close(txChannel)

	log.Println("Bot Started")

	go worker.FetchWorker(ctx, dataChannel)
	go worker.TaskProcessor(ctx, dataChannel, txChannel)
	go worker.Minter(ctx, txChannel)

	<-ctx.Done()
	log.Println("Shuttiing down gracefully...")
}
