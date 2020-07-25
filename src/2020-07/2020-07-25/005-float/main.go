package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a float32 = 3.5
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%v, %.2f", a, a)
}
