package main

import "fmt"

// 空接口，万能类型

func MyFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)
}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}
	MyFunc(book)

	MyFunc(10)
}
