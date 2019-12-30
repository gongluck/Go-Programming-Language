package main

import (
	"fmt"
	"os"
)

var dir string

func init() {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println("Can not get work dir.")
	}
}

func main() {
	fmt.Println("work dir is ", dir)
}
