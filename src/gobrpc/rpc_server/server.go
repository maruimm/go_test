package main

import (
	"net/rpc"
	"log"
	"net"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	time.Sleep(1*time.Second)
	*reply = "hello:" + request
	return nil
}



func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}