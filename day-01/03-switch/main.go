package main

import "fmt"

func main() {
	fmt.Println(grade(62))
}

func grade(score int) (g string) {
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "G"
		fallthrough
	case score < 100:
		g = "H"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return
}
