// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func testLog() {
// 	log.Print("my log")
// 	log.Printf("my log %d", 10086)
// 	log.Println("--- my log ---")
// }

// func testPanic() {
// 	defer fmt.Println("panic 结束后执行...") // 会执行
// 	log.Print("my log")
// 	log.Panic("my panic") // 发生了异常
// 	fmt.Println("end...") // 不会执行
// }

// func testFatal() {
// 	defer fmt.Println("fatal 结束后执行...") // 不会执行
// 	log.Print("my log")
// 	log.Fatal("my fatal") // 发生了异常
// 	fmt.Println("end...") // 不会执行
// }

// func init() {
// 	// log 的配置
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 设置format
// 	log.SetPrefix("MyLog: ")                             // 设置前缀

// 	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		log.Fatal("日志文件错误")
// 	} else {
// 		log.SetOutput(f) // 设置输出方式，输出到文件
// 	}
// }

// func main() {
// 	// testLog()
// 	// testPanic()
// 	// testFatal()

// 	i := log.Flags()
// 	fmt.Printf("i: %v\n", i)
// 	log.Println("My Log")
// }
