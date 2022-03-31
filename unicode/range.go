package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for _, unicodeItem := range username {
		fmt.Print(string(unicodeItem), " ")
	}
	fmt.Println()
}
