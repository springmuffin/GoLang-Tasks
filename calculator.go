package main

import "fmt"

func calculator() {
	fmt.Println("############## Welcome to calculator app ##############")

	var (
		operation       int
		first           int
		second          int
		result          int
		operationStatus bool = true
	)

	fmt.Printf("First operand = ")
	fmt.Scan(&first)

	fmt.Printf("Second operand = ")
	fmt.Scan(&second)

	fmt.Println("Please enter the type of operation you would like to perform:")
	fmt.Printf("Add (1) \nSubtract (2) \nMultiply (3) \nDivide (4) \nYour Input: ")

	fmt.Scan(&operation)

	switch operation {
	case 1:
		result = first + second

	case 2:
		result = first - second

	case 3:
		result = first * second

	case 4:
		if second == 0 {
			operationStatus = false
			fmt.Printf("This operation will result in a divide by zero error")
		} else {
			result = first / second
		}

	default:
		fmt.Println("Incorrect input")
		operationStatus = false
	}

	if operationStatus {
		fmt.Printf("\nResult of the operation = %v", result)
	} else {
		fmt.Println("\nResult could not be determined")
	}

}
