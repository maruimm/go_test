package main

import (
	"cnet/ctcp"
	"fmt"
	"time"
	_ "io"
	"bytes"
	"encoding/binary"
	"flag"
	"math/rand"
)


//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}


func main() {
	sender := ctcp.NewTcpSender()
	err := sender.Init("127.0.0.1:8888",1*time.Second)
	if err != nil {
		fmt.Printf("err:%v\n",err)
	}

	bytesBuffer := bytes.NewBuffer([]byte{})

	var b = flag.Int("b",10,"每次执行的个数")
	var n = flag.Int("n",10,"每次执行的个数")
	flag.Parse()

	for index := 0; index < *n; index++ {

		bytesBuffer.Write(IntToBytes(*b))
		fmt.Printf("buf len:%d,buf cap:%d\n", len(bytesBuffer.Bytes()), cap(bytesBuffer.Bytes()))

		for i := 0; i < *b; i++ {
			n, _ := bytesBuffer.Write([]byte{byte(rand.Int())})
			fmt.Printf("write %d byte\n", n)
		}

		fmt.Printf("buf len:%d,buf cap:%d\n", len(bytesBuffer.Bytes()), cap(bytesBuffer.Bytes()))

		for i, x := range bytesBuffer.Bytes() {
			fmt.Printf("[%d]%d\n", i, x)
		}
		fmt.Printf("buf len:%d,buf cap:%d\n", len(bytesBuffer.Bytes()), cap(bytesBuffer.Bytes()))
	}
	bytesBuffer.Write([]byte{2})
	sender.SendBytes(bytesBuffer.Bytes())

}