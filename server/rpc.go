package main

// TODO: use JSON-RPC https://www.jsonrpc.org/specification

// Numbers represents RPC args
type Numbers struct {
	A, B int
}

// RPCServer provides RPC methods
type RPCServer int

// AddNumbers Adds two numbers
func (rpc *RPCServer) AddNumbers(numbers *Numbers, resp *int) error {
	// change the value pointed at by resp
	*resp = numbers.A + numbers.B
	return nil
}
