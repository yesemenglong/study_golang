package main

import (
	"fmt"
	"time"
)

// select 超时控制
func main() {
	//retCh:= make(chan string, 1)
	select {
	case retCh := <-AsyncService():
		fmt.Println(retCh)
	case <-time.After(2 * time.Second):
		fmt.Println("time out")
	}

}

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned")

		retCh <- ret
	}()
	return retCh

}
