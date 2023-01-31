package main

import "io/ioutil"

import "fmt"

func main() {
	const filename string = "F:\\study\\golang_study_lhf\\day-01\\02-if\\abc.txt"
	contents, err := ioutil.ReadFile(filename) //:= 的意思是声明变量并赋值
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
