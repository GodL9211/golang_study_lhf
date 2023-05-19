package main

import "fmt"

// 关闭channel
// 向一个已经关闭的channel发送数据会panic
// 关闭channel后，可以继续从channel接收数据
// nil channel, 接发都会阻塞
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 由于没有关闭通道，当子协程发送完数据后，主函数中的接收操作 <-c 将一直阻塞，导致所有的 Goroutine 都进入休眠状态，最终触发了死锁错误。
		//close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println("data is : ", data)
		} else {
			break
		}
	}

	fmt.Println("main finished..")
}
