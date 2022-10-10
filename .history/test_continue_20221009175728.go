package main

import "fmt"

func c1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数: %v\n", i)
		} else {
			continue
			fmt.Printf("奇数: %v\n", i)
		}
	}
}

func main() {
	c1()
}
