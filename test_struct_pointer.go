// package main

// import "fmt"

// func main() {
// 	var name string
// 	name = "Tom"
// 	var p_name *string
// 	p_name = &name
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Printf("p_name: %v\n", p_name)
// 	fmt.Printf("p_name: %v\n", *p_name)

// 	type Person struct {
// 		id, age int
// 		name    string
// 	}
// 	var tom = Person{
// 		id:   1,
// 		age:  18,
// 		name: "Tom",
// 	}
// 	var p_tom *Person
// 	p_tom = &tom
// 	fmt.Printf("p_tom: %p\n", p_tom)
// 	fmt.Printf("p_tom: %v\n", *p_tom)

// 	var tom_new = new(Person) // tom_new 是指针类型
// 	tom_new.id = 101
// 	tom_new.age = 55
// 	tom_new.name = "Tom"
// 	fmt.Printf("tom_new: %v\n", tom_new)
// 	fmt.Printf("tom_new: %p\n", tom_new)
// 	fmt.Printf("tom_new: %v\n", *tom_new)

// }
