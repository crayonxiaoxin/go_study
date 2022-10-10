// package main

// import "fmt"

// func main() {
// 	// 匿名函数
// 	max := func(a int, b int) int {
// 		if a > b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}
// 	r := max(1, 2)
// 	fmt.Printf("r: %v\n", r)

// 	// 自己调用自己
// 	r2 := func(a int, b int) int {
// 		if a > b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}(2, 3)
// 	fmt.Printf("r2: %v\n", r2)
// }
