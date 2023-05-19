package main

import "fmt"

// channel和range
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c)
	}()

	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println("data is : ", data)
	//	} else {
	//		break
	//	}
	//}

	//使用range改写上面的for，简化
	for data := range c {
		fmt.Println("data is : ", data)
	}

	fmt.Println("main finished..")
}
