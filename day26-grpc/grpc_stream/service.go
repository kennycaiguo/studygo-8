package main

import (
	"day26/protoes"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type Streamer struct {}

func (this *Streamer)BidStream(stream protoes.Chat_BidStreamServer) error{
	ctx := stream.Context()
	select {
	case <- ctx.Done():
		log.Println("收到客户端通过context发出的终止信号")
		return ctx.Err()
	default:
		输入, err := stream.Recv()
		if err == io.EOF {
			log.Println("客户端发送数据结束")
			return nil
		}
		if err != nil {
			log.Println("接受数据出错")
			return err
		}
		switch 输入.Input {
		case "结束对话":
			log.Println("收到 结束对话 指令")
			if err := stream.Send(&protoes.Response{Output:"收到结束命令"}); err != nil {
				return err
			}
		case "返回数据流":
			log.Println("收到 返回数据流 指令")
			for i := 0; i < 10; i++ {
				if err := stream.Send(&protoes.Response{Output:"数据流 # " + strconv.Itoa(i)}); err != nil {
					return err
				}
			}
		default:
			log.Println("[收到消息]", 输入.Input)
			if err := stream.Send(&protoes.Response{Output:"服务端返回："+ 输入.Input}); err != nil {
				return err
			}
		}
	}
	return nil
}


func main() {
	log.Println("服务端启动")
	server := grpc.NewServer()
	protoes.RegisterChatServer(server, &Streamer{})
	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
