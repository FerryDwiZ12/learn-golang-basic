package main

import "fmt"

func getFullNamed() (firstName, middleName, lastName string) {
	firstName = "Ferry"
	middleName = "Dwi"
	lastName = "Zulkhifli"

	return
}

func main() {
	firstName, middleName, lastName := getFullNamed()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}