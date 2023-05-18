package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (receiver User) Call() {
	fmt.Println("User is called")
	fmt.Printf("input type is : %v", receiver)
}

func main() {
	user := User{1, "haige", 18}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is : ", inputType)

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is : ", inputValue)

	// 通过type获取里面的字段
	// 获取interface的reflect.Type, 通过type得到NumField, 进行遍历
	// 得到每个field，数据类型
	// 通过filed有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v=%v\n", field.Name, field.Type, value)
	}

	// 获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}
