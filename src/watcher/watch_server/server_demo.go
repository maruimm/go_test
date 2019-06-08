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
	"sync"
	"fmt"
)

type KVStoreService struct {}

func (p *KVStoreService) Hello(request string, reply *string) error {
	time.Sleep(1*time.Second)
	*reply = "hello:" + request
	return nil
}

type KVStoreServiceInterface = interface {
	Hello(request string, reply *string) error
}

const KVStoreServiceName = "KVStoreService"

func RegisterHelloService(svc KVStoreServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}


type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
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
