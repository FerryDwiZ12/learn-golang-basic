package main

import "fmt"

func main() {
	var name string

	name = "Zul"
	fmt.Println(name)
	fmt.Println(len(name))

	var friendName = "agus"
	fmt.Println(friendName)

	var age = 100
	fmt.Println(age)

	umur := 23
	fmt.Println(umur)


	var (
		firstName = "ferry"
		lastName = "zul"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//constants

	const jeneng string = "ferry"
	const akirJeneng = "zul"

	fmt.Println(jeneng)
	fmt.Println(akirJeneng)

	const (
		test = "test"
		lagi = "lagi"
	)

	fmt.Println(test)
	fmt.Println(lagi)
}