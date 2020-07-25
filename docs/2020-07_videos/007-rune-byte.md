# rune/byte
- https://www.bilibili.com/video/BV14T4y1g7h9?p=9


## 字符
```go
var ch1 = 'a'
var ch2 = '中'
```

## 字符的类型
- uint8
- rune

## 字节
1. 一个字母是 1 个字节
2. 一个汉字，默认使用 utf-8，所以3个字节
3. 如果是 gbk 就只占2个字节
4. 字符串的空间用 len(str)


```go
var str = "abc"
fmt.Println(str, len(str), str[1])
// len 取得的就是字节数

// 结论： 这里和 javsscript 不同
abc 3 98
中国人abc 12

// 这种有汉字，输出会有问题
for i := 0; i < len(a2); i++ {
  fmt.Println(a2[i])
}
```

## for/range 循环

```go
for _, item := range a2 {
  fmt.Println(item)
}
// 20013
// 22269
// 20154
// 97
// 98
// 99
```

