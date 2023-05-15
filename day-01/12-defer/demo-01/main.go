package main

import "fmt"

func main() {
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	defer fmt.Println("main end3")

	fmt.Println("hello main")
}
