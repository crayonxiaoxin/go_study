// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	s := os.Environ() // 环境变量
// 	fmt.Printf("s: %v\n", s)

// 	s2 := os.Getenv("GOPATH") // 获取 GOPATH
// 	fmt.Printf("s2: %v\n", s2)

// 	s3, b := os.LookupEnv("ANDROID_HOME") // 获取某个环境变量，b表示成功或失败
// 	if b {
// 		fmt.Printf("s3: %v\n", s3)
// 	}

// }
