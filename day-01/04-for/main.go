package main

import "fmt"

func main() {
	// 无限循环
	i := 1
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	fmt.Println("上面的是无限循环示例")
	// 条件循环
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}
	fmt.Println("上面的是条件循环示例")

	// 标准循环
	for k := 1; k <= 10; k++ {
		fmt.Println(k)
	}
	fmt.Println("上面的是标准循环的示例")
}
