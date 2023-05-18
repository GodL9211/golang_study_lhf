package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// tty: pair<type:*os.File, value:"./hello.json"文件描述符>
	tty, err := os.OpenFile("d:/hello.json", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	write, err := w.Write([]byte("hello"))
	if err != nil {
		return
	}
	fmt.Println("=======")
	fmt.Println(write)
}
