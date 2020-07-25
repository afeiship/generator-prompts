# defer
- https://www.bilibili.com/video/BV14T4y1g7h9?p=28
- 延时执行，被这个定义的语句



```go
fmt.Println("11")
defer func() {
  fmt.Println("aaa")
  fmt.Println("bbb")
}()
fmt.Println("22")
```
