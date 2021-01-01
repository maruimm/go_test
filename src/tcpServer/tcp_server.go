package tcpServer

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"sync"
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

func (s *tcpServer) handleConn(connHandler connHandler) {
	conn, err := s.conns.realConn(connHandler)
	if err != nil {
		PrintLog(LOGERROR, "real conn error:%+v\n", err)
		return
	}
	defer s.wg.Done()
	defer func() {
		PrintLog(LOGDEBUG, "connection closed")
		s.conns.freeConn(connHandler)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	session := bufio.NewReadWriter(reader, writer)
	defer session.Flush()

	for {
		//_ = conn.SetDeadline(time.Now().Add(15 * time.Second))
		content, err := session.ReadBytes('\n')
		if err != nil {
			PrintLog(LOGDEBUG, "err:%+v\n", err)
			return
		}
		_, _ = session.Write(content)
		_ = session.Flush()
		PrintLog(LOGDEBUG, "content:%s\n", content)
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
