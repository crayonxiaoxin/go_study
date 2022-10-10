package main

import "fmt"

func c1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数: %v\n", i)
		} else {
			fmt.Printf("奇数: %v\n", i)
			continue
		}
	}
}

func main() {
	c1()
}
