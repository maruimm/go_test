package main

import (
	"gobrpc/comm"
	"net/rpc"
	"net"
	"time"
)

func main() {
	rpc.Register(new(comm.HelloService))

	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}

}