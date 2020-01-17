package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Alignof(x))
	fmt.Println(unsafe.Offsetof(x.b))

	pb := (*int16)(unsafe.Pointer(&x.b))
	*pb = 100
	fmt.Println(x.b)
}
