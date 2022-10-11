// package main

// import (
// 	"fmt"
// 	"os"
// )

// // 创建文件
// func createFile() {
// 	f, err := os.Create("a.txt") // 如果存在，会覆盖
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f: %v\n", f.Name())
// 	}
// }

// // 创建目录
// func makeDir() {
// 	// err := os.Mkdir("a", os.ModePerm)
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }
// 	err2 := os.MkdirAll("a/b/c", os.ModePerm)
// 	if err2 != nil {
// 		fmt.Printf("err2: %v\n", err2)
// 	}
// }

// // 删除目录或者文件
// func removeDirOrFile() {
// 	// err := os.Remove("a.txt")
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }

// 	err := os.RemoveAll("a")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// // 获取工作目录
// func wd() {
// 	dir, err := os.Getwd() // 获取工作目录
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("dir: %v\n", dir)
// 	}

// 	err2 := os.Chdir("/Users/xin/work") // 切换工作目录
// 	if err2 != nil {
// 		fmt.Printf("err2: %v\n", err2)
// 	} else {
// 		dir, _ := os.Getwd()
// 		fmt.Printf("dir: %v\n", dir)
// 	}

// 	s := os.TempDir() // 获取临时目录
// 	fmt.Printf("s: %v\n", s)
// }

// // 重命名
// func rename() {
// 	err := os.Rename("test.txt", "test2.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// // 读文件
// func readFile() {
// 	b, err := os.ReadFile("test2.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("b: %v\n", string(b[:]))
// 	}
// }

// // 写文件
// func writeFile() {
// 	s := "hello world"
// 	err := os.WriteFile("test2.txt", []byte(s), os.ModePerm)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// func main() {
// 	// createFile()
// 	// makeDir()
// 	// removeDirOrFile()
// 	// wd()
// 	// rename()
// 	// readFile()
// 	// writeFile()
// }
