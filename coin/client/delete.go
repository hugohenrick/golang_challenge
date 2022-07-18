package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
)

func deleteCoin(c pb.CoinServiceClient, id string) {
	log.Println("---deleteCoin was invoked---")

	_, err := c.DeleteCoin(context.Background(), &pb.CoinId{Id: id})

	if err != nil {
		log.Printf("Error happened while deleting: %v\n", err)
	}

	log.Println("Coin was deleted!")

}
