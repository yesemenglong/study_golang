package main

import (
	"fmt"
	"time"
)

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("goroutine A received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("goroutine B received data: ", val)
	return
}

func main() {
	ch := make(chan int)
	go goroutineA(ch)
	go goroutineB(ch)
	ch <- 3
	time.Sleep(time.Second)
}

// Blog: https://blog.leonardwang.cn/archives/golangprunnext%E6%98%AF%E4%BB%80%E4%B9%88
// Goroutine 优先级
// P.runnext > p.localrunq > globalrunq
// 流程分析：
// 1.gA -> runnext
// 2.gA -> localrunq    gB -> runnext
// 当前状态： runnext: gB   localrunq: gA
// 所以 打印的永远是 gB
// version: go1.14
