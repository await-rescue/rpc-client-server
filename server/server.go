package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	server := new(NumbersService)
	rpc.Register(server)
	rpc.HandleHTTP()

	fmt.Println("server up")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
