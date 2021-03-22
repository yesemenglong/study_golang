package main

import "fmt"

// for range 从通道循环取值
func main(){
	ch1 := make(chan int, 100)
	go send(ch1)

	for ret := range ch1{
		fmt.Println(ret)
	}
}

func send(c chan int)  {
	for i := 0; i<10; i++{
		c <- i
	}
	close(c)
}