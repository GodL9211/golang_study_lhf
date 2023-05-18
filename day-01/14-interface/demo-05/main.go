package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriterBook()
}

type Book struct {
}

func (receiver Book) ReadBook() {
	fmt.Println("read a book")
}

func (receiver Book) WriterBook() {
	fmt.Println("writer a book")
}

func main() {
	// b: pair<type:Book, value:Book{}地址>
	b := &Book{}

	var r Reader
	// r: pair<type:, value:>
	r = b
	// r: pair<type:Book, value:Book{}地址>
	r.ReadBook()

	// w: pair<type:, value:>
	var w Writer
	// w: pair<type:Book, value:Book{}地址>
	w = r.(Writer) // 此处断言为什么会成功？因为w r 具体的type是一致的
	w.WriterBook()

}
