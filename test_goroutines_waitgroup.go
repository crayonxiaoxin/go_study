// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wp sync.WaitGroup

// func showMsg(i int) {
// 	// defer wp.Add(-1)
// 	defer wp.Done() // goroutine 结束-1
// 	fmt.Printf("i: %v\n", i)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wp.Add(1) // goroutine 启动+1
// 		go showMsg(i)
// 	}

// 	wp.Wait() // 等待所有 goroutine 结束

// 	// 主协程
// 	fmt.Println("--- End ---")
// }
