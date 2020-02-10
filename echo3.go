package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//print slice
	fmt.Println(os.Args[1:])
	// if large number of inputs
	fmt.Println(strings.Join(os.Args[1:], " "))
}
