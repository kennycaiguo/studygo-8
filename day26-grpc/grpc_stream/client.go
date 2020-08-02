package main

import (
	"bufio"
	"context"
	"day26/protoes"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Println("连接失败：[%v]", err)
		return
	}
	defer conn.Close()
	client := protoes.NewChatClient(conn)
	ctx := context.Background()
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Println("创建数据流失败：[%v]", err)
		return
	}
	go func() {
		log.Println("请输入消息：")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if err := stream.Send(&protoes.Request{Input:scanner.Text()}); err != nil {
				log.Println(err)
				return
			}
		}
	}()
	for {
		响应, err := stream.Recv()
		if err == io.EOF {
			log.Println("收到服务端的结束命令")
			break
		}
		if err != nil {
			log.Println("接受数据出错：", err)
		}
		log.Printf("【客户端接受到】: %v \n",响应.Output)
	}
}
