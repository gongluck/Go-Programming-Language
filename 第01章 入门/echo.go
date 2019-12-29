// 输出命令行参数
package main

import (
	"fmt"
	"os"
	//"strings"
)

func main() {
	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Printf("%s ", os.Args[i])
	// }

	for _, v := range os.Args {
		fmt.Printf("%s ", v)
	}

	//fmt.Println(strings.Join(os.Args[0:], " "))

	//fmt.Println(os.Args[0:])

	fmt.Println("")
}
