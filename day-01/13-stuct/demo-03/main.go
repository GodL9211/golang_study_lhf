package main

// 继承（组合）
import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat...")
}

func (this *Human) Walk() {
	fmt.Println("Human walk...")
}

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法
	level int
}

// 子类的新方法
func (s *SuperMan) fly() {
	fmt.Println("superman fly...")
}

func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.fly()
	s.Eat()
}
