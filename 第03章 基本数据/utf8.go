package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "golang 中文"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%d\t%q\n", i, r)
		i += size
	}
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
