package main

import "fmt"

func slice1() {
	var a1 []int
	a1 = append(a1, 0)
	fmt.Printf("a1: %v\n", a1)

	var a2 []string
	fmt.Printf("a2: %v\n", a2)
}

func main() {
	slice1()
}
