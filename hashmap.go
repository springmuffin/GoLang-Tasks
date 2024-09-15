package main

import "fmt"

func hashmap() {
	environment := []string{"dev", "qc", "uat", "prod"}
	ec2_instance := []string{"t3.micro", "t3.small", "t3.medium", "t3.large"}
	hashMap := make(map[string]string)

	for i := 0; i < len(environment); i++ {
		hashMap[environment[i]] = ec2_instance[i]
	}

	for key, value := range hashMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
