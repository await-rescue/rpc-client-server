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
	Result float32
}

func main() {
	fmt.Println("Waiting for server")
	time.Sleep(time.Second * 2)

	conn, err := net.Dial("tcp", "server:8001")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	jsonCodec := jsonrpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(jsonCodec)

	methods := make([]string, 0)
	methods = append(methods,
		"NumbersService.AddNumbers",
		"NumbersService.MultiplyNumbers",
		"NumbersService.DivideNumbers")

	for {
		time.Sleep(time.Millisecond * 500)

		args := NumbersArgs{rand.Intn(5), rand.Intn(5)}
		reply := new(NumbersReply)

		rand.Seed(time.Now().Unix())
		method := methods[rand.Intn(len(methods))]

		log.Printf("Calling %s with params %d and %d", method, args.A, args.B)
		err = client.Call(method, args, reply)
		if err != nil {
			log.Println(fmt.Sprintf("ERROR: %s", err))
			continue
		}

		res, _ := json.Marshal(reply)
		log.Println(fmt.Sprintf("Response received: %s", string(res)))
	}
}
