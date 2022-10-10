// package main

// import "fmt"

// // 闭包
// func add() func(y int) int {
// 	var x int
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func calc(base int) (func(a int) int, func(b int) int) {
// 	add := func(a int) int {
// 		base += a
// 		return base
// 	}
// 	sub := func(b int) int {
// 		base -= b
// 		return base
// 	}
// 	return add, sub
// }

// func main() {
// 	f := add()
// 	r := f(10)
// 	fmt.Printf("r: %v\n", r) // x=0  y=10	10
// 	r = f(20)
// 	fmt.Printf("r: %v\n", r) // x=10 y=20	30
// 	r = f(30)
// 	fmt.Printf("r: %v\n", r) // x=30 y=30	60

// 	fmt.Println("----------------")

// 	add, sub := calc(100)
// 	r = add(100)
// 	fmt.Printf("r: %v\n", r) // 100 + 100 = 200
// 	r = sub(50)
// 	fmt.Printf("r: %v\n", r) // 200 - 150 = 150
// }
