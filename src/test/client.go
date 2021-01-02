package main

import (
	"bufio"
	"fmt"
	"github/maruimm/myGoLearning/selfProto"
	"github/maruimm/myGoLearning/tcpClient"
	"time"
)

func main() {

	pool := tcpClient.NewConnPool(30,
		100,
		 60*time.Second,
		"127.0.0.1",
		8899)

	for i := 0; i < 100; i++ {
		go func(i int) {
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
					HeadLen: uint32(len(head)),
					BodyLen:uint32(len(body)),
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
				writer := bufio.NewWriter(conn)
				n, err := writer.Write(req)
				if err != nil {
					fmt.Printf("searial req failed err:%+v\n" ,err)
					return
				}
				err = writer.Flush()
				fmt.Printf("searial req failed n:%d,err:%+v\n" ,n,err)
				_ = pool.Release(conn, true)
			}
		}(i)
	}
	time.Sleep(100*time.Second)
	fmt.Printf("hello world\n")
}
