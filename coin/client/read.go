package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
)

func readCoin(c pb.CoinServiceClient, id string) *pb.Coin {
	log.Println("---readCoin was invoked---")

	req := &pb.CoinId{Id: id}
	res, err := c.ReadCoin(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Coin was read: %v\n", res)
	return res
}
