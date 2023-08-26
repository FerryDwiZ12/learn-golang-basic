package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string)bool  {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("Eko", blackList)

	registerUser("root", func(name string)bool{
		return name == "root"
	})
}