//导出C库
//go build -buildmode=c-shared -o export.so export.go
//go build -buildmode=c-archive -o export.a export.go

//查看函数表
//nm -g --defined-only export.so
//nm -g --defined-only export.a

package main

import "C"
import "fmt"

//export GoPrint
func GoPrint(value string) {
	fmt.Println(value)
}

func main() { //main函数是必须的 有main函数才能让cgo编译器去把包编译成C的库
}
