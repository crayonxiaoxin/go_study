// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Persion struct {
// 	Name     string
// 	Age      int
// 	Email    string
// 	Children []string
// }

// // struct => json
// func toJson() {
// 	persion := Persion{
// 		Name:  "Tom",
// 		Age:   15,
// 		Email: "abc@rmb.ee",
// 	}
// 	b, err := json.Marshal(persion)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("string(b): %v\n", string(b))
// 	}
// }

// // json => struct
// func toStruct() {
// 	b := []byte(`{"Name":"Tom","Age":15,"Email":"abc@rmb.ee"}`)
// 	var p Persion
// 	err := json.Unmarshal(b, &p)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("p: %v\n", p)
// 	}
// }

// // json => struct 嵌套
// func nestedToStruct() {
// 	b := []byte(`{"Name":"Tom","Age":15,"Email":"abc@rmb.ee","Children":["Anta","LiNing"]}`)
// 	// var p interface{}
// 	var p Persion
// 	err := json.Unmarshal(b, &p)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("p: %v\n", p)
// 	}
// }

// func main() {
// 	toJson()
// 	toStruct()
// 	nestedToStruct()
// }
