// package main

// import (
// 	"fmt"
// )

// type Animal struct {
// 	name string
// 	age  int
// }

// func (a Animal) eat() {
// 	fmt.Printf("%v eat------\n", a.name)
// }

// func (a Animal) sleep() {
// 	fmt.Printf("%v sleep------\n", a.name)
// }

// type Dog struct {
// 	a     Animal
// 	color string
// }

// type Cat struct {
// 	Animal // 可以理解为继承
// 	height int
// }

// func main() {
// 	dog := Dog{
// 		a:     Animal{name: "dog", age: 18},
// 		color: "red",
// 	}
// 	dog.a.eat()
// 	dog.a.sleep()

// 	cat := Cat{
// 		Animal{name: "dog", age: 18},
// 		32,
// 	}
// 	cat.eat()
// 	cat.sleep()
// }
