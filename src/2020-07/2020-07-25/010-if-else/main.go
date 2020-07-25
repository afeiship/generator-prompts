package main

import "fmt"

func main() {
	var a = 1

	if a > 2 {
		fmt.Println("> 2z")
	} else {
		fmt.Println("<= 2")
	}

	if age := 30; age > 20 {
		fmt.Println(age)
	}

	// 这里就不能使用了，上面用完即销毁
	// fmt.Println(age)
}
