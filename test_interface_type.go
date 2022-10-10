// package main

// import "fmt"

// type Pet interface {
// 	eat()
// }

// type Dog struct {
// 	name string
// }

// func (dog Dog) eat() {
// 	fmt.Printf("%v eat...\n", dog.name)
// }

// type Cat struct {
// 	name string
// }

// func (cat Cat) eat() {
// 	fmt.Printf("%v eat...\n", cat.name)
// }

// // 多态
// func main() {
// 	var pet Pet
// 	pet = Dog{name: "dog"}
// 	pet.eat()
// 	pet = Cat{name: "cat"}
// 	pet.eat()
// }
