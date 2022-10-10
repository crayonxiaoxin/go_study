package main

import "fmt"

func slice1() {
	var a1 []int
	a1 = append(a1, 0)
	fmt.Printf("a1: %v\n", a1)
}

func main() {
	slice1()
}
