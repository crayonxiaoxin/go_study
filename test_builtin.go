// package main

// import "fmt"

// // print len append panic 等

// func testAppend() {
// 	s := []int{1, 2, 3}
// 	i := append(s, 100)
// 	fmt.Printf("i: %v\n", i)

// 	s1 := []int{4, 5, 6}
// 	i2 := append(s, s1...)
// 	fmt.Printf("i2: %v\n", i2)
// }

// func testNew() {
// 	b := new(bool) // 返回指针
// 	fmt.Printf("b: %T\n", b)
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("b: %v\n", *b)
// }

// func testMake() { // slice map chan
// 	i := make([]int, 10, 100) // 长度为10，容量为100
// 	fmt.Printf("i: %v\n", i)  // 10个初始值为0的切片

// 	i2 := new([]int)
// 	fmt.Printf("i2: %v\n", i2) // 指针 &[]
// }

// func main() {
// 	// testAppend()
// 	// testNew()
// 	testMake()
// }
