package main

import (
	"net"
	"net/rpc"
	"log"
	myjson "jsonrpc/proto"
	jsonrpc "net/rpc/jsonrpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}


// server端的rpc处理器
type ServerHandler struct {}

// server端暴露的rpc方法
func (serverHandler ServerHandler) GetName(id int, returnObj *myjson.RpcObj) error {
	log.Println("server\t-", "recive GetName call, id:", id)
	returnObj.Id = id
	return nil
}

// server端暴露的rpc方法
func (serverHandler ServerHandler) SaveName(rpcObj myjson.RpcObj, returnObj *myjson.ReplyObj) error {
	log.Printf("server\t-", "recive SaveName call, RpcObj:%v", rpcObj)
	return nil
}


func main() {

	// 新建Server
	server := rpc.NewServer()

	// 开始监听,使用端口 8888
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("server\t-", "listen error:", err.Error())
	}
	defer listener.Close()

	log.Println("server\t-", "start listion on port 8888")

	// 新建处理器
	serverHandler := &ServerHandler{}

	// 注册处理器
	server.Register(serverHandler)

	// 等待并处理链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		// 在goroutine中处理请求
		// 绑定rpc的编码器，使用http connection新建一个jsonrpc编码器，并将该编码器绑定给http处理器
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
