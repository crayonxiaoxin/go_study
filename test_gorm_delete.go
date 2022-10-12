// package main

// import (
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

// func delete() {
// 	var user User
// 	db.First(&user)
// 	db.Delete(&user) // 软删除，并非真正删除，只是赋值给 deleted_at
// }

// func delete2() {
// 	db.Delete(&User{}, 1)
// }

// func main() {
// 	// delete()
// 	delete2()
// }
