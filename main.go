package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		os.Exit(-1)
	}

	csvFilePath := argsWithoutProg[0]
	fmt.Println("csvFilePath:", csvFilePath)
}
