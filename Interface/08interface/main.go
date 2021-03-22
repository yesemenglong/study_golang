package main

import "io"

type myWriter struct {
}

func (w myWriter) Write(p []byte) (n int, err error) {
	return
}

func main() {
	// 检查 *myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = (*myWriter)(nil)

	// 检查 myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = myWriter{}
}

//上述赋值语句会发生隐式地类型转换，在转换的过程中，
//编译器会检测等号右边的类型是否实现了等号左边接口所规定的函数。
