package main

import "fmt"

func add(x, y int) int {
	return x + y
}

//multiple returns
func multif(x, y, z string) (string, string, string) {
	return y, x, z
}

func main() {
	fmt.Println(add(42, 13))
	a, b, c := multif("white", "red", "black")
	fmt.Println(a, b, c)
}
