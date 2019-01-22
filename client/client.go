package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"time"
)

// NumbersArgs represents RPC args
type NumbersArgs struct {
	A, B int
}

// NumbersReply represents an RPC reply
type NumbersReply struct {
	Result int
}

func main() {
	fmt.Println("waiting for server")
	time.Sleep(time.Second * 5)
	fmt.Println("client up")
	client, err := rpc.DialHTTP("tcp", "server:8001")
	if err != nil {
		log.Fatal(err)
	}

	reply := new(NumbersReply)

	for {
		time.Sleep(time.Millisecond * 1000)
		args := NumbersArgs{rand.Intn(50), rand.Intn(50)}

		fmt.Println(fmt.Sprintf("Calling AddNumbers with %d and %d", args.A, args.B))

		err = client.Call("NumbersService.AddNumbers", &args, reply)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(fmt.Sprintf("response: %d", reply.Result))
	}
}
