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

// func select1() {
// 	var user User
// 	// SELECT * FROM users ORDER BY id LIMIT 1;
// 	db.First(&user)
// 	fmt.Printf("user: %v\n", user.ID)

// 	var user1 User
// 	db.Take(&user1)
// 	// SELECT * FROM users LIMIT 1;
// 	fmt.Printf("user: %v\n", user1.ID)

// 	var user2 User
// 	db.Last(&user2)
// 	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
// 	fmt.Printf("user: %v\n", user2.ID)
// }

// func select2() {
// 	// 查找单个id
// 	var user User
// 	// SELECT * FROM users ORDER BY id LIMIT 1;
// 	db.First(&user, 2)
// 	fmt.Printf("user: %v\n", user.ID)

// 	// 查找多个id
// 	var users []User
// 	db.Find(&users, []int{1, 2, 3})
// 	for _, v := range users {
// 		fmt.Printf("v: %v\n", v.ID)
// 	}
// }

// // 查找所有
// func select3() {
// 	var users []User
// 	d := db.Find(&users)
// 	fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)
// }

// func main() {
// 	// select1()
// 	// select2()
// 	select3()
// }
