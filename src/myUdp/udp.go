package main

import (
	"net"
	"fmt"
	"strconv"
	"os"
	"strings"
	"bufio"
	"config"
	pb "proto"
)


func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func serverUdpDemo() {

	fmt.Println(config.NAME)
	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, config.SERVER_RECV_LEN)
		_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		strData := string(data)
		fmt.Println("Received:", strData)

		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Send:", upper)
	}
}

func clientUdpDemo() {

	fmt.Printf("%v",pb.FOO_name)

	serverAddr := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	conn, err := net.Dial("udp", serverAddr)
	checkError(err)

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen - written > config.SERVER_RECV_LEN {
				toWrite = line[written : written + config.SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			checkError(err)

			fmt.Println("Write:", toWrite)

			msg := make([]byte, config.SERVER_RECV_LEN)
			n, err = conn.Read(msg)
			checkError(err)

			fmt.Println("Response:", string(msg))
		}
	}

}

func main() {

	recvsig := make(chan int)

	go serverUdpDemo()

	go clientUdpDemo()
	<- recvsig
}