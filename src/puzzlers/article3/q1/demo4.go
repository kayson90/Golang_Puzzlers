package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	//如果只单独编译这个文件，是不行的，这点和java有区别
	//go run demo4.go demo4_lib.go
	//或者go run ./q1
	hello(name)
}
