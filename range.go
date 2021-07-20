package main

import (
	"fmt"
)

func main() {
	/*
	   array: a[1,2,3]




	*/
	array1 := [...]int{1, 3, 5}
	for _, num := range array1 {
		fmt.Println(num)
	}

}
