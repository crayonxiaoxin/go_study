package main

import (
	"fmt"
	"hello/user"
)

func main() {
	fmt.Println("Hello Go")
	s := user.Hello()
	fmt.Printf("s: %v\n", s)
}
