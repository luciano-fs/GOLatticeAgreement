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

type IntSet map[int]bool

func (a IntSet) Join(b IntSet) IntSet{
	c := make(map[int]bool)

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
	p := &pb.Person{
		Id: 1234,
		Name: "John Doe",
		Email: "test@test.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-444", Type: pb.Person_HOME},
		},
	}

	// Serializing the struct and assigning it to body
	body, _ := proto.Marshal(p)

	// De-serializing the body and saving it to p1 for testing
	p1 := &pb.Person{}
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)

}
