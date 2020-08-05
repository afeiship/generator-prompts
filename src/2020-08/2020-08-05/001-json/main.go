package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	res := gjson.Parse(json)
	value := gjson.Get(json, "name")
	fmt.Println(res)
	fmt.Println(value.String())

	println("hello world")
}
