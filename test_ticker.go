// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Timer 只执行一次， Ticker 可以周期执行

// func main() {
// 	// ticker := time.NewTicker(time.Second)
// 	// counter := 1
// 	// for _ = range ticker.C {
// 	// 	fmt.Println("ticker...")
// 	// 	if counter >= 5 {
// 	// 		ticker.Stop()
// 	// 		break
// 	// 	}
// 	// 	counter++
// 	// }

// 	ticker := time.NewTicker(time.Second)
// 	chanInt := make(chan int)
// 	go func() { // 启动协程
// 		for _ = range ticker.C { // 遍历 ticker，每秒
// 			select { // 随机case，发送数据
// 			case chanInt <- 1:
// 			case chanInt <- 2:
// 			case chanInt <- 3:
// 			}
// 		}
// 	}()

// 	sum := 0
// 	for v := range chanInt { // 遍历 channel
// 		fmt.Printf("收到: %v\n", v)
// 		sum += v
// 		if sum >= 20 { // 退出条件
// 			break
// 		}
// 	}

// }
