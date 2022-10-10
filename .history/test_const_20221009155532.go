package main

import "fmt"

func main() {
	const PI float64 = 3.14
	const PI2 = 3.1415

	const (
		a = 100
		b = 200
	)

	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}
