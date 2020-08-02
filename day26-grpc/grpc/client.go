package grpc

import (
	"context"
	"day26/protoes"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := protoes.NewHelloServerClient(conn)
	ctx := context.Background()
	r1, err := c.SayHi(ctx, &protoes.HelloRequest{Name:"xiaoxiao"})
	r2, err := c.GetMsg(ctx, &protoes.HelloRequest{Name:"xiaoxiao"})
	fmt.Println(r1)
	fmt.Println(r2)
}
