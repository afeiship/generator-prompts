package main

import "fmt"

func main() {
	// var username string
	var passowrd string = "2000"
	fmt.Printf("value: %v\n", passowrd)

	// 同时定义多个变量
	var (
		name string = "afei"
		age  int    = 100
		site string = "https://github.com/marketplace?type=actions"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(site)

	// short var
	github := "https://www.mipengine.org/"
	fmt.Println(github)
}
