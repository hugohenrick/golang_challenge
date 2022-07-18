package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpVoteCoin(ctx context.Context, in *pb.Coin) (*emptypb.Empty, error) {
	log.Printf("UpVoteCoin was invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	filter := bson.M{"_id": oid}

	collection.FindOneAndUpdate(ctx, filter, bson.M{"$inc": bson.M{"votes": 1}}, options.FindOneAndUpdate().SetReturnDocument(1))

	return &emptypb.Empty{}, nil
}
