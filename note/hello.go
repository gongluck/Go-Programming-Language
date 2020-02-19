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

	arr := [...]int{1, 2, 3: 6, 5, 7, 2, 9}
	fmt.Println(arr)
	arrlen := len(arr)
	for i := 0; i < arrlen; i++ {
		for j := i + 1; j < arrlen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)

	slice := arr[2:5]
	fmt.Println(slice)

	slice = make([]int, 3, 5)
	//slice[4] = 8 //error
	slice = append(slice, 8)
	fmt.Println(slice)

	var m map[int]string = make(map[int]string)
	m[1] = "hello"
	m[2] = "world"
	m[4] = "golang"
	delete(m, 2)
	for k, v := range m {
		fmt.Println(k, v)
	}

	var mm map[int]map[int]string
	mm = make(map[int]map[int]string)
	v, ok := mm[2]
	if !ok {
		mm[2] = make(map[int]string)
		v = mm[2]
	}
	fmt.Println(v)

	l, o := A(10, 20, "test")
	B()
	fmt.Println(l, o)

	p := person{Name: "joe"}
	p.Age = 10
	fmt.Println(p)
}

func A(a, b int, c string) (int, string) {
	defer fmt.Println("end A")
	a, b, c = 1, 2, "hello"
	return a, c
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

type person struct {
	Name string
	Age  int
}
