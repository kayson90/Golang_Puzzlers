package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 示例1。
	// 只能发不能收的通道。
	var uselessChan = make(chan<- int, 1)
	// 只能收不能发的通道。
	var anotherUselessChan = make(<-chan int, 1)
	// 这里打印的是可以分别代表两个通道的指针的16进制表示。
	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan, anotherUselessChan)
	//这里理解起来有点绕，我们关注的是数据流的方向，上面是发送给通道，所以是发送通道
	//下面是数据从通道接收，所以是接收收通道

	// 示例2。双向通道转换为发送通道
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	// 示例4。
	intChan2 := getIntChan()
	//要注意for range操作通道时阻塞的情况
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// 示例5。这是一种类型转换
	a := GetIntChan(getIntChan)
	// 分别输出的是GetIntChan类型，接收通道类型
	fmt.Printf("%T,%T\n", a, a())
}

// 示例2。约束了函数的行为，只能往通道写数据
func SendInt(ch chan<- int) {
	//发送数据给通道，所以是发送通道
	ch <- rand.Intn(1000)
}

// 示例3。
type Notifier interface {
	SendInt(ch chan<- int)
}

// 示例4。返回结果是一个接收通道
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 示例5。自定义函数类型，函数的结果是一个接收通道
type GetIntChan func() <-chan int
