package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "lian",
		"course":  "golang",
		"site":    "mhs",
		"quality": "notbad",
	}
	// 使用make创建map
	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)
	fmt.Println("Traversing map")
	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	// 删除
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
