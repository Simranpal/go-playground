// program to count duplicate lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

var countmap map[string]int

func main() {
	countmap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//fetch each line
	for input.Scan() {
		line := input.Text()
		countmap[line] = countmap[line] + 1
		if len(line) == 0 {
			break
		}
	}
	for line, n := range countmap {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
