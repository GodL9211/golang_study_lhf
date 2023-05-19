package main

import (
	"fmt"
	"time"
)

// 有缓存的channel
// 当channel已经满，再向channel写数据，会阻塞
// 当channel为空时，从channel取数据会阻塞
func main() {
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子协程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子协程正在运行，发送的元素=", i, " len(c)=", len(c), " cap(c)=", cap(c))
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	time.Sleep(time.Duration(1) * time.Second) // 保证defer和最后一次写c channel后的打印能正常打印，避免主线程直接退出，子线程来不及打印即退出
	fmt.Println("main 结束")
}
