package main

import (
	"fmt"
	"github.com/maruimm/myGoLearning/src/myLib/heap"
)

func main() {
	h := heap.NewHeap(5)
	fmt.Printf("befor adjust %+v\n",h )
	h.Adjust()
	fmt.Printf("after adjust %+v\n",h )
}
