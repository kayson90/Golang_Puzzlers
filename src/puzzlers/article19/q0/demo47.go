package main

//对于panic是goroutine死掉还是整个进程死掉？
func main() {
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
}
