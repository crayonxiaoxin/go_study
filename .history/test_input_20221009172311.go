package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)
	fmt.Println("请输入 name, age: (空格分隔)")
	fmt.Scan(&name, &age) // 输入
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}
