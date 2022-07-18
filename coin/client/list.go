package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listCoin(c pb.CoinServiceClient) {
	log.Println("---listCoin was invoked---")

	stream, err := c.ListCoins(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Printf("Error happened while ListCoins: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something Happened: %v\n", err)
		}

		log.Println(res)
	}

}
