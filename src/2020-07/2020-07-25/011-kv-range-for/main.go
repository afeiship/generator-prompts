package main

import "fmt"

func main() {
	var arr = []string{"php", "java", "golang"}
	for k, v := range arr {
		fmt.Println(k, v)
	}

	// 0 php
	// 1 java
	// 2 golang
}
