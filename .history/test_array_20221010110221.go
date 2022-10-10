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

func main() {
	arr1()
}
