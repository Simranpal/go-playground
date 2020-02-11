package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Red"
	a[1] = "Green"
	a[2] = "Blue"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	primenos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primenos)
}
