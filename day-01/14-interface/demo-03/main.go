package main

// 变量内置pair结构

import "fmt"

func main() {
	var a string
	a = "haige"
	fmt.Println(a)

	var allType interface{}
	allType = a

	fmt.Println(allType)

	str, _ := allType.(string)
	fmt.Println(str)
}
