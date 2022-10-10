// package main

// import "fmt"

// // 阶乘 for
// func f1(a int) int {
// 	var s int = 1
// 	// 5! = 5*4*3*2*1
// 	for i := 1; i <= a; i++ {
// 		s *= i
// 	}
// 	fmt.Printf("s: %v\n", s)
// 	return s
// }

// // 阶乘 递归
// func f2(a int) int {
// 	if a == 1 {
// 		return a
// 	}
// 	return a * f2(a-1)
// }

// // 斐波那契数列  f(n)=f(n-1)+f(n-2) 且 f(2)=f(1)=1
// func f(n int) int {
// 	if n == 1 || n == 2 {
// 		return 1
// 	}
// 	return f(n-1) + f(n-2)
// }

// func main() {
// 	f1(5)
// 	r := f2(5)
// 	fmt.Printf("r: %v\n", r)

// 	r = f(5)
// 	fmt.Printf("r: %v\n", r)
// }
