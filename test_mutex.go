// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var i int = 100

// var wg sync.WaitGroup

// var lock sync.Mutex

// func add() {
// 	defer wg.Done() // 2
// 	lock.Lock()     // 3
// 	i += 1
// 	fmt.Printf("i++: %v\n", i)
// 	lock.Unlock()
// }

// func sub() {
// 	defer wg.Done() // 2
// 	lock.Lock()     // 3
// 	i -= 1
// 	fmt.Printf("i--: %v\n", i)
// 	lock.Unlock()
// }

// func main() {
// 	// 1. 都是在主协程执行，同步（101和100交替出现，最后是100）
// 	// for i := 0; i < 100; i++ {
// 	// 	add()
// 	// 	sub()
// 	// }

// 	// 2. 异步，不同协程，结果不固定
// 	// for i := 0; i < 100; i++ {
// 	// 	wg.Add(1)
// 	// 	go add()
// 	// 	wg.Add(1)
// 	// 	go sub()
// 	// }
// 	// wg.Wait()

// 	// 3. 异步，加锁，最后结果和1一样
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go add()
// 		wg.Add(1)
// 		go sub()
// 	}
// 	wg.Wait()
// }
