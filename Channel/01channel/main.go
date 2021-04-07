package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(cap(c))
}

// 把一个有缓冲的channel 关闭后，再读取channel 的值都会是 0
