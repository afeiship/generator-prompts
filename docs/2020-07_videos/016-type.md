# type
- https://www.bilibili.com/video/BV14T4y1g7h9?p=26
> 可以自定义类型

```go
type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

func main() {
	var c calc = add
	fmt.Printf("%T", c)
}
```
