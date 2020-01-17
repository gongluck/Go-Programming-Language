package main

import (
	"unsafe"
)

//
//源码中使用C代码
//

/*
#include <stdio.h>
#include <stdlib.h>
void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

//
//使用#cgo定义库路径
//

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lhello
#include "hello.h"
*/
import "C"

func main() {
	// Go语言调用C函数
	cs := C.CString("Hello World")
	defer C.free(unsafe.Pointer(cs))
	C.myprint(cs)

	C.hello()
}
