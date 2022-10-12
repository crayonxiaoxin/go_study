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
// 	Active   bool
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

// func update() {
// 	var u User
// 	db.Last(&u)
// 	fmt.Printf("u: %v\n", u)

// 	u.Active = !u.Active
// 	db.Save(&u) // 更新

// 	db.First(&u, u.ID)
// 	fmt.Printf("u: %v\n", u)
// }

// func update2() {
// 	db.Model(&User{}).Where("active = ?", true).Update("Name", "Hello Man")
// }

// func update3() {
// 	var u User
// 	db.First(&u)
// 	db.Model(&u).Updates(User{Name: "nihao", Age: 18, Active: true})
// }

// func main() {
// 	createTable()
// 	// update()
// 	// update2()
// 	update3()
// }
