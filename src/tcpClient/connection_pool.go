package tcpClient

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type ConnPool interface {
	Release(closer io.ReadWriteCloser, pooled bool) error
	Acquire(timeWait time.Duration) (io.ReadWriteCloser, error)
}

type myConn struct{
	conn io.ReadWriteCloser
	lasUsedTime time.Time
}

func (c myConn) Read(p []byte) (n int, err error) {
	return c.conn.Read(p)
}

func (c myConn) Write(p []byte) (n int, err error) {
	return c.conn.Write(p)
}

func (c myConn) Close() error {
	return c.conn.Close()
}


type factoryConn func(network, address string) (net.Conn, error)

type myConnPool struct {
	conns list.List
	lk *sync.RWMutex

	idleTime time.Duration
	maxConnCount int
	idleConnCount int
	usedChan chan struct{}
	factoryConn
	addr string
	port uint16
}

func NewConnPool() ConnPool {
	pool := &myConnPool{
		idleConnCount:2,
		maxConnCount:5,
		idleTime:10* time.Second,
		lk: new(sync.RWMutex),
		factoryConn: net.Dial,
		addr: "127.0.0.1",
		port: 8899,
	}
	pool.usedChan = make(chan struct{}, pool.maxConnCount)
	for i := 0; i < pool.maxConnCount; i++{
		pool.usedChan <- struct{}{}
	}

	return pool
}

func (p *myConnPool) takeToken(timeWait time.Duration) error{
	select {
	case <-p.usedChan: {
		return nil
	}
	case <-time.After(timeWait): {
		return fmt.Errorf("takeToken time out")
	}
	}
}

func (p *myConnPool) freeToken() {
	select {
	case p.usedChan<- struct{}{}:
	default:
		fmt.Printf("freeToken default some error")
	}
}

func (p *myConnPool) Acquire(timeWait time.Duration) (io.ReadWriteCloser, error) {

	if err := p.takeToken(timeWait); err != nil {
		return nil ,err
	}
	p.lk.Lock()
	if p.conns.Len() == 0 {
		p.lk.Unlock()
		conn, err := p.factoryConn("tcp", fmt.Sprintf("%s:%d", p.addr, p.port))
		if err != nil {
			p.freeToken()
			return nil, errors.New(fmt.Sprintf("dial failed:%+v\n", err))
		}
		myconn := myConn {
			conn: conn,
			lasUsedTime:time.Now(),
		}

		return interface{}(myconn).(io.ReadWriteCloser), nil
	}
	try := 0
	for {
		if p.conns.Len() == 0 {
			p.lk.Unlock()
			return nil, fmt.Errorf("no usable conn")
		}
		val := p.conns.Remove(p.conns.Front())
		conn := val.(myConn)
		if time.Now().Sub(conn.lasUsedTime) > 5 * time.Second {
			err := conn.conn.Close()
			fmt.Printf("close result: %+v\n", err)
			if try < 3 {
				try = try + 1
				continue
			} else {
				p.lk.Unlock()
				return nil, fmt.Errorf("no usable conn and pop try 3times")
			}
		}
		p.lk.Unlock()
		return val.(io.ReadWriteCloser), nil
	}
}

func (p *myConnPool) Release(rwCloser io.ReadWriteCloser, pooled bool) error {
	p.freeToken()
	p.lk.Lock()
	defer p.lk.Unlock()

	if pooled {
		conn := rwCloser.(myConn)
		conn.lasUsedTime = time.Now()
		p.conns.PushBack(conn)
		return nil
	}
	return rwCloser.Close()
}