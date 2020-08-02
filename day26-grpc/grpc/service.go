package grpc

import (
	"context"
	"day26/protoes"
	"google.golang.org/grpc"
	"net"
)

type server struct {}

func (this *server) SayHi(ctx context.Context, in *protoes.HelloRequest) (*protoes.HelloReplay, error) {
	return &protoes.HelloReplay{Message:"hi"}, nil
}
func (this *server) GetMsg(ctx context.Context, in *protoes.HelloRequest) (*protoes.HelloMessage, error) {
	return &protoes.HelloMessage{Msg:"hello"}, nil
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	protoes.RegisterHelloServerServer(srv, &server{})
	err = srv.Serve(ln)
	if err != nil {
		panic(err)
	}
}
