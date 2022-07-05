package main

import (
	"log"
	pb "github.com/luciano-fs/GOLatticeAgreement/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//. "github.com/luciano-fs/GOLatticeAgreement/types"
)

const address = "localhost:8000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	c := pb.NewProposeClient(conn)

    value := map[int32]bool{
		1: true,
		2: true,
		3: true,
	}

	c.MakeProposal(
		context.Background(),
		&pb.Proposal {
            Value: value,
            Seq: 17,
            Uid: 23,
		},
	)
}
