package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
)

func CreateCoin(c pb.CoinServiceClient) string {
	log.Println("---CreateCoin was invoked---")

	coin := &pb.Coin{
		Initials: "BTC",
		Name:     "Bitcoin",
		Votes:    0,
	}

	res, err := c.CreateCoin(context.Background(), coin)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Coin has been created: %s\n", res.Id)
	return res.Id
}
