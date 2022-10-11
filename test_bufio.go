// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func testReader1() {
// 	r := strings.NewReader("hello golang")
// 	r2 := bufio.NewReader(r)
// 	s, _ := r2.ReadString('\n')
// 	fmt.Printf("s: %v\n", s)
// }

// func testReader2() {
// 	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
// 	br := bufio.NewReader(r)
// 	p := make([]byte, 10)
// 	for {
// 		n, err := br.Read(p)
// 		if err == io.EOF {
// 			break
// 		} else {
// 			fmt.Printf("n: %v\n", n)
// 			fmt.Printf("p: %v\n", string(p))                 // 不正确的，如果最后一次只有5个，那么后5个还是旧数据
// 			fmt.Printf("string(p[:n]): %v\n", string(p[:n])) // 长度是多少，就读多少个
// 		}
// 	}
// }

// func testWrite1() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
// 	defer f.Close()
// 	w := bufio.NewWriter(f)
// 	w.WriteString("hello world 111111")
// 	w.Flush()
// }

// func testScanner1() {
// 	r := strings.NewReader("ABC DEF GHI JKL 123")
// 	s := bufio.NewScanner(r)
// 	s.Split(bufio.ScanWords) // ScanWords 以空格分割
// 	for s.Scan() {           // s.Scan() bool
// 		fmt.Printf("s.Text(): %v\n", s.Text())
// 	}
// }

// func testScanner2() {
// 	r := strings.NewReader("hello 世界！")
// 	s := bufio.NewScanner(r)
// 	// s.Split(bufio.ScanBytes) // 中文会显示乱码
// 	s.Split(bufio.ScanRunes) // 中文可以正常显示
// 	for s.Scan() {           // s.Scan() bool
// 		fmt.Printf("%v\t", s.Text())
// 	}
// }

// func main() {
// 	// testReader1()
// 	// testReader2()
// 	// testWrite1()
// 	// testScanner1()
// 	testScanner2()
// }
