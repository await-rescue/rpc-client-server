package main

// TODO: use JSON-RPC https://www.jsonrpc.org/specification

// NumbersArgs represents RPC args
type NumbersArgs struct {
	A, B int
}

// NumbersReply represents an RPC reply
type NumbersReply struct {
	Result int
}

// NumbersService serves RPC methods
type NumbersService struct{}

// AddNumbers Adds two numbers
func (ns *NumbersService) AddNumbers(args *NumbersArgs, reply *NumbersReply) error {
	reply.Result = args.A + args.B
	return nil
}
