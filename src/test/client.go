package main

import (
	"fmt"
	"github/maruimm/myGoLearning/selfProto"
	"github/maruimm/myGoLearning/tcpClient"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	pool := tcpClient.NewConnPool(3,
		10,
		 60*time.Second,
		"127.0.0.1",
		8899)

	var wg sync.WaitGroup
	var count int32

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10240 ; j++ {
				conn, err := pool.Acquire(2 * time.Second)
				//_ = conn.(net.Conn)
				if err != nil {
					fmt.Printf("acquire error:%+v\n", err)
					continue
				}
				deCencer := selfProto.NewDecEncer()
				head := fmt.Sprintf("head i:%d,j:%d\n", i,j)
				body := fmt.Sprintf("body i:%d,j:%d\n", i,j)
				proto := selfProto.Proto{
					Start:0xee,
					Head:[]byte(head),
					Body:[]byte(body),
					End:0xff,
				}
				buff, err := deCencer.Enc(proto)
				if err != nil {
					fmt.Printf("searial req failed err:%+v\n" ,err)
					return
				}

				req := buff.Bytes()
				n, err := conn.Write(req)
				if err != nil || n != len(req) {
					fmt.Printf("searial req failed err:%+v\n" ,err)
					_ = pool.Release(conn, false)
					return
				}
				atomic.AddInt32(&count,1)
				//fmt.Printf("searial req n:%d,err:%+v,req:%+v\n" ,n,err,req)
				_ = pool.Release(conn, true)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("hello world:%d\n",count)
}
