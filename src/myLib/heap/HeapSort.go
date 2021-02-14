package heap

import (
	"errors"
	"fmt"
)

type Elem uint32


var (
	NoLeftChild = errors.New("no left child")
	NoRightChild = errors.New("no right child")
)


type Heap struct {
	size uint32
	els []Elem
	height uint32
}

func (heep *Heap) leftChild(node uint32) (uint32, error){
	lc := uint32(node << 1)
	if lc > heep.size {
		return 0, NoLeftChild
	}
	return lc , nil
}

func (heep *Heap) rightChild(node uint32) (uint32, error) {
	lc ,err := heep.leftChild(node)
	if err != nil {
		return 0, err
	}
	rc := lc + 1
	if rc > heep.size {
		return 0, NoRightChild
	}
	return rc , nil
}


func (heep* Heap) Swap(i, j uint32) {
	heep.els[i],heep.els[j] = heep.els[j],heep.els[i]
}

func (heep* Heap) Compare(i, j uint32) bool{
	if heep.els[i] > heep.els[j] {
		return true
	}
	return false
}

func (heep* Heap) SelectChild(adjustNode uint32) (uint32, error){
	var swapNode uint32
	lc, err := heep.leftChild(uint32(adjustNode))
	if err != nil { //不可能没有左子树
		fmt.Printf("error :%s,adjustNode:%d\n", err, adjustNode)
		return 0, err
	}
	rc, err := heep.rightChild(uint32(adjustNode))
	if err != nil {
		swapNode = lc //没有右子树，直接拿左子树比较
		fmt.Printf("select child node:%d, adjustNode:%d\n", swapNode,adjustNode)
		return swapNode, nil
	}
	if heep.Compare(lc, rc) {
		swapNode = lc
	} else {
		swapNode = rc
	}
	fmt.Printf("select child node:%d, adjustNode:%d\n", swapNode,adjustNode)
	return swapNode,nil
}

func (heep *Heap) Adjust() {
	adjustLevel := heep.height - 1
	for i := adjustLevel; i >= 1; i-- {
		rightNode := (1 << i) - 1
		leftNode := 1 << (i - 1)
		fmt.Printf("adjustLevel:%d, lefttNode:%d, rightNode:%d\n",
			i, leftNode, rightNode)
		for adjustNode := rightNode; adjustNode >= leftNode; adjustNode-- {
			tmpadjustNode := uint32(adjustNode)
			tmpi := i
			for ;tmpi < heep.height; tmpi++ {
				swapNode, err := heep.SelectChild(tmpadjustNode)
				if err != nil {
					break
				}
				if heep.Compare(tmpadjustNode, swapNode) { //说明已经是大根堆了
					break
				} else {
					heep.Swap(tmpadjustNode, swapNode)
					tmpadjustNode = swapNode
				}
			}
		}
	}
}


func NewHeap(size uint32) *Heap{
	//els := make([]Elem, 1, size + 1)

	els := []Elem{0,1,4,3,2,5}

	height := uint32(0)
	tmpSize := size
	for ;tmpSize != 0; tmpSize = tmpSize >> 1 {
		height = height + 1
	}

	return &Heap {
		size: size,
		els: els,
		height:height,
	}
}


/*
func main() {
	heep := NewHeap(5)
	fmt.Printf("heep:%+v\n", heep)

	heep.Adjust()

	fmt.Printf("heep:%+v\n", heep)

}*/
