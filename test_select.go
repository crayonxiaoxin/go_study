// package main

// import (
// 	"fmt"
// 	"time"
// )

// var chanInt = make(chan int, 0)
// var chanStr = make(chan string)

// func main() {
// 	go func() {
// 		chanInt <- 100
// 		chanStr <- "hello"
// 		defer close(chanInt)
// 		defer close(chanStr)
// 	}()

// 	// 死循环，一直读：
// 	// 1. 如果没有 defer，会走 default
// 	// 2. 如果 defer default都没有，会死锁
// 	// 3. 如果有 defer ，读完之后再读是 channel 的默认值
// 	for {
// 		select {
// 		case r := <-chanInt:
// 			fmt.Printf("r: %v\n", r)
// 		case r := <-chanStr:
// 			fmt.Printf("r: %v\n", r)
// 		default:
// 			fmt.Println("default....")
// 		}

// 		time.Sleep(time.Second)
// 	}

// }
