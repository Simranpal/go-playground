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
	//length of map
	fmt.Println("Length of map: ", len(mymap))

	mymap["trucks"] = 1
	//iterate over the map
	for key, value := range mymap {
		fmt.Printf("\nKey = %s, Value = %d\n", key, value)
	}

	//initialize map with values
	map2 := map[string]int{
		"sum":  3711,
		"perc": 2138,
		"gra":  1908,
		"mes":  912,
	}
	fmt.Println("\nLength of map2: ", len(map2))

	val, ifExists := map2["gra"]
	fmt.Printf("\n\"gra\" val = %d, ifExists = %t \n", val, ifExists)
}
