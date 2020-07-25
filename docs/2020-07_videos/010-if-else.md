# if/else
- https://www.bilibili.com/video/BV14T4y1g7h9?p=13

```go
func main() {
	var a = 1

	if a > 2 {
		fmt.Println("> 2z")
	} else {
		fmt.Println("<= 2")
	}
}


// 局部变量
if age := 30; age > 20 {
  fmt.Println(age)
}

// 这里就不能使用了，上面用完即销毁
// fmt.Println(age)
```
