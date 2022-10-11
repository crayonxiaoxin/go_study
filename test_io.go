// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// func testCopy() {
// 	r := strings.NewReader("hello golang")
// 	written, err := io.Copy(os.Stdout, r) // 输出到控制台
// 	if err != nil {
// 		log.Fatal(err)
// 		fmt.Printf("written: %v\n", written)
// 	}
// }

// func main() {
// 	// r := strings.NewReader("hello world")
// 	// buf := make([]byte, 16)
// 	// r.Read(buf)
// 	// fmt.Printf("string(buf): %v\n", string(buf))

// 	testCopy()
// }
