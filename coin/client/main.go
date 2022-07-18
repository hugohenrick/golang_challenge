package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/hugohenrick/golang_challenge/coin/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewCoinServiceClient(conn)

	CreateCoin(c)
	//id := CreateCoin(c)
	//readCoin(c, id) //Valid
	//readCoin(c, "aNomExistingId") //Invalid
	//updateCoin(c, id)
	//listCoin(c)
	//deleteCoin(c, id)

}
