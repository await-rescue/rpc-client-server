package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	numbers := new(NumbersService)

	server := rpc.NewServer()
	server.Register(numbers)

	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server up")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
