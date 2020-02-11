package main

import "fmt"

func main() {
	var mymap map[string]int
	mymap = make(map[string]int)
	mymap["cars"] = 2
	i := mymap["cars"]
	fmt.Println("value of mymap existing key \"cars\":", i)
	//if value does not exist in map
	j := mymap["bike"]
	fmt.Println("value of mymap undefined key \"bike\":", j)
}
