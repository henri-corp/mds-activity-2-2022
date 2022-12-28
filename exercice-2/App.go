package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("ERROR : ./app : Missing arg 1")
		fmt.Println("ARGS GIVEN:", os.Args)
		os.Exit(1)
	}
	fmt.Println("Hello", os.Args[1])
}
