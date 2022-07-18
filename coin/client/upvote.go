package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
)

func upVoteCoin(c pb.CoinServiceClient, id string) {
	log.Println("---upVoteCoin was invoked---")

	newCoin := &pb.Coin{
		Id: id,
	}

	_, err := c.UpVoteCoin(context.Background(), newCoin)

	if err != nil {
		log.Printf("Error happened while upvoting: %v\n", err)
	}

	log.Println("Coin was upvoted!")
}
