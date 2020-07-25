# float
- https://www.bilibili.com/video/BV14T4y1g7h9?p=7


## float
- float32
- float64

## unsafe pakcag
> sizeOf 可以查看所占的内存空间

```go
	var a float32 = 3.5
	fmt.Println(a)
  fmt.Println(unsafe.Sizeof(a))
```

## format outputh
```go
fmt.Printf("%v, %.4f", a, a)
```

## decimal ieee754
> 同样的存在精度丢失的问题

```go
	var f1 = 0.1
	var f2 = 0.2
	var sum = f1 + f2
	fmt.Printf("%v \n", sum)
```
