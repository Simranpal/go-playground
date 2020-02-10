package main

import "os"

func main() {

	s, sp := "", ""
	for _, arg := range os.Args[1:] {
		s += sp + arg
		sp = " "
	}
	println(s)
}
