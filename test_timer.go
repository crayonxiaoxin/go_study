// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// fmt.Printf("time.Now(): %v\n", time.Now()) // now
// 	// timer := time.NewTimer(time.Second * 2)
// 	// t1 := <-timer.C            // 阻塞的，直到指定时间到了
// 	// fmt.Printf("t1: %v\n", t1) // now +2s

// 	// fmt.Printf("time.Now(): %v\n", time.Now()) // now
// 	// timer2 := time.NewTimer(time.Second * 2)
// 	// <-timer2.C                                 // 阻塞的，直到指定时间到了
// 	// fmt.Printf("time.Now(): %v\n", time.Now()) // now +2s

// 	// <-time.After(time.Second * 2) // 等待2s
// 	// fmt.Println("----")

// 	// 停止 timer
// 	timer := time.NewTimer(time.Second)
// 	// timer.Reset(time.Second * 2) // 重置
// 	go func() {
// 		<-timer.C
// 		fmt.Println("func......")
// 	}()
// 	stop := timer.Stop()
// 	if stop {
// 		fmt.Println("stop timer......")
// 	}
// }
