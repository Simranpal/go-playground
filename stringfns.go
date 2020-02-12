package main

import "fmt"

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(HasPrefix("sim", "s"))
	fmt.Println(HasPrefix("simran", "r"))
}
