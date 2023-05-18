package main

import (
	"fmt"
	"reflect"
)

// 使用反射解析结构体标签

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagString, " doc: ", tagDoc)
	}
}

func main() {
	var re resume
	re = resume{"haiGe", "male"}
	findTag(&re)
}
