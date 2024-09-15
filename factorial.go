package main

import "fmt"

func factorial() {
	var (
		num int
		res int = 1
	)
	fmt.Println("############## Welcome to factorial finder ##############")

	fmt.Println("Enter the number to find the factorial")

	fmt.Scan(&num)

	if num <= 0 {
		fmt.Print("Please enter a positive integer")
	} else {
		for i := num; i >= 1; i-- {
			res *= i
		}

		fmt.Println("Factorial = ", res)
	}

}
