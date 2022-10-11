// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// func test_add_sub() {
// 	var i int32 = 100
// 	atomic.AddInt32(&i, 1)
// 	fmt.Printf("i: %v\n", i)
// 	atomic.AddInt32(&i, -1)
// 	fmt.Printf("i: %v\n", i)

// 	var j int64 = 100
// 	atomic.AddInt64(&j, 1)
// }

// func test_read_write() {
// 	var i int32 = 100
// 	atomic.LoadInt32(&i) // read
// 	fmt.Printf("i: %v\n", i)

// 	atomic.StoreInt32(&i, 200) // write
// 	fmt.Printf("i: %v\n", i)
// }

// func test_cas() {
// 	// cas  CompareAndSwap
// 	var i int32 = 100
// 	b := atomic.CompareAndSwapInt32(&i, 100, 200) // 如果是100，则改成200
// 	fmt.Printf("b: %v\n", b)                      // true
// 	fmt.Printf("i: %v\n", i)                      // 200
// 	c := atomic.CompareAndSwapInt32(&i, 100, 300) // 如果是100，则改成300
// 	fmt.Printf("b: %v\n", c)                      // false（实际上值变成了200，200!=100）
// 	fmt.Printf("i: %v\n", i)                      // 200
// }

// func main() {
// 	// test_add_sub()
// 	// test_read_write()
// 	test_cas()
// }
