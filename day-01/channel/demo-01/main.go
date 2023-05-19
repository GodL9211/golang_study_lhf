package main

import (
	"fmt"
)

// 无缓冲channel

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		fmt.Println("goroutine 正在运行...")
		// 将666发送给c
		c <- 666
		fmt.Println("777")
	}()

	num := <-c

	fmt.Println("从管道中得到的数据为： ", num)

	fmt.Println("main goroutine 结束...")
}

// 无缓冲channel有同步机制（阻塞）, 如果没有num接收，子goroutine也会阻塞
