package main

import (
	"fmt"
	pb "github.com/luciano-fs/GOLatticeAgreement/protofiles"
	"google.golang.org/protobuf/proto"
)

type Elem interface {
	Join(Elem) Elem
	Leq(Elem)  bool
}

type IntSet map[int32]bool

func (a IntSet) Join(b IntSet) IntSet{
	c := make(map[int32]bool)

	for elemA,_ := range a {
		c[elemA] = true
	}
	for elemB,_ := range b {
		c[elemB] = true
	}

	return c
}

func (a IntSet) Leq(b IntSet) bool {
	for elemA,_ := range a {
		if !b[elemA] {
			return false
		}
	}
	return true
}

func main() {
    value := map[int32]bool{
		1: true,
		2: true,
		3: true,
	}

	p := &pb.Proposal{
		Value: value,
		Seq: 17,
		Uid: 23,
	}

    /*
	body, _ := proto.Marshal(p)

	p1 := &pb.Proposal{}
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)
    */

}
