package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}



func main() {
	var zul HasName
	
}