package main

import (
	"personHandler"
)

func main() {
	/*
		array1 := []string{"ruima","joz","1","yuanbao","python","c++","java"}
		slice1 := array1[:2]
		//slice1 = slice1[:cap(slice1)]
		//slice1[4] = "c#"
		//slice1 = append(slice1,"slice1","slice2","slice3","slice4")
		slice3 := append(slice1,"c#")
		slice3[0] = "new ruima"
		fmt.Println("hello ruima...!!!!!!!!!####!!!!xxx")
		fmt.Println("myMap len : ",len(array1), "cap :" ,cap(array1), "slice1 cap:" , cap(slice1), "slice1 len :" ,len(slice1))
		fmt.Println("array1 index 4:",array1[0])
		fmt.Println("slice1 index 4:",slice3[0])
		fmt.Println("slice3 cap:", cap(slice3),"slice3 len",len(slice3))
		var slice4 []string =  nil
		slice5 := append(slice4,"xx")
		fmt.Println("cap :", cap(slice5))*/

	handler := personHandler.GetPersonHandler()
	origs := make(chan personHandler.Person, 100)
	personHandler.FecthPerson(origs)
	dest := handler.Batch(origs)
	sign := personHandler.SavePerson(dest)
	<- sign
}
