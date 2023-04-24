package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
