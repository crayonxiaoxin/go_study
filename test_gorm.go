// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // 模型
// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// // 创建表
// func create(db *gorm.DB) {
// 	db.AutoMigrate(&Product{})
// }

// // 插入数据
// func insert(db *gorm.DB) {
// 	product := Product{Code: "1001", Price: 100}
// 	db.Create(&product)
// }

// // 查询数据
// func query(db *gorm.DB) {
// 	var product Product
// 	db.First(&product, 1) // 根据 primary key
// 	fmt.Printf("product: %v\n", product)

// 	db.First(&product, "code = ?", 1001) // 根据条件
// 	fmt.Printf("product: %v\n", product)
// }

// // 更新数据
// func update(db *gorm.DB) {
// 	var product Product
// 	db.First(&product, 1) // 根据 primary key
// 	db.Model(&product).Update("Price", 1000)

// 	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
// 	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "code": "F43"})
// }

// // 删除数据
// func delete(db *gorm.DB) {
// 	var product Product
// 	db.First(&product, 1) // 根据 primary key
// 	db.Delete(&product)
// }

// func main() {
// 	dsn := "root:New4you!@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	create(db)

// 	// insert(db)

// 	query(db)

// 	// update(db)

// 	// delete(db)
// }
