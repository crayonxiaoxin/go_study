// package main

// import "fmt"

// var c = make(chan int)

// func main() {

// 	// // 1. normal （发送2次，读2次）
// 	// go func() {
// 	// 	for i := 0; i < 2; i++ {
// 	// 		c <- i
// 	// 	}
// 	// }()

// 	// r := <-c
// 	// fmt.Printf("r: %v\n", r)
// 	// r = <-c
// 	// fmt.Printf("r: %v\n", r)

// 	// // 2. deadlock! 死锁（发送2次，读3次）
// 	// go func() {
// 	// 	for i := 0; i < 2; i++ {
// 	// 		c <- i
// 	// 	}
// 	// }()

// 	// r := <-c
// 	// fmt.Printf("r: %v\n", r)
// 	// r = <-c
// 	// fmt.Printf("r: %v\n", r)
// 	// r = <-c
// 	// fmt.Printf("r: %v\n", r)

// 	// 3. 关闭channel （发送2次，读3次）
// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			c <- i
// 		}
// 		close(c) // 防止死锁
// 	}()

// 	r := <-c
// 	fmt.Printf("r: %v\n", r) // 0
// 	r = <-c
// 	fmt.Printf("r: %v\n", r) // 1
// 	r = <-c
// 	fmt.Printf("r: %v\n", r) // 0（第3次没有，返回默认值0）

// }
