package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a float32 = 3.5
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%v, %.4f \n", a, a)

	var f1 = 0.1
	var f2 = 0.2
	var sum = f1 + f2
	fmt.Printf("%v \n", sum)
}
