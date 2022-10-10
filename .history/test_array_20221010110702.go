package main

import "fmt"

func arr1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [4]bool
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a3: %T\n", a3)
	fmt.Printf("a1: %v\n", a1) // 默认值 [0 0]
	fmt.Printf("a2: %v\n", a2) // 默认值 [   ]
	fmt.Printf("a3: %v\n", a3) // 默认值 [false false false false]
}

func arr2() {
	// 数组初始化
	var a1 = [2]int{1, 2} // [1 2]
	// var a1 = [3]int{1, 2} // [1 2 0]
	fmt.Printf("a1: %v\n", a1)

	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
}

func arr3() {
	// 数组初始化，默认省略长度
	var a1 = [...]int{1, 2} // [1 2]
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a1 len: %v\n", len(a1))

	var a2 = [...]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a2 len: %v\n", len(a2))
}

func main() {
	// arr1()
	// arr2()
	arr3()
}
