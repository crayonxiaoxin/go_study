package main

import "fmt"

func arr1() {
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}

func main() {
	arr1()
}
