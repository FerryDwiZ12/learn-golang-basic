package main

import "fmt"

func main() {
	name := "Zul"

	switch name {
	case "Zul":
		fmt.Println("Hallo Zul")
		fmt.Println("Hallo mas Zul")
	case "Joko":
		fmt.Println("Hallo Joko")
	default:
		fmt.Println("kenalan bwang")
	}

	// sort statement
	switch length := len(name) ; length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama terlalu pendek")
	}


	length2 := len(name)

	switch {
	case length2 > 10:
		fmt.Println("nama terlalu panjang")
	case length2 > 5:
		fmt.Println("nama lumayan panjang")
	default :
		fmt.Println("nama pas")
	}
}