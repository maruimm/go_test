package tcpServer

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github/maruimm/myGoLearning/selfProto"
	"net"
	"reflect"
	"sync"
	"sync/atomic"
)

type Server interface {
	Run() error
	ShutDown()
}

type tcpServer struct {
	addr  string
	port  uint16
	wg    *sync.WaitGroup
	ctx   context.Context
	stopd chan struct{}
	lis   net.Listener
	conns myConnection
	count int32
}

func NewTcpServer(ctx context.Context,
	addr string,
	port uint16) (Server, error) {
	s := &tcpServer{
		addr:  addr,
		port:  port,
		ctx:   ctx,
		wg:    &sync.WaitGroup{},
		stopd: make(chan struct{}),
	}
	var err error
	s.lis, err = net.Listen("tcp",
		fmt.Sprintf("%s:%d", s.addr, s.port))
	if err != nil {
		return nil, err
	}
	return s, nil
}

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {

	deCencer := selfProto.NewDecEncer()
	//fmt.Printf("atEOF:%v,dataLen:%d,data:%+v,StartLen:%d\n",atEOF,len(data),data,deCencer.StartLen())
	if uint32(len(data)) >= deCencer.StartLen() {
		bytesBuffer := bytes.NewBuffer(data)
		var startFlag uint16
		var headLen, bodyLen uint32
		_ = binary.Read(bytesBuffer, binary.LittleEndian, &startFlag)
		_ = binary.Read(bytesBuffer, binary.LittleEndian, &headLen)
		_ = binary.Read(bytesBuffer, binary.LittleEndian, &bodyLen)
		//fmt.Printf("startFlag:%d,headLen:%d,bodyLen:%d\n",startFlag,headLen,bodyLen)
		num := int(reflect.TypeOf(startFlag).Size() +
			reflect.TypeOf(headLen).Size() +
			reflect.TypeOf(bodyLen).Size() +
			uintptr(headLen) +
			uintptr(bodyLen) +
			reflect.TypeOf(startFlag).Size())
		//fmt.Printf("data:%+v,len(data):%d,num:%d\n",data,len(data),num)
		if len(data) >= num {
			return num, data[:num], nil
		}
	} else {
		if atEOF == true {
			fmt.Printf("retrun :%+v\n",bufio.ErrInvalidUnreadByte)
			return 0, nil, bufio.ErrInvalidUnreadByte
		} else {
			return 0, nil, nil
		}
	}
	return 0, nil, nil
}

func (s *tcpServer) handleConn(connHandler connHandler) {

	defer func() {
		fmt.Printf("s.count:%d\n",s.count)
	}()

	conn, err := s.conns.realConn(connHandler)
	if err != nil {
		PrintLog(LOGERROR, "real conn error:%+v\n", err)
		return
	}
	defer s.wg.Done()
	defer func() {
		PrintLog(LOGDEBUG, "self connection closed\n")
		s.conns.freeConn(connHandler)
		conn.Close()
	}()
	writer := bufio.NewWriter(conn)

	defer writer.Flush()
	scanner := bufio.NewScanner(conn)
	scanner.Split(split)
LOOP:
	for {
		deCencer := selfProto.NewDecEncer()

		//_ = conn.SetDeadline(time.Now().Add(15 * time.Second))
		for {

			ok :=  scanner.Scan()
			//PrintLog(LOGDEBUG, "Scan result:%+v\n", ok)
			if ok  == false{
				break LOOP
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input: %s\n", err)
				return
			}
			buff := bytes.NewBuffer(scanner.Bytes())
			req, err := deCencer.Dec(buff)
			if err != nil {
				PrintLog(LOGERROR, "deCencer de err:%+v\n", err)
				return
			}
			PrintLog(LOGDEBUG, "raw req:%+v\n", scanner.Bytes())
			PrintLog(LOGDEBUG, "req:%+v, remote port:%+v\n", req,conn.RemoteAddr())
			atomic.AddInt32(&s.count, 1)
			//_, _ = writer.Write(scanner.Bytes())
			//_ = writer.Flush()
		}
	}
}

func (s *tcpServer) ShutDown() {
	err := s.lis.Close()
	if err != nil {
		PrintLog(LOGDEBUG, "error:%+v\n", err)
	}
	<-s.stopd
}

func (s *tcpServer) Run() error {

	defer close(s.stopd)
	defer s.wg.Wait()
LOOP:
	for {
		conn, err := s.lis.Accept()
		if err != nil {
			PrintLog(LOGDEBUG, "accept exit...\n")
			s.conns.traversalAndCloseConns()
			break LOOP
		}
		el := s.conns.saveConn(conn)
		s.wg.Add(1)
		go s.handleConn(el)
	}
	return nil
}
