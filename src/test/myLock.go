package main

import (
	"fmt"
	"sync"
)


type Data struct {
	lk *sync.RWMutex
	counter int
}

func (p *Data) Incr() {
	p.lk.Lock()
	defer p.lk.Unlock()
	p.counter = p.counter + 1
}

func (p *Data) Get() int {

	p.lk.Lock()
	defer p.lk.Unlock()

	return p.counter
}


func main() {

	data := Data {
		lk: &sync.RWMutex{},
	}
	data.Incr()
	fmt.Printf("data:%+v, c:%d\n", data,data.Get())
}


