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

func arr4() {
	// 数组初始化，默认省略长度，指定索引对应的值
	var a1 = [...]int{0: 91, 3: 21, 10: 100}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a1 len: %v\n", len(a1))

	var a2 = [...]bool{0: true, 3: false, 5: true}
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a2 len: %v\n", len(a2))

	// 根据 长度和下标 遍历
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}

	// 根据 range 遍历
	for _, v := range a1 {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	// arr1()
	// arr2()
	// arr3()
	arr4()
}
