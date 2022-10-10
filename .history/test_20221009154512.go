package main

import "fmt"

func main1() {
	// var name string
	// var age int
	// var email string

	var (
		name  string
		age   int
		email string
	)
	name, age, email = "批量", 1, "初始化"
	fmt.Printf("name: %v\n", name)
	fmt.Printf("name: %v\n", age)
	fmt.Printf("name: %v\n", email)

}
