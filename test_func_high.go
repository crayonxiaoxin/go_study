// package main

// import "fmt"

// func sayHello(name string) {
// 	fmt.Printf("Hello, %v\n", name)
// }

// // 函数作为参数
// func test(name string, f func(string)) {
// 	f(name)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// func sub(a int, b int) int {
// 	return a - b
// }

// // 函数作为返回值
// func calc(operator string) func(int, int) int {
// 	switch operator {
// 	case "+":
// 		return add
// 	default:
// 		return sub
// 	}
// }

// func main() {
// 	// test 第二个参数是函数
// 	test("world", sayHello)

// 	// calc("+") 返回值是一个函数
// 	addRes := calc("+")(2, 5)
// 	fmt.Printf("calc(\"+\")(2, 5): %v\n", addRes)
// 	subRes := calc("-")(5, 9)
// 	fmt.Printf("calc(\"-\")(5, 9): %v\n", subRes)
// }
