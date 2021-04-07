package main

import (
	"context"
	"fmt"
)

var keyA string = "keyA"
var keyC string = "keyA"

func main(){
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, keyA, "valueA")
	fmt.Println("In ctx:")
	fmt.Println("keyA => ", ctx1.Value(keyA))

	ctx2 := context.WithValue(ctx, keyC, "valueC")
	fmt.Println("In ctx2:")
	fmt.Println("keyA => ", ctx2.Value(keyA))
	fmt.Println("keyC => ", ctx2.Value(keyC))
}
