package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoin(ctx context.Context, in *pb.Coin) (*pb.CoinId, error) {
	log.Printf("CreateCoin was invoked with %v\n", in)

	data := CoinItem{
		Initials: in.Initials,
		Name:     in.Name,
		Votes:    in.Votes,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.CoinId{
		Id: oid.Hex(),
	}, nil

}
