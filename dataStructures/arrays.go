package main

import (
	"fmt"
)

func main() {
	var arr = […]int{1,2,3,4}
	var i, v int
	fmt.Println(arr)

	for i = 0; i < len(arr); i++ {
		fmt.Println("Elements are: ", arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}