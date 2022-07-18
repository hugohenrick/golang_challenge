package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
)

func updateCoin(c pb.CoinServiceClient, id string) {
	log.Println("---updateCoin was invoked---")

	newCoin := &pb.Coin{
		Id:       id,
		Initials: "BTC2",
		Name:     "Bitcoin 2.0",
		Votes:    1000,
	}

	_, err := c.UpdateCoin(context.Background(), newCoin)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Coin was updated!")
}
