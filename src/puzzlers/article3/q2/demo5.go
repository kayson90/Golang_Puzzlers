package main

//路径是从$GOPAH和$GOROOT进行搜索的
import (
	"flag"
	"puzzlers/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
