package main

import (
	"bufio"
	"fmt"
	"strings"
)

type UserData struct {
	Environment string
	UserData    string
}

var userDataSlice []UserData

func structure(reader *bufio.Reader) {
	var (
		newUserData UserData
		addMore     string
	)

	fmt.Println("Enter the following details to add a new record for user data: ")
	fmt.Println("Environment: ")
	env, _ := reader.ReadString('\n')
	newUserData.Environment = env

	fmt.Println("UserData: ")
	data, _ := reader.ReadString('\n')
	newUserData.UserData = data

	userDataSlice = append(userDataSlice, newUserData)

	fmt.Println("Do you want to add more entries? (Y/N)")
	addMore, _ = reader.ReadString('\n')
	addMore = strings.TrimSpace(strings.ToUpper(addMore))

	if addMore == "Y" {
		structure(reader)
	} else {
		fmt.Println("List of all the user data entries: ")

		for _, user := range userDataSlice {
			fmt.Printf("Environment: %s\nUserData: %s\n", user.Environment, user.UserData)
		}
	}
}
