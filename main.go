package main

import (
	"fmt"
	"github.com/maruimm/myGoLearning/src/myLib/heap"
)

func main() {
	h := heap.NewHeap([]heap.Elem{1,2,35,5,4,0})
	fmt.Printf("befor adjust %+v\n",h )
	h.Adjust()
	fmt.Printf("after adjust %+v\n",h )
	h.InsertNode(heap.Elem(77))
	h.InsertNode(heap.Elem(77))
	h.InsertNode(heap.Elem(6))
	fmt.Printf("after insert %+v\n",h )
	var p heap.Elem
	for {
		err := h.PopRootNode(&p)
		if err != nil {
			fmt.Printf("pop err:%+v\n", err)
			break
		}
		fmt.Printf("pop elem:%+v, heap:%+v\n", p, h)
	}
}
