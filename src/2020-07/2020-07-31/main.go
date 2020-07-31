package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Address
}

type Address struct {
	City string
}

func main() {
	p1 := Person{}
	p1.Name = "feizheng"
	p1.Age = 100

	// 两种形式均可以达到目的
	p1.City = "sh"
	p1.Address.City = "wh"

	fmt.Printf("%#v \n", p1)
}
