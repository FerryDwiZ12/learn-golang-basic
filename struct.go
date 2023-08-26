package main

import "fmt"

type Customer struct {
	Name, Addres string
	Age          int
}

func main() {
	var eko Customer
	eko.Name = "Zul"
	eko.Addres = "yogyakarta"
	eko.Age = 23

	fmt.Println(eko.Name)
	fmt.Println(eko.Addres)
	fmt.Println(eko.Age)


	// struct Literals

	joko := Customer {
		Name: "Joko",
		Addres: "Jogja",
		Age: 22,
	}

	fmt.Println(joko)


	budi := Customer{"Budi", "Jakarta", 19}
	fmt.Println(budi)

}