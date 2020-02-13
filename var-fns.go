package main

import "fmt"

//var len args
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
	total += val
	}
	return total
	}
	
}

func main() {
	
	fmt.Println(sum())
	 //"0"
	fmt.Println(sum(3))
	 //"3"
	fmt.Println(sum(1, 2, 3, 4)) //
}
