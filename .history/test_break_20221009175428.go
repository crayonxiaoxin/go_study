// package main

// import "fmt"

// func b1() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i == 5 {
// 			fmt.Printf("break %v\n", i)
// 			break
// 		}
// 	}
// }

// func b2() {
// 	i := 2
// 	switch i {
// 	case 1:
// 		fmt.Printf("i: %v\n", 1)
// 		break
// 	case 2:
// 		fmt.Printf("i: %v\n", 2)
// 		// break
// 		fallthrough // 穿透 2,3
// 	case 3:
// 		fmt.Printf("i: %v\n", 3)
// 	case 4:
// 		fmt.Printf("i: %v\n", 4)
// 	}
// }

// func b3() {
// LABEL:
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i >= 5 {
// 			break LABEL
// 		}
// 	}
// 	fmt.Println("end...")
// }

// func main() {
// 	// b1()
// 	// b2()
// 	b3()
// }
