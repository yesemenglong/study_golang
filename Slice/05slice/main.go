package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]    // 2,3,4
	s2 := s1[2:6:7] // 4,5,6,7
	fmt.Println(cap(s2))
	s2 = append(s2, 100)
	s2 = append(s2, 200)
	fmt.Println(cap(s2))
	fmt.Println(s1)
	fmt.Println(s2) // [4 5 6 7 100 200]
	fmt.Println(s)  // [0 1 2 3 4 5 6 7 100 9]
	// s 中 下标为 9 值变成了 100，是因为 s2 = append(s2, 100) 的原因
	// 下标为 10 的为啥没变：s2 的cap是5， s2 = append(s2, 200) 操作使s2 进行了扩容，s2将原有的元素复制到新的位置，cap翻倍，变为 10
}
