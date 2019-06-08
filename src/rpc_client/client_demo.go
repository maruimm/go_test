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
	for  {

		 <- ticks

		fmt.Printf("< start....\n")
		go func() {

			var reply string
			err = client.Call("path/to/pkg.HelloService.Hello", "btest", &reply)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(reply)
		}()
	}

	chan1 := make(chan int)
	<- chan1
}