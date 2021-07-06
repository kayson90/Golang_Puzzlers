package main

import "fmt"

//关于扩容，值得细细琢磨
func main() {
	// 示例1。
	s6 := make([]int, 0)
	fmt.Printf("The capacity of s6: %d\n", cap(s6))
	//从这里看到确实是2倍的关系
	for i := 1; i <= 5; i++ {
		s6 = append(s6, i)
		fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6))
	}
	fmt.Println()

	// 示例2。
	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
	//从这里看到确实是1.25的关系
	s7e1 := append(s7, make([]int, 200)...)
	fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
	fmt.Println()
	//todo 这种情况是比较复杂的，后面再研究一下
	// 示例3。
	s8 := make([]int, 10)
	fmt.Printf("The capacity of s8: %d\n", cap(s8))
	s8a := append(s8, make([]int, 11)...)
	fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))
	s8b := append(s8a, make([]int, 23)...)
	fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))
	s8c := append(s8b, make([]int, 45)...)
	fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))
}
