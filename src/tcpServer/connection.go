package tcpServer

import (
	"container/list"
	"net"
	"sync"
)

type myConnection struct {
	connections list.List
	lk sync.RWMutex
}


type connHandler *list.Element

func (s *myConnection) saveConn(conn net.Conn) connHandler {
	s.lk.Lock()
	defer s.lk.Unlock()
	return s.connections.PushFront(conn)
}

func (s *myConnection) freeConn(conn *list.Element) {
	s.lk.Lock()
	defer s.lk.Unlock()
	s.connections.Remove(conn)
}

func (s *myConnection) traversalAndCloseConns() {
	s.lk.Lock()
	defer s.lk.Unlock()
	for el := s.connections.Front(); el != nil; el = el.Next() {
		conn, ok := el.Value.(net.Conn)
		if !ok {
			PrintLog(LOGERROR, "el type is failed %+v\n", el)
			return
		}
		conn.Close()
	}
}

func (s *myConnection) realConn(connHandler connHandler) (net.Conn, error){
	conn, ok := connHandler.Value.(net.Conn)
	if !ok {
		return nil,SysError
	}
	return conn, nil
}
