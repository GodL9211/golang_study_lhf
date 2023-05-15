package main

import "fmt"

// Book 定义一个结构体
type Book struct {
	tittle string
	auth   string
}

func changeBook(b Book) {
	// 传递一个book的副本
	b.tittle = "python"
}

func changeBook2(b *Book) {
	// 传递一个book的指针
	b.tittle = "python"
}

func main() {
	var book1 Book
	book1.tittle = "Golang"
	book1.auth = "zhang3"
	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
