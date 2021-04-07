package main

import (
	"fmt"
	"unsafe"
)

// 使用 unsafe 修改私有成员
type Programmer struct {
	name     string
	age      int
	language string
}

func main() {
	p := Programmer{"stefno", 1, "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)
}
