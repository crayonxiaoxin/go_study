package main

import "fmt"

func main() {
	var name string = "tom"
	var age int = 20
	b := true

	// %v 值   %T 类型
	fmt.Printf("name: %v\ntype: %T\n\n", name, name)
	fmt.Printf("age: %v\ntype: %T\n\n", age, age)
	fmt.Printf("b: %v\ntype: %T\n\n", b, b)

	a := [...]int{1, 2}
	var c = [...]int{1, 2}
	var d = [2]int{1, 2}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
}
