package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	// convert to slice of rune
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
	fmt.Println()
}
