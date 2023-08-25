package main

import "fmt"

func getFullName() (string, string) {
	return "Zul", "Ferry"
}

func main() {
	firstName, _ := getFullName()
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)
	fmt.Println(firstName)
}