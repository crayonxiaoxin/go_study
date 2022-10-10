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

func b2() {
	i := 2
	switch i {
	case 1:
		fmt.Printf("i: %v\n", i)
		break
	case 2:
		fmt.Printf("i: %v\n", i)
		// break
		fallthrough
	case 3:
		fmt.Printf("i: %v\n", i)
		break
	}
}

func main() {
	b1()
}
