package main

import (
	"fmt"
	"time"
)

func main() {
	var timeObj = time.Now()
	var str = timeObj.Format("2006-01-02 03:04:05")
	var str2 = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str)
	fmt.Println(str2)
}
