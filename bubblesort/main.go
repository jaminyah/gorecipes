package main

import "fmt"

func main() {

	array := []int{14, 8, 22, 4, 10, 31, 13, 27, 9, 25}

	swapped := true
	for swapped {
		swapped = false

		for i := 1; i < len(array); i++ {
			if array[i-1] > array[i] {
				array[i], array[i-1] = array[i-1], array[i]
				swapped = true
			}
		}
	}
	fmt.Println("After: ", array)
}
