package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (h *Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.Level)
}

func main() {
	// 创建一个Hero
	hero := Hero{Name: "zhang3", Ad: 100}
	hero.Show()
}
