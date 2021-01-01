package main

import (
	"fmt"
	"github/maruimm/myGoLearning/tcpClient"
	"time"
)

func main() {

	pool := tcpClient.NewConnPool()

	for i := 0; i < 100; i++ {
		go func(i int) {
			for j := 0; j < 10240 ; j++ {
				conn, err := pool.Acquire(2 * time.Second)
				//_ = conn.(net.Conn)
				if err != nil {
					fmt.Printf("acquire error:%+v\n", err)
					continue
				}
				time.Sleep(5*time.Second)
				_, _ = conn.Write([]byte(fmt.Sprintf("hello world..i:%d,j:%d\n", i,j)))
				_ = pool.Release(conn, true)
			}
		}(i)
	}
	time.Sleep(100*time.Second)
	fmt.Printf("hello world\n")
}
