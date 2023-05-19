package main

import "fmt"

// 单流程下一个go只能监控一个channel的状态，select可以完成监控多个channel的状态

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 11; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacii(c, quit)

}
