package main

import (
	"fmt"
	
)

func main() {
	person := map[string]string{
		"name":   "ferry",
		"addres": "jogja",
	}

	// tambah data person
	person["tittle"] = "Programmer"
	
	// mengubah value dalam map
	person["tittle"] = "Design"

	fmt.Println("map :",person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])
	fmt.Println(person["tittle"])

	// membuat map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "zul"
	book["ups"] = "salah"

	fmt.Println(book)
 
	delete(book,"ups")

	fmt.Println(book)
}