package main

import "errors"

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

// MultiplyNumbers multiplies two numbers
func (ns *NumbersService) MultiplyNumbers(args *NumbersArgs, reply *NumbersReply) error {
	reply.Result = args.A * args.B
	return nil
}

// DivideNumbers multiplies two numbers
func (ns *NumbersService) DivideNumbers(args *NumbersArgs, reply *NumbersReply) error {
	if args.B == 0 {
		return errors.New("Cannot divide by zero")
	}
	reply.Result = args.A / args.B
	return nil
}
