package main

import (
	_ "net"
	"net/rpc"
	_ "log"
	_ "net/http"
	_ "net/rpc/jsonrpc"
	_ "io"
	"log"
	"net"
	"time"
)

// 需要传输的对象
type RpcObj struct {
	Id   int `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
	Name string `json:"name"`
}

// 需要传输的对象
type ReplyObj struct {
	Ok  bool `json:"ok"`
	Id  int `json:"id"`
	Msg string `json:"msg"`
}

// server端的rpc处理器
type ServerHandler struct {}

// server端暴露的rpc方法
func (serverHandler ServerHandler) GetName(id int, returnObj *RpcObj) error {
	log.Println("server\t-", "recive GetName call, id:", id)
	returnObj.Id = id
	returnObj.Name = "名称1"
	return nil
}

// server端暴露的rpc方法
func (serverHandler ServerHandler) SaveName(rpcObj RpcObj, returnObj *ReplyObj) error {
	log.Println("server\t-", "recive SaveName call, RpcObj:", rpcObj)
	returnObj.Ok = true
	returnObj.Id = rpcObj.Id
	returnObj.Msg = "存储成功"
	return nil
}

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	time.Sleep(1*time.Second)
	*reply = "hello:" + request
	return nil
}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

const HelloServiceName = "path/to/pkg.HelloService"

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {

	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}

}
