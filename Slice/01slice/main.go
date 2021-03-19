package main

import (
	"fmt"
)

func main() {
	// 基于数组定义切片
	// a := [5]int{55, 56, 57, 58, 59}
	// b := a[1:4] // 基于数组a创建切片，包括元素a[1],a[2],a[3]
	// fmt.Println(b)
	// fmt.Printf("type of b:%T\n", b)

	// 切片再切片
	// a := [...]string{"北京", "上海", "广州", "成都", "深圳", "重庆"}
	// fmt.Printf("a:%v type:%T  len:%d  cap:%d\n", a, a, len(a), cap(a))
	// b := a[1:3]
	// fmt.Printf("b:%v  type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	// c := b[1:5]
	// fmt.Printf("c:%v  type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))

	a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := a[:5]
	fmt.Printf("len:%d  cap:%d\n", len(s1), cap(s1))
	s2 := a[1:2]
	fmt.Printf("len:%d  cap:%d\n", len(s2), cap(s2))
	s3 := a[1:3]
	fmt.Printf("len:%d  cap:%d\n", len(s3), cap(s3))
	s4 := a[2:3]
	fmt.Printf("len:%d  cap:%d\n", len(s4), cap(s4))
}
