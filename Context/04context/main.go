package main

import (
	"context"
	"fmt"
	"time"
)

// 任务取消

func main(){
	ctx,cancel := context.WithCancel(context.Background())
	for i:=0;i<5;i++{
		go func(i int,ctx context.Context) {
			for {
				if isCancelled(ctx){
					break
				}
				time.Sleep(time.Millisecond*5)
			}
			fmt.Println(i,"Cancelled")
		}(i,ctx)
	}
	cancel()
	time.Sleep(time.Second*1)
}


func isCancelled(ctx context.Context)bool{
	select {
	case <- ctx.Done():
		return true
	default:
		return false
	}
}

