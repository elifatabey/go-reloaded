package main

import (
	"os"
	"strings"
)

func main() {
	//1. Reading the file - receiving argument
	file, err := os.Args[1:]
	if err != nil {
		panic(err)
	}
	theinput := string(file) //the input
	slice := strings.Fields(theinput)
	fmt.Printf(slice[1])
}
