package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ApiConfig struct {
	ApiUrl         string
	TransactionUrl string
}

type Collection struct {
	Contract   string   `json:"contract"`
	Deployer   string   `json:"deployer"`
	Name       string   `json:"name"`
	TotalMints string   `json:"totalMints"`
	IsReported []string `json:"isReported"`
}

type CollectionResponse struct {
	Collections []Collection `json:"collections"`
}

type Transaction struct {
	To       string `json:"to"`
	CallData string `json:"callData"`
	NftCount string `json:"nftCount"`
	EthValue string `json:"ethValue"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}

func GetApiConfig() *ApiConfig {
	return &ApiConfig{
		ApiUrl:         os.Getenv("MINTFUN_COLLECTION_URL"),
		TransactionUrl: os.Getenv("MINTFUN_TRANSACTION_URL"),
	}
}

func FetchCollection(ctx context.Context) ([]Collection, error) {

	config := GetApiConfig()

	response, err := http.Get(config.ApiUrl)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error: %s\n", response.Status)
		return nil, errors.New("failed to fetch collection data")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error reading response body %s\n", err)
		return nil, err
	}

	var collectionResponse CollectionResponse

	err = json.Unmarshal(body, &collectionResponse)
	if err != nil {
		log.Printf("Error decoding json response %s\n", err)
		return nil, err
	}

	return collectionResponse.Collections, nil
}

func GetTransaction(ctx context.Context, collection Collection) ([]Transaction, error) {

	url := fmt.Sprintf("https://mint.fun/api/mintfun/contract/%s/transactions", collection.Contract)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error: %s\n", response.Status)
		return nil, errors.New("failed to fetch transaction data")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error reading response body %s\n", err)
		return nil, err
	}
	var transactionResponse TransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		log.Printf("Error decoding json response %s\n", err)
		return nil, err
	}
	return transactionResponse.Transactions, nil
}
