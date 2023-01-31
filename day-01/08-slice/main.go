package main

import "fmt"

type Person struct {
	name string
	age  int
}

func slice1() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	//[3:5]中的这个5 如果超过数组长度，那么是会报错的，也就是下面例子中的cap()方法算出来的长度，不能超过，超过就会报错。
	fmt.Println(s1, s2)
	fmt.Printf("s1 = %v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2)=%d, cap(s2)=%d \n", s2, len(s2), cap(s2))
}

func main() {

	// 从数组生成slice, 打印slice
	slice1()

	for i := 1; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	type List []interface{}
	var alist = make(List, 0, 3) // 使用make创建slice, 最佳实践是初始化时指定length和cap, 这里分别为0和3
	blist := make(List, 3, 3)

	alist = append(alist, 1) // 向slice添加元素
	alist = append(alist, "hello world")
	alist = append(alist, Person{name: "xiaoming", age: 18})
	alist = append(alist, blist...) // 将blist加入到alist中
	fmt.Println(alist)

	for _, value := range alist { // 切片遍历
		switch value.(type) {
		case int:
			fmt.Println("type of value is int")
		case string:
			fmt.Println("type of value is string")
		}
	}

}
