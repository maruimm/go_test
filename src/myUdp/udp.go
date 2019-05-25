package main

import (
	"net"
	"fmt"
	"strconv"
	"os"
	"config"
	pb "myProto"
	"github.com/golang/protobuf/proto"
)


func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func serverUdpDemo() {

	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	addr, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		fmt.Println("err:",err)
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
			fmt.Println("server:",err)
			continue
		}
		test := &pb.Test{}
		proto.Unmarshal(data, test)
		fmt.Printf("server get data:%v\n",test)

		_, err = conn.WriteToUDP([]byte("ruima..."), rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}

func clientUdpDemo() {


	test := &pb.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}

	data, _ := proto.Marshal(test)


	serverAddr := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	conn, err := net.Dial("udp", serverAddr)
	checkError(err)

	defer conn.Close()

	toWrite := data

	n, err := conn.Write([]byte(toWrite))
	checkError(err)

	fmt.Printf("client Write:%v,%d\n", toWrite,n)

	msg := make([]byte, config.SERVER_RECV_LEN)
	n, err = conn.Read(msg)
	checkError(err)

	//fmt.Println("Response:", msg))


}

func main() {

	recvsig := make(chan int)
	go clientUdpDemo()
	go serverUdpDemo()

	//go clientUdpDemo()
	<- recvsig
}