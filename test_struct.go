// package main

// import "fmt"

// // 结构体
// type Persion struct {
// 	id    int
// 	name  string
// 	age   int
// 	email string
// }

// type Customer struct {
// 	id, age     int
// 	name, email string
// }

// func main() {
// 	p := Persion{} // {0  0 }
// 	fmt.Printf("p: %v\n", p)

// 	xiaoming := Persion{
// 		id:    1,
// 		name:  "xiaoming",
// 		age:   18,
// 		email: "aaa@bbb.com",
// 	}
// 	xiaoming.age = 20
// 	xiaoming.email = "info@rmb.ee"
// 	fmt.Printf("xiaoming: %v\n", xiaoming)

// 	// 匿名结构体
// 	var tom struct {
// 		id, age int
// 		name    string
// 	}
// 	tom.age = 25
// 	tom.id = 2
// 	tom.name = "Tom"
// 	fmt.Printf("tom: %v\n", tom)
// }
