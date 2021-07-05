package lib

//是internal的直接父级
import (
	"os"
	in "puzzlers/article3/q4/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
