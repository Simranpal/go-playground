package main

import "os"

func main() {

	var s, sp string
	i := 1
	for i < len(os.Args) {
		s += sp + os.Args[i]
		sp = " "
	}
	i++
	println(s)
}
