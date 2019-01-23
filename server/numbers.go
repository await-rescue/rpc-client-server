package main

import "errors"

// NumbersArgs represents RPC args
type NumbersArgs struct {
	A, B int
}

// NumbersReply represents an RPC reply
type NumbersReply struct {
	Result float32
}

// NumbersService serves RPC methods
type NumbersService struct{}

// AddNumbers Adds two numbers
func (ns *NumbersService) AddNumbers(args *NumbersArgs, reply *NumbersReply) error {
	reply.Result = float32(args.A) + float32(args.B)
	return nil
}

// MultiplyNumbers multiplies two numbers
func (ns *NumbersService) MultiplyNumbers(args *NumbersArgs, reply *NumbersReply) error {
	reply.Result = float32(args.A) * float32(args.B)
	return nil
}

// DivideNumbers divided two numbers
func (ns *NumbersService) DivideNumbers(args *NumbersArgs, reply *NumbersReply) error {
	if args.B == 0 {
		return errors.New("Cannot divide by zero")
	}
	reply.Result = float32(args.A) / float32(args.B)
	return nil
}
