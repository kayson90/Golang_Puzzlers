package main

import "fmt"

//不要让接收方关闭chan，因为往关闭的通道发送是会panic
func main() {
	ch1 := make(chan int, 2)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		//ok用来检测chan是否已经关闭，检测可能有延时，因为关闭了的通道还有元素，依然显示是未关闭的
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
