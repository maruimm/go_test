package main

import (
	"net/rpc"
	"log"
	"fmt"
	"time"
	"math/rand"
)



func doClientWork(client *rpc.Client) {

	r := rand.Int()
	t := time.Now().UnixNano()
	fmt.Printf("send start...%d time now:%d\n",r,t)
	helloCall := client.Go("path/to/pkg.HelloService.Hello", "hello", new(string), nil)

	// do some thing

	helloCall = <-helloCall.Done
	fmt.Printf("send end...%d dur:%d\n",r,time.Now().UnixNano() - t)
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}


func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var ticker* time.Ticker = time.NewTicker(1)
	ticks := ticker.C
	for  {

		 <- ticks

		//fmt.Printf("< start....\n")
		go func() {

			//var reply string
			//err = client.Call("path/to/pkg.HelloService.Hello", "btest", &reply)

			doClientWork(client)

			if err != nil {
				log.Fatal(err)
			}

			//fmt.Println(reply)
		}()
	}

	chan1 := make(chan int)
	<- chan1
}