package main

import "fmt"

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		"喜剧之王", 2000, 10, []string{"星爷", "zhangbozhi"},
	}
	fmt.Println(movie)
}
