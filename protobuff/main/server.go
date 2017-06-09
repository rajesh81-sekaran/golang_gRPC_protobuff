package main

import (
	"fmt"
	"net"
	pb "protobuff/cbuf"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type server struct{}
type myError struct {
	err string
}

func (me myError) Error() string {
	return me.err
}

func (s server) IntegerToString(ctx context.Context, from *pb.Integer) (op *pb.String, me error) {
	op = &pb.String{}
	switch from.Integer {
	case 1:
		op.Str = "one"

	case 2:
		op.Str = "two"

	case 3:
		op.Str = "three"

	case 4:
		op.Str = "four"

	case 5:
		op.Str = "five"

	default:
		op.Str = "zero"
		me = myError{err: "Only digits 1 to 5 are supported defaulting to \"zero\""}
	}
	return
}

func (s server) StringToInteger(ctx context.Context, from *pb.String) (op *pb.Integer, me error) {
	op = &pb.Integer{}
	switch from.Str {
	case "one":
		op.Integer = 1

	case "two":
		op.Integer = 2

	case "three":
		op.Integer = 3

	case "four":
		op.Integer = 4

	case "five":
		op.Integer = 5

	default:
		op.Integer = 0
		me = myError{err: "Only strings \"one\", \"two\" \"three\" \"four\" and \"five\" are supported"}
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterConversionServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
