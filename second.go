package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println(os.Args[1])
	}else {
		os.Exit(1)
	}
}
