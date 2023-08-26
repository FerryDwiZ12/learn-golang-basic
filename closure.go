package main

import "fmt"

func main() {
	name := "Zul"
	counter := 0

	increment := func(){
		name = "Apa"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	// fmt.Println("Increment")
}