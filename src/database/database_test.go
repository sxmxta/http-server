package database

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Title string
	Code  string
	Price uint
}

func TestSql(t *testing.T) {
	// Migrate the schema
	db := GetDB()
	db.AutoMigrate(&Product{})

	// 插入内容
	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

	// 读取内容
	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Printf("%+v \n", product)

	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Printf("%+v \n", product)

	var products []Product
	db.Find(&products)
	fmt.Printf("%+v \n", products)



	// 更新操作：更新单个字段
	db.Model(&product).Update("Price", 2000)
	fmt.Printf("%+v \n", product)

	// 更新操作：更新多个字段
	db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	fmt.Printf("%+v \n", product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})
	fmt.Printf("%+v \n", product)

	// 删除操作：
	tx := db.Delete(&product, 1)
	fmt.Printf("%+v \n", tx.Error)
}
