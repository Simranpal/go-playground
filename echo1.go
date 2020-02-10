package main

import "os"

func main() {

	//declare 2 string variables
	var s, sp string
	for i := 1; i < len(os.Args); i++ {
		s += sp + os.Args[i]
		sp = " "
	}
	println(s)
}

/* way to declare a variable
by default all variables initialized to 0, ""
s := "sss" declare & initialize
var s string
var s = ""
var s string = ""
*/
