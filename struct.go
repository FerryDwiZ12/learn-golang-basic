package main

import "fmt"

type Customer struct {
	Name, Addres string
	Age          int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello",name, "My name is", customer.Name)
}



func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Addres = "yogyakarta"
	eko.Age = 23

	fmt.Println(eko.Name)
	fmt.Println(eko.Addres)
	fmt.Println(eko.Age)

	eko.sayHi("Zul")


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