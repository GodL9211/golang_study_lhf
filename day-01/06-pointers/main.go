package main

import "fmt"

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	// 使用指针交换两数的值
	c := 5
	d := 15
	swap(&c, &d)
	fmt.Println(c, d)
}

// go语言只有值传递一种方式
func swap(c *int, d *int) {
	*c, *d = *d, *c
}
