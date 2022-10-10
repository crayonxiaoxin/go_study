// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 创建 int 类型通道，只能传入 int 值
// var values = make(chan int) // 无缓冲

// func send() {
// 	rand.Seed(time.Now().UnixNano())
// 	value := rand.Intn(10)
// 	fmt.Printf("value send: %v\n", value)
// 	// time.Sleep(time.Second * 5)
// 	values <- value // 发送
// }

// func main() {
// 	defer close(values)
// 	go send()
// 	fmt.Println("wait......")
// 	value := <-values // 接收
// 	fmt.Printf("value receive: %v\n", value)
// 	fmt.Println("end.....")
// }
