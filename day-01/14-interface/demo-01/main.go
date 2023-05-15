package main

import "fmt"

// 接口和多态

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Cat 具体的类
type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "cat"
}

// Dog 具体的类
type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("dog is sleep")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "dog"
}

func ShowAnimal(a AnimalIF) {
	a.Sleep()
}

func main() {
	var cat AnimalIF
	cat = &Cat{"green"}

	var dog AnimalIF
	dog = &Dog{"yello"}

	ShowAnimal(cat)
	ShowAnimal(dog)
}
