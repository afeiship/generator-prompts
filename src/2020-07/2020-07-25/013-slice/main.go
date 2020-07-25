package main

import "fmt"

func main() {
	// 不指定长度：slice
	var arr1 []int

	// 指定长度: arr
	var arr2 [5]string

	fmt.Printf("%v - %T \n", arr1, arr1)
	fmt.Printf("%v - %T \n", arr2, arr2)
}
