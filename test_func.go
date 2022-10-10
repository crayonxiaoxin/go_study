// package main

// import "fmt"

// func sum(a int, b int) (result int) {
// 	return a + b
// }

// func max(a int, b int) (max int) {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func f2() (name string, age int) {
// 	name = "hello"
// 	age = 15
// 	// return name, age
// 	return // 等价于 return name, age，因为返回值定义了名称
// }

// func f3(a int) {
// 	a += 100
// }

// func f33(s []int) {
// 	s[0] = 100
// }

// func f4(args ...int) {
// 	for _, v := range args {
// 		fmt.Printf("v: %v\n", v)
// 	}
// }

// func main() {
// 	sum := sum(50, 22)
// 	fmt.Printf("sum: %v\n", sum)

// 	max := max(451, 2567)
// 	fmt.Printf("max: %v\n", max)

// 	name, age := f2()
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Printf("age: %v\n", age)

// 	a := 100
// 	f3(a)
// 	fmt.Printf("a: %v\n", a)

// 	// 参数是 slice、map 等，会改变原始值
// 	s := []int{1, 200}
// 	fmt.Printf("修改前 s: %v\n", s)
// 	f33(s)
// 	fmt.Printf("修改后 s: %v\n", s)

// 	// 可变参数
// 	f4(999, 888)
// }
