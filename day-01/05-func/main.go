package main

import "fmt"

func main() {
	a, b := div(13, 3)
	fmt.Println(a, b)

	c := sumArgs(1, 2, 3, 4, 5, 6)
	fmt.Println(c)
}

func sumArgs(values ...int) interface{} {
	sums := 0
	for i := range values {
		sums += values[i]
	}
	return sums
}

func div(a int, b int) (q int, r int) {
	q = a / b
	r = a % b
	return
}
