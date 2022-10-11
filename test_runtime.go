// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	go show("java") // 子协程

// 	// 主协程
// 	for i := 0; i < 2; i++ {
// 		runtime.Gosched() // 先让给其他协程执行（注释掉这句会发现上面子协程不能执行完毕）
// 		fmt.Printf("i: %v\n", "golang")
// 	}
// 	fmt.Println("end...")

// 	core := runtime.NumCPU() // cpu 核心数
// 	fmt.Printf("runtime.NumCPU(): %v\n", core)
// }

// func show(msg string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%v: %v\n", i, msg)
// 		if i >= 5 {
// 			runtime.Goexit() // 退出协程
// 		}
// 	}
// }
