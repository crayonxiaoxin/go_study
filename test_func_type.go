// package main

// import "fmt"

// func sum(a int, b int) int {
// 	return a + b
// }

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func main() {
// 	type total func(int, int) int // 函数的类型定义
// 	var a total = sum             // 函数变量
// 	fmt.Printf("a(1, 2): %v\n", a(1, 2))

// 	var b total = max
// 	fmt.Printf("b(1, 2): %v\n", b(1, 2))
// }
