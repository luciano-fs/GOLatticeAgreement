package main

import (
        "os"
	"net"
	"log"
	pb "github.com/luciano-fs/GOLatticeAgreement/protofiles"
	la "github.com/luciano-fs/GOLatticeAgreement/operations"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{
    Accepted map[int32]bool
    pb.UnimplementedProposeServer
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":" + os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

    pb.RegisterProposeServer(s, &server{Accepted: make(map[int32]bool)})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) MakeProposal(ctx context.Context, in *pb.Proposal) (*pb.Response, error) {
    //fmt.Println("Got proposal ", in.Value)
    //fmt.Println("Got from ", in.Uid)
    //fmt.Println("For sequence", in.Seq)
    //fmt.Println("Accepted so far:", s.Accepted)

    nack := make(map[int32]bool)

    for elemA,_ := range s.Accepted {
		if !in.Value[elemA] {
            nack[elemA] = true
		}
	}

    s.Accepted = la.Join(s.Accepted, in.Value)

    if len(nack)== 0 {
        return &pb.Response{Accept: true, Value: nil}, nil
    } else {
        return &pb.Response{Accept: false, Value: nack}, nil
    }
}
