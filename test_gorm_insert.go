// package main

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // 模型
// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      int
// 	Birthday time.Time
// }

// var db *gorm.DB

// func init() {
// 	dsn := "root:New4you!@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db = d
// }

// // 创建表
// func createTable() {
// 	db.AutoMigrate(&User{})
// }

// // 插入数据
// func insert1() {
// 	user := User{
// 		Name:     "Tom",
// 		Age:      20,
// 		Birthday: time.Now(),
// 	}
// 	result := db.Create(&user)

// 	fmt.Printf("d.RowsAffected: %v\n", result.RowsAffected)
// 	fmt.Printf("user.ID: %v\n", user.ID)
// }

// func insert2() {
// 	user := User{
// 		Name:     "Jerry",
// 		Age:      18,
// 		Birthday: time.Now(),
// 	}
// 	result := db.Select("Name", "Age", "CreatedAt").Create(&user) // 筛选只添加哪些字段

// 	fmt.Printf("d.RowsAffected: %v\n", result.RowsAffected)
// 	fmt.Printf("user.ID: %v\n", user.ID)
// }

// func insert3() {
// 	var users = []User{
// 		{Name: "Spider Man"},
// 		{Name: "Iron Man"},
// 		{Name: "Super Man"},
// 	}
// 	db.Create(&users)
// }

// func main() {
// 	createTable()
// 	// insert1()
// 	// insert2()
// 	insert3()
// }
