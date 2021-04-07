package main

import (
	"context"
	"fmt"
)

func main() {
	keyA := "keyA"
	ctx := context.Background()
	ctxA := context.WithValue(ctx, keyA, "valA")

	keyC := "keyA"
	ctxC := context.WithValue(ctxA, keyC, "eggo")

	fmt.Println("keyA =>", ctxC.Value(keyA))
	fmt.Println("keyC =>", ctxC.Value(keyC))

	type keyTypeE string
	type keyTypeF string
	var keyE keyTypeE = "keyE"
	ctx1 := context.Background()
	ctxE := context.WithValue(ctx1, keyE, "valE")
	var keyF keyTypeF = "keyE"
	ctxF := context.WithValue(ctxE, keyF, "eggo")
	fmt.Println("keyE =>", ctxF.Value(keyE))
	fmt.Println("keyF =>", ctxF.Value(keyF))
}
