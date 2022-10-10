// package main

// import "fmt"

// // defer 延迟处理
// // 先进后出 （先定义的后执行）
// func main() {
// 	fmt.Println("START")       // 1
// 	defer fmt.Println("step1") // 5
// 	defer fmt.Println("step2") // 4
// 	defer fmt.Println("step3") // 3
// 	fmt.Println("END")         // 2
// }
