package main

// 声明变量的四种方式
import "fmt"

func main() {
	fmt.Println("hello world")

	var a int
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b int = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c = 101
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	d := 102
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	// 多变量声明
	var e, f int = 103, 104
	fmt.Println(e, f)
}
