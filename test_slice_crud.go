// package main

// import "fmt"

// // add
// func test1() {
// 	var s1 = []int{}
// 	s1 = append(s1, 100)
// 	s1 = append(s1, 200)
// 	s1 = append(s1, 300)
// 	fmt.Printf("s1: %v\n", s1)
// }

// // del
// func test2() {
// 	var s1 = []int{1, 2, 3, 4, 5, 6}
// 	// 删除了索引为2的元素，即3
// 	s1 = append(s1[:2], s1[3:]...)
// 	fmt.Printf("s1: %v\n", s1)
// }

// // update
// func test3() {
// 	var s1 = []int{1, 2, 3, 4, 5, 6}
// 	s1[2] = 100
// 	fmt.Printf("s1: %v\n", s1)
// }

// // copy
// func test4() {
// 	var s1 = []int{1, 2, 3, 4, 5, 6}
// 	s2 := s1                   // 赋值地址
// 	s2[2] = 100                // 同时修改索引为2的值
// 	fmt.Printf("s1: %v\n", s1) // [1 2 100 4 5 6]
// 	fmt.Printf("s2: %v\n", s2) // [1 2 100 4 5 6]

// 	// var s3 = []int{} // index out of range [2] with length 0
// 	var s3 = make([]int, len(s2))
// 	copy(s3, s1)               // 拷贝
// 	s3[2] = 200                // 不会改变s1
// 	fmt.Printf("s2: %v\n", s2) //[1 2 100 4 5 6]
// 	fmt.Printf("s3: %v\n", s3) // [1 2 200 4 5 6]
// }

// func main() {
// 	// test1()
// 	// test2()
// 	// test3()
// 	test4()
// }
