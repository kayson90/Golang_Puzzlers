package main

import "fmt"

type Printer func(contents string) (n int, err error)

//函数签名和Printer一样，所以是它的实现
func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	//由于是实现可以直接赋值
	p = printToStd
	p("something")
}
