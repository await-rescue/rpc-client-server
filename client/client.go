package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc/jsonrpc"
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
	time.Sleep(time.Second * 2)

	conn, err := net.Dial("tcp", "server:8001")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	for {
		args := NumbersArgs{rand.Intn(50), rand.Intn(50)}
		reply := new(NumbersReply)

		err = client.Call("NumbersService.AddNumbers", args, reply)
		if err != nil {
			panic(err)
		}

		log.Println(reply.Result)
		time.Sleep(time.Millisecond * 1000)
	}
}
