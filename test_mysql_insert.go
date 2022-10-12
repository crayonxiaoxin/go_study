// package main

// // https://pkg.go.dev/github.com/go-sql-driver/mysql

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// func initDB() (err error) {
// 	dsn := "root:New4you!@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
// 	db, err = sql.Open("mysql", dsn) // 这里不适用 := ,因为变量名已经定义过了
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Ping() // 尝试与数据库连接
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func insert(username string, password string) {
// 	sql := "insert into user_tbl(username,password) values(?,?)"
// 	r, err := db.Exec(sql, username, password)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		i, _ := r.LastInsertId()
// 		fmt.Printf("i: %v\n", i)
// 	}
// }

// func main() {
// 	err := initDB()
// 	if err != nil {
// 		fmt.Printf("连接数据库失败: %v\n", err)
// 	} else {
// 		fmt.Println("连接数据库成功！")
// 	}
// 	insert("zhangsan", "zs123")
// }
