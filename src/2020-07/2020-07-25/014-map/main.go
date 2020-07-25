package main

import "fmt"

func main() {
	var userinfo = map[string]string{
		"username": "feizheng",
		"age":      "20",
	}

	fmt.Println(userinfo)

	for k, v := range userinfo {
		fmt.Println(k, v)
	}
}
