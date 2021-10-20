package main

import "fmt"

var channels = [3]chan int{
	nil,
	//创建一个无缓冲的通道
	make(chan int),
	nil,
}

var numbers = []int{1, 2, 3}

func main() {
	court := make(chan int)
	court <- 1
	a := <-court
	println(a)
	//select {
	//case getChan(0) <- getNumber(0):
	//	fmt.Println("The first candidate case is selected.")
	//case getChan(1) <- getNumber(1):
	//	fmt.Println("The second candidate case is selected.")
	//case getChan(2) <- getNumber(2):
	//	fmt.Println("The third candidate case is selected")
	//default:
	//	fmt.Println("No candidate case is selected!")
	//}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
