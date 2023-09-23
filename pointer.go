package main

import "fmt"

type Address struct {
	City, Province, Contry string
}


// materi pointer parameter di function
func ChangeCountryToIndonesia(address *Address){
	address.Contry = "Jerman"
}

func main() {
	var address1 Address = Address{"Jakarta", "Yogya", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	// &Address dia membuat baru dengan reference
	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// merubah semua nilai addres 1 atau 2 menjadi nilai addres2 dibawah
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	// new
	var address4 *Address = new(Address)
	address4.City = "Jogja"
	fmt.Println(address4)


	// materi function parameter pointer
	var alamat = Address{
		City: "Subang",
		Province : "Jabar",
		Contry : "",
	}

	// var alamatPointer *Address = &alamat
	// ChangeCountryToIndonesia(alamatPointer)

	ChangeCountryToIndonesia(&alamat)
	fmt.Println("func pointer",alamat)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}