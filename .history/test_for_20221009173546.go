package main

import "fmt"

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f3() {
	i := 0
	for i < 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

func main() {
	f1()
	f2()
	f3()
}
