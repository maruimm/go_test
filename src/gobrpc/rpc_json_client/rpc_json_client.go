package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"flag"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	var b = flag.Int("b",10,"执行次数")
	flag.Parse()
	log.Printf("b:%v",*b)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	for i := 0; i < *b; i++ {
		var reply string
		err = client.Call("HelloService.Hello", "hello", &reply)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("i:%d,%s\n",i ,reply)
	}
}
