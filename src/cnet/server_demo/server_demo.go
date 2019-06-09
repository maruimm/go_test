package main

import (
	"bufio"
	"bytes"
	ctcp "cnet/ctcp"
	"fmt"
	"net"
	_ "net"
	_ "strconv"
	"strings"
	"encoding/binary"
	"reflect"
)

const (
	DELIM byte = '\n'
)

func generateTestContent(content string) string {
	var respBuffer bytes.Buffer
	respBuffer.WriteString(strings.TrimSpace(content))
	respBuffer.WriteByte(DELIM)
	return respBuffer.String()

}

func requestHandler(showLog bool) func(conn net.Conn) {
	return func(conn net.Conn) {
		reader := bufio.NewReader(conn)
		scanner := bufio.NewScanner(reader)
		//for {
		// 自定义匹配函数
		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			fmt.Printf("atEOF:%v,data:%d\n",atEOF,len(data))
			if (len(data) >= 4) {
				bytesBuffer := bytes.NewBuffer(data)
				var x int32
				binary.Read(bytesBuffer, binary.BigEndian, &x)
				fmt.Printf("num:%d\n",x)
				num := int(reflect.TypeOf(x).Size() + uintptr(x))
				if len(data) >= num {
					return num, data[:num], nil
				}

			} else {
				if atEOF == true {
					return 0, nil, bufio.ErrInvalidUnreadByte
				} else {
					return 0, nil, nil
				}
			}



			return 0, nil, nil
		}
		// 设置匹配函数
		scanner.Split(split)
		// 开始扫描
		for scanner.Scan() {
			//fmt.Printf("%s\n", scanner.Text())
			for i,x := range scanner.Bytes() {
				fmt.Printf("[%d]:%d\n",i,x)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		//}
	}
}

func generateTcpListener(serverAddr string, showLog bool) ctcp.TcpListener {
	listener := ctcp.NewTcpListener()
	var hasError bool
	if showLog {
		fmt.Printf("Start Listening at address %s ...\n", serverAddr)
	}
	err := listener.Init(serverAddr)
	if err != nil {
		hasError = true
		fmt.Errorf("Listener Init error: %s", err)
	}
	err = listener.Listen(requestHandler(showLog))
	if err != nil {
		hasError = true
		fmt.Errorf("Listener Listen error: %s", err)
	}
	if !hasError {
		return listener
	} else {
		if listener != nil {
			listener.Close()
		}
		return nil
	}
}

func main() {
	fmt.Println("hello world")
	listener := generateTcpListener("127.0.0.1:8888", true)

	if listener == nil {
		fmt.Println("listerner err")
		return
	}

	defer listener.Close()

}
