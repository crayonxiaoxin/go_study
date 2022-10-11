// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func test1() {
// 	// 强转
// 	var i int = 100
// 	var b byte = 10
// 	b = byte(i)
// 	fmt.Printf("b: %v\n", b)
// 	i = int(b)
// 	fmt.Printf("i: %v\n", i)

// 	var s string = "hello world"
// 	var b1 = []byte{104, 101, 108}
// 	fmt.Printf("string(b1): %v\n", string(b1))
// 	fmt.Printf("[]byte(s): %v\n", []byte(s))
// }

// func test2() {
// 	s := []byte("你好世界abc")
// 	r := bytes.Runes(s)
// 	fmt.Printf("len(s): %v\n", len(s))
// 	fmt.Printf("len(r): %v\n", len(r)) // 真实长度
// }

// func main() {
// 	// test1()
// 	test2()
// }
