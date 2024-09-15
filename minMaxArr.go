package main

import (
	"fmt"
	"sort"
)

func minMaxArr() {
	var arr [5]int

	fmt.Println("Please enter 5 random numbers")

	for i := range arr {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println("Incorrect value entered, please try again")
			return
		}
	}

	fmt.Println("Original Array: ", arr)

	sort.Ints(arr[:])

	fmt.Printf("Minimum: %d \nMaximum: %d\nSorted Array: %v", arr[0], arr[len(arr)-1], arr)
}
