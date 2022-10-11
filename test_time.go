// package main

// import (
// 	"fmt"
// 	"time"
// )

// func test1() {
// 	t := time.Now()
// 	fmt.Printf("t: %T\n", t)
// 	fmt.Printf("t: %v\n", t)

// 	year := t.Year()
// 	month := t.Month()
// 	day := t.Day()
// 	hour := t.Hour()
// 	minute := t.Minute()
// 	second := t.Second()
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// 	fmt.Printf("%T-%T-%T %T:%T:%T\n", year, month, day, hour, minute, second)
// }

// func test2() {
// 	t1 := time.Now()
// 	timestamp := t1.Unix()
// 	fmt.Printf("type: %T, timestamp: %v\n", timestamp, timestamp)

// 	t := time.Unix(timestamp, 0) // timestamp 转 时间
// 	year := t.Year()
// 	month := t.Month()
// 	day := t.Day()
// 	hour := t.Hour()
// 	minute := t.Minute()
// 	second := t.Second()
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// 	fmt.Printf("%T-%T-%T %T:%T:%T\n", year, month, day, hour, minute, second)

// }

// func format() {
// 	t := time.Now()
// 	s := t.Format("2006-01-02 15:04:05") // 固定是 2006-01-02 15:04:05 （2006 1 2 3 4 5），而不是 Y-m-d H:i:s
// 	fmt.Printf("s: %v\n", s)
// }

// func main() {
// 	// test1()
// 	// test2()
// 	format()
// }
