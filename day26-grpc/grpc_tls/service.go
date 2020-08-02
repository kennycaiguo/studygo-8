package grpc

import (
	"context"
	"day26/protoes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	creads, _ := credentials.NewClientTLSFromFile("server.pem", "wang")

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer(grpc.Creds(creads))
	protoes.RegisterHelloServerServer(srv, &server{})
	err = srv.Serve(ln)
	if err != nil {
		panic(err)
	}
}
