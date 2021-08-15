package main

import "fmt"

//接送操作符的使用
//箭头方向代码数据传输的方向
func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}
