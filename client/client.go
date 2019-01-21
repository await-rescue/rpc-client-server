package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"time"
)

// Numbers represents RPC args
type Numbers struct {
	A, B int
}

func main() {
	fmt.Println("waiting for server")
	time.Sleep(time.Second * 5)
	fmt.Println("client up")
	client, err := rpc.DialHTTP("tcp", "server:8001")
	if err != nil {
		log.Fatal(err)
	}

	var reply int

	for {
		time.Sleep(time.Millisecond * 100)
		args := Numbers{rand.Intn(50), rand.Intn(50)}

		fmt.Println(fmt.Sprintf("Calling AddNumbers with %d and %d", args.A, args.B))

		err = client.Call("RPCServer.AddNumbers", args, &reply)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(fmt.Sprintf("response: %d", reply))
	}
}
