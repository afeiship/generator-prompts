package main

import "fmt"

// https://www.cnblogs.com/xxnn/p/10875693.html

type Person struct {
	Name string
	Age  int
}

// func (r *Rect) Area() int
func (p *Person) Print() {
	fmt.Printf("hello %v \n", p.Name)
	fmt.Printf("you can live %v years.\n", p.Age)
}

func main() {
	var p1 = Person{
		Name: "feizheng",
		Age:  100,
	}

	p1.Print()
}
