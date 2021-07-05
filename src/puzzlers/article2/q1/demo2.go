package main

import (
	"flag"
	"fmt"
)

var name string

//todo flag的使用不必一开始就掌握
func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
