package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	args := Args{18, 4}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
}
