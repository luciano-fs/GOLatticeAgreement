package main

import (
    "strconv"
    "reflect"
	"log"
	pb "github.com/luciano-fs/GOLatticeAgreement/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    "os"
	la "github.com/luciano-fs/GOLatticeAgreement/operations"
)

func propose(c pb.ProposeClient, value map[int32]bool, seq int32, id int32, r chan pb.Response) {
    resp,_ := c.MakeProposal(
		context.Background(),
		&pb.Proposal {
            Value: value,
            Seq: seq,
            Uid: id,
		},
	)

    if resp != nil {
	    r <- *resp
	    close(r)
    }
}

func main() {
    id,_ := strconv.Atoi(os.Args[1])
    uid := int32(id)

    value := map[int32]bool {
        int32(uid): true,
    }

    n,_ := strconv.Atoi(os.Args[2])
    f := n/2
    addr := os.Args[3:]

    rep := make([]pb.ProposeClient, n)
    for i := 0; i<n; i++ {
        conn, err := grpc.Dial(addr[i], grpc.WithInsecure())
        if err != nil {
            log.Println("Error while making connection, %v", err)
        }
        rep[i] = pb.NewProposeClient(conn)
    }

    var seq int32
    seq = 1
    for {
        //log.Println("Beginning iteration ", seq)
        chans := make([]chan pb.Response, n)
        cases := make([]reflect.SelectCase, n)
        for i := 0; i < n; i++ {
            ch := make(chan pb.Response)
            chans = append(chans, ch)
            go propose(rep[i], value, seq, uid, ch)
            cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
        }

        remaining := n-f
        nack := make(map[int32]bool)
        for remaining > 0 {
            //log.Println("Remaining responses: ", remaining)
            active, rresp, ok := reflect.Select(cases)
            if !ok {
                // The chosen channel has been closed, so zero out the channel to disable the case
                cases[active].Chan = reflect.ValueOf(nil)
                remaining -= 1
                continue
            }
            response := rresp.Interface().(pb.Response)
            //log.Println("Got a response from ", active)
            //log.Println(response.Accept, " ", response.Value)

            if response.Accept == false {
                nack = la.Join(response.Value, nack)
                //log.Println("Current nack value: ", nack)
            }
        }

        if len(nack) == 0 {
            log.Println(uid," decides ", value)
            break
        }

        value = la.Join(value, nack)
        seq++
    }
}
