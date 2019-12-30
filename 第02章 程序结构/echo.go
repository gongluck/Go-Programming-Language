// echo
package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Bool("n", false, "n flag")
	var s = flag.String("s", "", "test")

	flag.Parse()
	fmt.Println("n = ", *n)
	fmt.Println("s = ", *s)
}
