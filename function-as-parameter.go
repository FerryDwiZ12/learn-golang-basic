package main

import "fmt"

// menggunakan type func declataion 
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string  {
	if name == "Anjing" {
		return "..."
	}else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Zul", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}