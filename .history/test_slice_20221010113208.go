package main

import "fmt"

func slice1() {
	var a1 []int
	a1 = append(a1, 0)
	fmt.Printf("a1: %v\n", a1)

	var a2 []string
	fmt.Printf("a2: %v\n", a2)
}

func slice2() {
	// make 方式声明
	var a1 = make([]int, 2)
	fmt.Printf("a1: %v\n", a1)
}

func slice3() {
	var a1 = []int{2, 3, 3}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("len(a1): %v\n", len(a1)) // 长度
	fmt.Printf("cap(a1): %v\n", cap(a1)) // 容量
}

func slice4() {
	var a1 = []int{1, 2, 3, 4, 5, 6, 7} // 切片，数组也可
	var a2 = a1[0:3]
	fmt.Printf("a2: %v\n", a2)

	a3 := a1[3:]
	fmt.Printf("a3: %v\n", a3)

	a4 := a1[:]
	fmt.Printf("a4: %v\n", a4)

	a5 := a1[:4]
	fmt.Printf("a5: %v\n", a5)
}

func main() {
	// slice1()
	// slice2()
	// slice3()
	slice4()
}
