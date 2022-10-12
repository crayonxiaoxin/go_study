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

// type User struct {
// 	id                 int
// 	username, password string
// }

// func queryOneRow(id int) {
// 	sql := "select * from user_tbl where id = ?"
// 	var u User
// 	err := db.QueryRow(sql, id).Scan(&u.id, &u.username, &u.password)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("u: %v\n", u)
// 	}
// }

// func queryMiltipleRows() {
// 	sql := "select * from user_tbl"
// 	rows, err := db.Query(sql)
// 	defer rows.Close()
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		for rows.Next() {
// 			var u User
// 			err2 := rows.Scan(&u.id, &u.username, &u.password)
// 			if err2 != nil {
// 				fmt.Printf("err2: %v\n", err2)
// 			} else {
// 				fmt.Printf("u: %v\n", u)
// 			}
// 		}
// 	}
// }

// func main() {
// 	err := initDB()
// 	if err != nil {
// 		fmt.Printf("连接数据库失败: %v\n", err)
// 	} else {
// 		fmt.Println("连接数据库成功！")
// 	}
// 	queryOneRow(2)
// 	fmt.Println("------------------")
// 	queryMiltipleRows()
// }
