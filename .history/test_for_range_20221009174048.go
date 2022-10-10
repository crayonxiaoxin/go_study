package main

import "fmt"

func f1() {
	var a = []int{1, 2, 3, 5}
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {

}
