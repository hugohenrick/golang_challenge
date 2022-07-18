package main

import (
	pb "github.com/hugohenrick/golang_challenge/coin/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CoinItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Initials string             `bson:"initials"`
	Name     string             `bson:"name"`
	Votes    int32              `bson:"votes"`
}

func documentToCoin(data *CoinItem) *pb.Coin {
	return &pb.Coin{
		Id:       data.ID.Hex(),
		Initials: data.Initials,
		Name:     data.Name,
		Votes:    data.Votes,
	}
}
