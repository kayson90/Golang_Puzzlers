package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	//变量的重声明，必须要有一个新变量
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}
