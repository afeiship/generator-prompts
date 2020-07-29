// https://github.com/imdario/mergo
package main

import "fmt"

type title struct {
	Title    string
	SubTitle string
}

type desc struct {
	Remark      string
	Description string
}

type Article struct {
	title
	desc
}

func main() {
	var a1 = Article{
		title: title{
			Title:    "title",
			SubTitle: "sub title",
		},
	}

	fmt.Printf("%v", a1)
}
