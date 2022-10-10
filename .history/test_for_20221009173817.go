package main

import "fmt"

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f2() {
	// 初始值可以写在外面
	i := 0
	for ; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f3() {
	i := 0
	for i < 10 {
		fmt.Printf("i: %v\n", i)
		// 结束条件可以写在里面
		i++
	}
}

func f4() {
	for {
		fmt.Println("一直执行...")
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
