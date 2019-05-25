package main

import (
	"fmt"
	"basic"
	"sort"
	liba "go-binary"

)

func generateKey(x interface{}) string {
	xa := interface{}(x).([]int)
	xac := make([]int, len(xa))
	copy(xac, xa)
	sort.Ints(xac)
	return fmt.Sprintf("%v", xac)
}

func compare(i, j interface{}) int {
	ia := interface{}(i).([]int)
	ja := interface{}(j).([]int)
	sort.Ints(ia)
	sort.Ints(ja)
	il := len(ia)
	jl := len(ja)
	result := 0
	if il < jl {
		result = -1
	} else if il > jl {
		result = 1
	} else {
		for i, iv := range ia {
			jv := ja[i]
			if iv != jv {
				if iv < jv {
					result = -1
				} else if iv > jv {
					result = 1
				}
				break
			}
		}
	}
	return result
}

func main() {

	fmt.Println("hello ruima...!!!!!!!!!####!!!!xxx")
	matrixSet := basic.SimpleSet{KeyGenerator: generateKey, Comparator: compare}

	matrix := make([][]int, 10)

	for _, v := range matrix {
		matrixSet.Add(interface{}(v))
	}
	fmt.Println(liba.Hello("haha"))
}
