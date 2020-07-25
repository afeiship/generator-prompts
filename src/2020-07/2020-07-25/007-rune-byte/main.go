package main

import "fmt"

func main() {
	var str = "abc"
	var a2 = "中国人abc"
	fmt.Println(str, len(str), str[1])
	fmt.Println(a2, len(a2))

	// for i := 0; i < len(a2); i++ {
	// 	fmt.Println(a2[i])
	// }

	for _, item := range a2 {
		fmt.Println(item)
	}
}
