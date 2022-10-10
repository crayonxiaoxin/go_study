// package main

// import "fmt"

// type Persion struct {
// 	id, age int
// 	name    string
// }

// // 值传递
// func showPersion(persion Persion) {
// 	persion.name = "Jerry"
// 	persion.age = 88
// 	fmt.Printf("persion: %v\n", persion)
// }

// // 指针传递
// func showPersionPointer(persion *Persion) {
// 	persion.name = "Jerry"
// 	persion.age = 88
// 	fmt.Printf("persion: %v\n", *persion)
// }

// func main() {
// 	tom := Persion{
// 		id:   101,
// 		age:  35,
// 		name: "Tom",
// 	}
// 	// fmt.Printf("tom: %v\n", tom)
// 	// fmt.Println("-----------------------")
// 	// showPersion(tom)
// 	// fmt.Println("-----------------------")
// 	// fmt.Printf("tom: %v\n", tom)

// 	fmt.Printf("tom: %v\n", tom)
// 	fmt.Println("-----------------------")
// 	showPersionPointer(&tom)
// 	fmt.Println("-----------------------")
// 	fmt.Printf("tom: %v\n", tom)
// }
