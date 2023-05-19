package main

import (
	"fmt"
	"time"
)

// 内核线程  用户线程  协程调度器  M:N
// 调度器策略： 复用线程  利用并行（cpu） 抢占策略  全局G队列（work stealing机制，从全局偷取）

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}

func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main Goroutine : i = %d\n", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}

// main goroutine退出，子goroutine立即退出
