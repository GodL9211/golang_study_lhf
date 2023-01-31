package main

import (
	"fmt"
	"math"
)

// 普通枚举类型
func consts() {
	const (
		filename string = "abc.txt"
		a, b     int    = 3, 4
	)

	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)

}

// 自增型枚举类型
func enus() {
	const (
		cpp = iota
		java
		python
		golang
	)

	fmt.Println(cpp, java, python, golang)

}

const (
	name1 = "boby"
	name2
	name3
	name4 = "jack"
	name5
)

func IotaPtintln() {
	fmt.Println(name1, name2, name3, name4, name5)
}

func main() {
	fmt.Println("hello world")
	consts()
	enus()
	IotaPtintln()
}
