package main

import "fmt"

func main() {
	var a *int
	var b interface{}
	b = a
	fmt.Println(b)
	// 条件始终为 'false'，因为 'b' 始终不是 'nil'
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is non-nil")
	}
}
