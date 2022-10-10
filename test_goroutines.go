// package main

// import (
// 	"fmt"
// 	"time"
// )

// func showMsg(msg string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("msg: %v\n", msg)
// 		time.Sleep(time.Microsecond * 100)
// 	}
// }

// // go 添加协程
// func main() {
// 	go showMsg("java")                  // 1 go 启动了一个协程
// 	go showMsg("golang")                // 2
// 	time.Sleep(time.Microsecond * 2000) // 等待协程完成
// 	fmt.Println("main end")             // 3 主函数退出，程序就结束了（没执行完的协程也就没用了）
// }
