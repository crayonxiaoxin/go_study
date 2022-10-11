// package main

// import (
// 	"os"
// )

// // 写 []byte
// func write() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
// 	f.Write([]byte("hello golang!"))
// 	f.Close()
// }

// // 写 字符串
// func writeString() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
// 	f.WriteString("hello kotlin!!")
// 	f.Close()
// }

// func main() {
// 	// write()
// 	writeString()
// }
