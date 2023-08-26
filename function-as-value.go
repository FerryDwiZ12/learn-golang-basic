package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	// menampung func dalam variable
	sayGoodBye := getGoodBye

	result := sayGoodBye("Zul")
	fmt.Println(result)

}