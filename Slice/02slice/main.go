package main

import (
	"fmt"
)

func main() {
	// make()
	// a := make([]int, 2, 10)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	//  切片的赋值拷贝
	// s1 := make([]int, 3)
	// s2 := s1 //将s1直接赋值给s2，s1和s2共用一个底层数组
	// s2[0] = 100
	// fmt.Println(s1)
	// fmt.Println(s2)
	// s1[1] = 100
	// fmt.Println(s1)
	// fmt.Println(s2)

	s := []int{1, 3, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	for index, value := range s {
		fmt.Println(index, value)
	}
}
