package main

import "fmt"

func f1() {
	var a = [...]int{1, 2, 3, 5}
	// for _, v := range a {
	// 	fmt.Printf("v: %v\n", v)
	// }
	for i, v := range a {
		fmt.Printf("v: %v-%v\n", i, v)
	}
}

func f2() {
	var a = []int{1, 2, 3, 5}
	for i, v := range a {
		fmt.Printf("v: %v-%v\n", i, v)
	}
}

func f3() {
	// key:value
	var m = make(map[string]string, 0)
	m["name"] = "abc"
	m["age"] = "20"
	for _, v := range m {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
