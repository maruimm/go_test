package main

import (
	"net/rpc"
	"log"
	"fmt"
	"time"
	"math/rand"
	"flag"
)



func doClientWorkSet(client *rpc.Client,key string,value string) {

	r := rand.Int()
	t := time.Now().UnixNano()
	fmt.Printf("send start...%d time now:%d\n",r,t)
	err := client.Call("KVStoreService.Set", [2]string{key,value}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}

}


func doClientWorkGet(client *rpc.Client,key string) {

	r := rand.Int()
	t := time.Now().UnixNano()
	fmt.Printf("send start...%d time now:%d\n",r,t)
	var value string
	err := client.Call("KVStoreService.Get" , key,&value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get value from sever:",value)

}


func doClientWorkWatch(client *rpc.Client, t int) {

	var value string
	err := client.Call("KVStoreService.Watch", t,&value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get value from sever:",value)

}


func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var b = flag.String("b","Set","Set/Get")
	var k = flag.String("k","defaultkey","key")
	var v = flag.String("v","defaultvalue","value")
	var t = flag.Int("t",0,"watch time out")
	flag.Parse()
	oper := *b
	switch oper {
		case "Get":
			doClientWorkGet(client,*k)
			break
		case "Set":
			doClientWorkSet(client,*k,*v)
			break
		case "Watch":
			doClientWorkWatch(client,*t)
			break
		default:
			fmt.Println("unkown oper")
	}
}