package main

import "fmt"

func main() {
	var name string = "tom"
	var age int = 20
	b := true
	fmt.Printf("name: %v\n%T\n", name, name)
	fmt.Printf("age: %v\n%T\n", age, age)
	fmt.Printf("b: %v\n%T\n", b, b)
}
