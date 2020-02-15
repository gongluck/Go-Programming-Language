package main

import (
	"fmt"
	"math"
)

const PI = 3.14

type (
	text string
)

func main() {
	fmt.Println("Hello golang, 你好 golang")
	fmt.Println(PI)

	var a [6]byte
	fmt.Println(a)
	fmt.Println(math.MaxInt64)

	b := "中文"
	fmt.Println(b)

	var c, d, f, _ int = 1, 2, 3, 4
	fmt.Println(c, d, f)

	var e float32 = 100.1
	c = int(e)
	fmt.Println(c)

	const (
		g = 10
		h = iota
		i = 10
		j
		k = iota
	)
	fmt.Println(g, h, i, j, k)

	if i := 0; i < 5 {
		fmt.Println("i<5")
	}

	n := 0
	for {
		n++
		if n > 20 {
			break
		}
		fmt.Println("n<=20")
	}

	for i := 0; i < 3; i++ {
		fmt.Println("i = ", i)
	}

	switch {
	case n > 15:
		fmt.Println("n>15")
		fallthrough
	case n > 10:
		fmt.Println("n>10")
	default:
		fmt.Println("default")
	}

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("end")
}
