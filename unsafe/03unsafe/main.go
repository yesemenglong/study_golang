package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// string 转 slice
func main() {
	var s = "sss"
	b := []byte(s)
	fmt.Println(bytes2string(b))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}
