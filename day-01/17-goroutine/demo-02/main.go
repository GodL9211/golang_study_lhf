package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 退出当前协程
			//runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	// 协程函数待参数
	go func(a, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(time.Duration(1) * time.Second)
	}
}
