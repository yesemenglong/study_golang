package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
)

// withCancel 用法
func UseContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("context is done with error %s", ctx.Err())
			return
		default:
			log.Printf("nothing just loop...")
			time.Sleep(time.Second * time.Duration(1))
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go UseContext(ctx)
	time.Sleep(time.Second * time.Duration(1))
	cancel()
	time.Sleep(time.Second * time.Duration(2))
}
