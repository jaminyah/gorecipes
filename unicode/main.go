package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	// convert to slice of rune
	//runes := []rune(username)

	for _, unicodeItem := range username {
		fmt.Print(string(unicodeItem), " ")
	}
	fmt.Println()
}
