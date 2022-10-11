// package main

// import (
// 	"fmt"
// 	"os"
// )

// func openCloseFile() {
// 	f, err := os.Open("a.txt") // 打开文件
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		f2, _ := os.Create("a.txt") // 创建文件 os.OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
// 		f = f2
// 	}
// 	fmt.Printf("f: %v\n", f.Name())
// 	f.Close() // 关闭文件
// }

// func openCloseFile2() {
// 	f, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755) // 打开，不存在则创建
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f: %v\n", f.Name())
// 		f.Close() // 关闭文件
// 	}
// }

// func readOps() {
// 	// f, _ := os.Open("a.txt") // 打开文件
// 	// for {
// 	// 	buf := make([]byte, 3) // 缓冲区，大小10
// 	// 	n, err := f.Read(buf)  // 每次读10byte
// 	// 	if err == io.EOF {     // 到达文件尾
// 	// 		break
// 	// 	}
// 	// 	fmt.Printf("n: %v\n", n)
// 	// 	fmt.Printf("string(buf): %v\n", string(buf))
// 	// }
// 	// f.Close()

// 	de, _ := os.ReadDir("a") // 读取目录 a
// 	for _, v := range de {
// 		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
// 		fmt.Printf("v.Name(): %v\n", v.Name())
// 	}
// }

// func main() {
// 	// openCloseFile()
// 	// openCloseFile2()
// 	readOps()
// }
