package main

import "fmt"

func main() {
	fmt.Println("11")
	defer func() {
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()
	fmt.Println("22")
}
