package main

import "fmt"

func b1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			fmt.Printf("break %v\n", i)
			break
		}
	}
}

func main() {
	b1()
}
