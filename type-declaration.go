package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtmZul NoKTP = "2842312620"
	var marriedStatus Married = false
	fmt.Println(noKtmZul)
	fmt.Println(marriedStatus)
}