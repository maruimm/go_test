package main

import (
	"net/rpc"
	"log"
	"fmt"
	"time"
)



func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var ticker* time.Ticker = time.NewTicker(1)
	ticks := ticker.C

	go func() {
		for _ = range ticks {
			var reply string
			err = client.Call("HelloService.Hello", "hello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_testhello_test", &reply)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(reply)
		}
	}()

	chan1 := make(chan int)
	<- chan1
}