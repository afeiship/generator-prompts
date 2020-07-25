# function
- https://www.bilibili.com/video/BV14T4y1g7h9?p=26
- https://www.bilibili.com/video/BV14T4y1g7h9?p=27

```go
func add(x int, y int) int{
  return x + y;
}

// 表示，第一个参数与后面的参数类型一致，可以省略
func add(x, y int) int {
	return x + y
}


// 可变参数
func add(x...int){
  // x: [1,2,3,]
}
```
