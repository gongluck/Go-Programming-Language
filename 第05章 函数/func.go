package main

import (
	"fmt"
)

func main() {
	var f []func()
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	defer func() {
		for _, fun := range f {
			fun()
		}
	}()

	for i, v := range s {
		// 捕获迭代变量
		i := i
		v := v
		f = append(f, func() {
			fmt.Println(i, " is ", v)
		})
	}
}
