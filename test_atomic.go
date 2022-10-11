// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )

// // 原子操作

// // 方式一：锁
// // var i int = 100

// // var lock sync.Mutex

// // func add() {
// // 	lock.Lock()
// // 	i++
// // 	lock.Unlock()
// // }
// // func sub() {
// // 	lock.Lock()
// // 	i--
// // 	lock.Unlock()
// // }

// // 方式二：atomic
// var i int32 = 100

// func add() {
// 	atomic.AddInt32(&i, 1)
// }

// func sub() {
// 	atomic.AddInt32(&i, -1)
// }

// func main() {
// 	for i := 0; i < 100; i++ {
// 		go add()
// 		go sub()
// 	}
// 	time.Sleep(time.Second * 2)
// 	fmt.Printf("i: %v\n", i)
// }
