// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// func readString() {
// 	r := strings.NewReader("hello golang!")
// 	b, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("string(b): %v\n", string(b))
// 	}
// }

// func readFile() {
// 	f, _ := os.Open("a.txt") // File 已经实现了 Reader
// 	defer f.Close()
// 	b, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("string(b): %v\n", string(b))
// 	}
// }

// func readFile2() {
// 	b, _ := ioutil.ReadFile("a.txt")
// 	fmt.Printf("string(b): %v\n", string(b))
// }

// func writeFile() {
// 	err := ioutil.WriteFile("a.txt", []byte("kotlin"), 0755)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// func readDir() {
// 	fi, _ := ioutil.ReadDir(".") // 读取当前目录
// 	for _, v := range fi {
// 		fmt.Printf("v: %v\n", v.Name())
// 	}
// }

// func main() {
// 	// readString()
// 	// readFile()
// 	// readDir()
// 	// readFile2()
// 	writeFile()
// }
