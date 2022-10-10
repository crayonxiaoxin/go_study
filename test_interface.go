// package main

// import "fmt"

// type USB interface {
// 	read()
// 	write()
// }

// type Computer struct {
// 	name string
// }

// func (c Computer) read() {
// 	fmt.Printf("%v read...\n", c.name)
// }

// func (c Computer) write() {
// 	fmt.Printf("%v write...\n", c.name)
// }

// type Mobile struct {
// 	name string
// }

// func (m Mobile) read() {
// 	fmt.Printf("%v read...\n", m.name)
// }

// func (m Mobile) write() {
// 	fmt.Printf("%v write...\n", m.name)
// }

// func main() {
// 	c := Computer{
// 		name: "联想",
// 	}
// 	c.read()
// 	c.write()

// 	m := Mobile{
// 		name: "小米10",
// 	}
// 	m.read()
// 	m.write()
// }
