// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	for i := 0; i <= 100; i++ {
// 		time.Sleep(50 * time.Millisecond)
// 		// j := i / 5
// 		// a := strings.Repeat("=", j) + strings.Repeat(" ", 100/5-j)
// 		// fmt.Printf("\r[%v%%][%s] ", i, a)

// 		a := strings.Repeat("=", i) + strings.Repeat(" ", 100-i)
// 		fmt.Printf("\r[%v%%][%s] ", i, a)

// 		// a := strings.Repeat(" ", 100-i) // \r 的情况下，回车会把前面的内容覆盖掉，效果显示为光标会回退
// 		// fmt.Printf("\r%v", a)

// 		// a := strings.Repeat("=", i)
// 		// fmt.Printf("\r%v", a)

// 		// fmt.Printf("\r[%v%%]", i)

// 		os.Stdout.Sync()
// 	}
// 	fmt.Println("\nok")
// }
