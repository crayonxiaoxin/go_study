// package main

// import "fmt"

// type Persion struct {
// 	name string
// 	age  int
// }

// // 构造函数
// func NewPersion(name string, age int) (*Persion, error) {
// 	if name == "" {
// 		return nil, fmt.Errorf("name 不能为空")
// 	}
// 	if age < 0 {
// 		return nil, fmt.Errorf("age 不能小于0")
// 	}
// 	return &Persion{name: name, age: age}, nil
// }

// func main() {
// 	persion, error := NewPersion("Tom", -20)
// 	fmt.Printf("persion: %v\n", persion)
// 	fmt.Printf("error: %v\n", error)
// }
