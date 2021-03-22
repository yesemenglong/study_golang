package main

import (
	"fmt"
	"unsafe"
)

// 如何打印出接口类型的 Hash 值？
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

type Person interface {
	growUp()
}

type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}

func main() {
	var qcrao = Person(Student{age: 18})

	iface := (*iface)(unsafe.Pointer(&qcrao))
	fmt.Printf("iface.tab.hash = %#x\n", iface.tab.hash)
}

// 构造接口 qcrao 的时候，即使我把 age 写成其他值，得到的 hash 值依然不变的，
//这应该是可以预料的，hash 值只和他的字段、方法相关。
