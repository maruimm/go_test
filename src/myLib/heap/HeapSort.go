package heap

import (
	"errors"
	"fmt"
)

type Elem uint32


var (
	NoLeftChild = errors.New("no left child")
	NoRightChild = errors.New("no right child")
	NoNode = errors.New("no node in heap")
)


type Heap struct {
	size uint32
	els []Elem
	height uint32
}

func (heap *Heap) leftChild(node uint32) (uint32, error){
	lc := uint32(node << 1)
	if lc > heap.size {
		return 0, NoLeftChild
	}
	return lc , nil
}

func (heap *Heap) rightChild(node uint32) (uint32, error) {
	lc ,err := heap.leftChild(node)
	if err != nil {
		return 0, err
	}
	rc := lc + 1
	if rc > heap.size {
		return 0, NoRightChild
	}
	return rc , nil
}


func (heap * Heap) Swap(i, j uint32) {
	heap.els[i], heap.els[j] = heap.els[j], heap.els[i]
}

func (heap * Heap) Compare(i, j uint32) bool{
	if heap.els[i] > heap.els[j] {
		return true
	}
	return false
}

func (heap * Heap) SelectChild(adjustNode uint32) (uint32, error){
	var swapNode uint32
	lc, err := heap.leftChild(uint32(adjustNode))
	if err != nil { //不可能没有左子树
		fmt.Printf("error :%s,adjustNode:%d\n", err, adjustNode)
		return 0, err
	}
	rc, err := heap.rightChild(uint32(adjustNode))
	if err != nil {
		swapNode = lc //没有右子树，直接拿左子树比较
		fmt.Printf("select child node:%d, adjustNode:%d\n", swapNode,adjustNode)
		return swapNode, nil
	}
	if heap.Compare(lc, rc) {
		swapNode = lc
	} else {
		swapNode = rc
	}
	fmt.Printf("select child node:%d, adjustNode:%d\n", swapNode,adjustNode)
	return swapNode,nil
}

func (heap *Heap) Adjust() {
	adjustLevel := heap.height - 1
	for i := adjustLevel; i >= 1; i-- {
		rightNode := (1 << i) - 1
		leftNode := 1 << (i - 1)
		fmt.Printf("adjustLevel:%d, lefttNode:%d, rightNode:%d\n",
			i, leftNode, rightNode)
		for adjustNode := rightNode; adjustNode >= leftNode; adjustNode-- {
			tmpadjustNode := uint32(adjustNode)
			tmpi := i
			for ;tmpi < heap.height; tmpi++ {
				swapNode, err := heap.SelectChild(tmpadjustNode)
				if err != nil {
					break
				}
				if heap.Compare(tmpadjustNode, swapNode) { //说明已经是大根堆了
					break
				} else {
					heap.Swap(tmpadjustNode, swapNode)
					tmpadjustNode = swapNode
				}
			}
		}
	}
}

func (heap *Heap) InsertNode(val Elem) {
	heap.els = append(heap.els, val)
	heap.height = heap.calHeight()
	heap.size = uint32(len(heap.els)) - 1

	newNodeIdx := heap.size
	for parentIndex := newNodeIdx >> 1; parentIndex != 0;  {
		if heap.Compare(newNodeIdx, parentIndex) {
			heap.Swap(parentIndex, newNodeIdx)
			newNodeIdx = parentIndex
			parentIndex = parentIndex >> 1
		} else {
			break
		}
	}
}

func (heap *Heap) PopRootNode(p *Elem) error {
	if heap.size < 1 {
		return NoNode
	}
	*p = heap.els[1]
	heap.Swap(1, heap.size)
	heap.els = heap.els[:heap.size]
	heap.size = heap.size - 1
	heap.height = heap.calHeight()
	fmt.Printf("after swap:%+v\n", heap)

	parentNodeIdx := uint32(1)
	for h := uint32(1); h < heap.height; h++ {
		selectedChild,err := heap.SelectChild(parentNodeIdx)
		if err != nil {
			continue
		}
		if heap.Compare(selectedChild, parentNodeIdx) {
			heap.Swap(selectedChild, parentNodeIdx)
			parentNodeIdx = selectedChild
		} else {
			break
		}
	}

	return nil
}


func (heap* Heap) calHeight() uint32 {
	size := heap.size
	height := uint32(0)
	for ; size != 0; size = size >> 1 {
		height = height + 1
	}
	return height
}

func NewHeap(val []Elem) *Heap{
	els := make([]Elem, 1)
	els = append(els, val...)
	size := uint32(len(val))

	heap := Heap {
		size: size,
		els: els,
	}
	height := heap.calHeight()
	heap.height = height
	return &heap
}

