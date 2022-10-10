package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "tom"
	var age int = 20
	b := true

	// %v var   %T 类型
	fmt.Printf("name: %v\ntype: %T\n\n", name, name)
	fmt.Printf("age: %v\ntype: %T\n\n", age, age)
	fmt.Printf("b: %v\ntype: %T\n\n", b, b)

	a := [...]int{1, 2}
	var c = [...]int{1, 2}
	var d = [2]int{1, 2}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	if age >= 18 && b {
		fmt.Println("你已经成年了\n")
	} else {
		fmt.Println("你还未成年\n")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}

	haha := "0123456789"
	fmt.Printf("haha[2:6]: %v\n", haha[2:6])
	fmt.Printf("haha[2:6]: %v\n", haha[2:])
	fmt.Printf("haha[2:6]: %v\n", haha[:6])

	fmt.Printf("strings.Split(\"hello world\", \" \"): %v\n", strings.Split("hello world", " "))

	type WebSite struct {
		Name string
	}
	site := WebSite{Name: "tom"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
}
