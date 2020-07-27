package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cmd = exec.Command("go", "version")
	fmt.Printf("type: %T\n", cmd)
	var res, _ = cmd.Output()
	fmt.Printf("%s, %T\n", res, res)
	// exec.Cmd("go version")

	// fmt.Println("Hello, 世界")
	valuint8 := []byte{115, 101, 46, 49, 95, 49}
	valabc8 := []uint8{97, 98, 99}
	strout := string(valuint8)
	fmt.Println("out:", strout)

	// int/int8 to string
	fmt.Printf("out: %s \n", valabc8)
	fmt.Println("out: %v", string(97))
	fmt.Println("out: %v", string(valabc8))
}
