// package main

// import "fmt"

// func main() {
// 	var ip *int
// 	fmt.Printf("ip: %v\n", ip) // nil
// 	fmt.Printf("ip: %T\n", ip) // *int

// 	var i int = 100
// 	ip = &i                     // 取地址
// 	fmt.Printf("ip: %v\n", ip)  // 内存地址
// 	fmt.Printf("ip: %v\n", *ip) // 内存地址对应的值 100

// 	var sp *string
// 	var s string = "hello"
// 	sp = &s
// 	fmt.Printf("sp: %v\n", sp)
// 	fmt.Printf("sp: %v\n", *sp)

// 	arrayPointers()
// }

// // 数组指针
// func arrayPointers() {
// 	a := [...]int{1, 3, 5, 7}
// 	var pa [len(a)]*int
// 	fmt.Printf("pa: %v\n", pa) // [<nil> <nil> <nil> <nil>]
// 	for i := 0; i < len(a); i++ {
// 		pa[i] = &a[i]
// 	}
// 	fmt.Printf("pa: %v\n", pa)

// 	for _, v := range pa {
// 		fmt.Printf("v: %v\n", *v) // 取值
// 	}
// }
