package main

import (
	"log"
	"net/rpc"
	"net"
	_ "fmt"
	"self_proto/comm"
)


func main() {

	svr := rpc.NewServer()

	comm.RegisterHelloService(svr,new(comm.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go svr.ServeConn(conn)
	}
}
