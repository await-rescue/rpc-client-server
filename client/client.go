package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
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

	jsonCodec := jsonrpc.NewClientCodec(conn)

	client := rpc.NewClientWithCodec(jsonCodec)

	for {
		args := NumbersArgs{rand.Intn(50), rand.Intn(50)}
		reply := new(NumbersReply)

		err = client.Call("NumbersService.AddNumbers", args, reply)
		if err != nil {
			log.Println(err)
		}

		res, _ := json.Marshal(reply)
		fmt.Println(string(res))
		time.Sleep(time.Millisecond * 1000)
	}
}
