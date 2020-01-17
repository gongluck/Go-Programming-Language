//导出C库
package main

import "C"
import "fmt"

//export GoPrint
func GoPrint(value string) {
	fmt.Println(value)
}

func main() { //main函数是必须的 有main函数才能让cgo编译器去把包编译成C的库
}
