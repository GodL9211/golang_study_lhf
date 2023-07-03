package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title string
	Code  string
	Price uint
}

func main() {
	operateSqlite()
}

func operateSqlite() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 插入内容
	//db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	//db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

	// 读取内容
	var product Product
	db.First(&product, 2) // find product with integer primary key
	fmt.Println(product)
	fmt.Println("-----------------")

	product = Product{}                  // 重新赋值
	db.Find(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)

	// 查询多条记录
	var productList []Product
	db.Find(&productList, "code = ?", "D42")
	fmt.Println(productList)

	//// 更新操作： 更新单个字段
	//db.Model(&product).Update("Price", 2000)
	//
	//// 更新操作： 更新多个字段
	//db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})
	//
	//// 删除操作：
	//db.Delete(&product, 1)
}
