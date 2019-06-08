package main

import (
	comm "self_proto/comm"
	"log"
	"flag"
	"time"
	"sync"
)


func main() {
	client, err := comm.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var b = flag.Int("b",10,"执行次数")
	flag.Parse()
	log.Printf("b:%v",*b)

	var wg sync.WaitGroup
	wg.Add(*b)
	for i:=1; i<*b;i++ {
		go func(i int) {
			var reply comm.String
			hello := comm.String{Value: "ruima"}
			err = client.Hello(hello, &reply)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("rsp[%d]:%s\n", i, reply.GetValue())
			wg.Add(1)
		}(i)
		time.Sleep(1)
	}
	wg.Wait()
}
