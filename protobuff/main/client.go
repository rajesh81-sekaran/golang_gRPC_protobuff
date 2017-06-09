package main

import (
	"context"
	"fmt"
	pb "protobuff/cbuf"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("192.168.33.10:9091", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	client := pb.NewConversionClient(conn)

	inpI := pb.Integer{Integer: 1}
	fmt.Println("inpI:")
	spew.Dump(inpI)
	opS, err := client.IntegerToString(context.Background(), &inpI)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ops := opS.Str
	fmt.Println("String value:", ops)

	inpS := pb.String{Str: "one"}
	fmt.Println("inpS:")
	spew.Dump(inpS)
	opI, err := client.StringToInteger(context.Background(), &inpS)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	opi := opI.Integer
	fmt.Println("String value:", opi)

}
