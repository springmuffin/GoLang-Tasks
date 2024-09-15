package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {
	fmt.Print("#### Welcome to my Go tasks program ####")
	var option string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Options(
				huh.NewOption("Calculator", "Calculator"),
				huh.NewOption("Odd", "Odd"),
				huh.NewOption("Sum", "Sum"),
				huh.NewOption("Factorial", "Factorial"),
				huh.NewOption("Min/Max Arr", "Min/Max Arr"),
				huh.NewOption("Structure", "Structure"),
				huh.NewOption("Hashmap", "Hashmap"),
			).Value(&option),
		),
	)
	form.Run()

	switch option {
	case "Calculator":
		calculator()
	case "Odd":
		odd()
	case "Sum":
		sum()
	case "Factorial":
		factorial()
	case "Min/Max Arr":
		minMaxArr()
	case "Structure":
		reader := bufio.NewReader(os.Stdin)
		structure(reader)
	case "Hashmap":
		hashmap()
	default:
		fmt.Println("Invalid selection")
	}

	// calculator()
	// odd()
	// sum()
	// factorial()
	// minMaxArr()

	// reader := bufio.NewReader(os.Stdin)
	// structure(reader)

	// hashmap()
}
