package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Enter function main.")
	go caller1()
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Exit function main.")
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
	fmt.Println("Exit function caller2.")
}
