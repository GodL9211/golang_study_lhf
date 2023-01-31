package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(a)
	c <- total // send total to c

}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:3], c)
	go sum(a[3:], c)
	go sum(a[:], c)

	x := <-c
	y := <-c
	z := <-c

	fmt.Println(x, y, z, x+y+z)

}
