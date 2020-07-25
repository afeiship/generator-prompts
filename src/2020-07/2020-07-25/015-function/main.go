package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var a = 1
	var b = 2
	var c = add(a, b)

	fmt.Println(c)
}
