package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 3: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1]) //从map中获取
}
