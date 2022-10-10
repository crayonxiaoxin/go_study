// package main

// import "fmt"

// func test1() {
// 	var m1 map[string]string // assignment to entry in nil map
// 	fmt.Printf("m1: %v\n", m1)
// 	m1["a"] = "aaa"
// 	m1["b"] = "bbb"
// 	m1["c"] = "ccc"
// 	fmt.Printf("m1: %v\n", m1)
// }

// func test2() {
// 	// var m1 map[string]string     // map声明
// 	// m1 = make(map[string]string) // 初始化
// 	var m1 = make(map[string]string)
// 	fmt.Printf("m1: %v\n", m1)
// 	m1["a"] = "aaa"
// 	m1["b"] = "bbb"
// 	m1["c"] = "ccc"
// 	fmt.Printf("m1: %v\n", m1)
// }

// func test3() {
// 	// 初始化1
// 	var m1 = map[string]string{
// 		"a": "aaa",
// 		"b": "bbb",
// 		"c": "ccc",
// 	}
// 	fmt.Printf("m1: %v\n", m1)
// 	m1["d"] = "ddd"
// 	fmt.Printf("m1: %v\n", m1)

// 	// 初始化2
// 	m2 := make(map[string]string)
// 	m2["a"] = "aaa"
// 	m2["b"] = "bbb"
// 	m2["c"] = "ccc"
// 	m2["d"] = "eee"
// 	fmt.Printf("m2: %v\n", m2)

// 	// 判断key是否存在：ok=true 存在，否则不存在
// 	value, ok := m2["e"]
// 	fmt.Printf("value: %v\n", value)
// 	fmt.Printf("ok: %v\n", ok)
// }

// func test4() {
// 	var m1 = map[string]string{
// 		"a": "aaa",
// 		"b": "bbb",
// 		"c": "ccc",
// 	}
// 	// 只获取 key
// 	for k := range m1 {
// 		fmt.Printf("key: %v\n", k)
// 	}
// 	// 只获取 value
// 	for _, v := range m1 {
// 		fmt.Printf("value: %v\n", v)
// 	}
// 	// 同时获取 key value
// 	for k, v := range m1 {
// 		fmt.Printf("%v: %v\n", k, v)
// 	}
// }

// func main() {
// 	// test1()
// 	// test2()
// 	// test3()
// 	test4()
// }
