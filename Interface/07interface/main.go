package main

import (
	"fmt"
	"unsafe"
)

// 如何打印出接口的动态类型和值？
type iface struct {
	itab, data uintptr
}

func main() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)
	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}

// a 的动态类型和动态值的地址均为 0，也就是 nil；
// b 的动态类型和 c 的动态类型一致，都是 *int；
// 最后，c 的动态值为 5。
