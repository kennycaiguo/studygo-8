package server

import (
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Que, Rem int
}

type Arith int

func (this *Arith) Multiply (args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (this *Arith) Divide (args *Args, quo *Quotient) error {
	quo.Que = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":1234")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
