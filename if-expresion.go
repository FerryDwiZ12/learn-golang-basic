package main

import "fmt"

func main() {
	name := "Zul"
	if name == "Zul" {
		fmt.Println("Hallo Zul")
	}

	username := "oda"
	if username == "odoy"{
		fmt.Println("hallo odoy")
	}else{
		fmt.Println("kenalan bwang !")
	}

	username1 := "jamal"
	if username1 == "odoy"{
		fmt.Println("hallo odoy")
	}else if username1 == "jamal"{
		fmt.Println("hallo bwang jamal")
	}else{
		fmt.Println("kenalan bwang !")
	}

	// sort statement 
	// length := len(name); bagian ini bisa di singkat

	if length := len(name) ;length > 5 {
		fmt.Println("terlalu panjang")
	}else{
		fmt.Println("nama sudah benar")
	}

}