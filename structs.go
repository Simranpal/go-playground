package main

import "fmt"

type human struct {
	city string
	job  string
}

func main() {
	fmt.Println(human{"NYC", "Chef"})
	//instance of human struct
	Simran := human{"YVR", "SE"}
	//access the struct variables with dot operator
	Simran.city = "Vancouver"
	fmt.Println("Simran:", Simran)
	//access the struct via pointer
	p := &Simran
	p.city = "Burnaby"
	fmt.Println("Simran:", Simran)
}
