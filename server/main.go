package main

import (
	"fmt"
	"net"
	"log"
	pb "github.com/luciano-fs/GOLatticeAgreement/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{
    pb.UnimplementedProposeServer
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterProposeServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) MakeProposal(ctx context.Context, in *pb.Proposal) (*pb.Response, error) {
	fmt.Println("Got proposal ", in.Value)
	fmt.Println("Got from ", in.Uid)
	fmt.Println("For sequence", in.Seq)
    m := make(map[int32]bool)
    return &pb.Response{Accept: true, Value: m}, nil
}
