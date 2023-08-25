package main

import "fmt"


// harus membuat tipe pengembalianya ()string <==
func getHello(name string) string {
	// return "Hello " + name
	if name == ""{
		return "Hello Bro"
	}else{
		return "Hello" + name
	}
}

func main() {
	result := getHello("")
	fmt.Println(result)
}