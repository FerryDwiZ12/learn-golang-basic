package main

import (
	"fmt"
	"strings"
)

func main()  {
	// mencari kata
	fmt.Println(strings.Contains("Zul Khifli", "Zul"))
	// memotong 
	fmt.Println(strings.Split("Zul Khifli", " "))
	// menjadi kecil huruf semua
	fmt.Println(strings.ToLower("ZulKhifli"))
	// menjadi besar huruf semua
	fmt.Println(strings.ToUpper("ZulKhifli"))
	// menjadi sama kayak uppercase
	fmt.Println(strings.ToTitle("ferry dwi zulkhifli"))
	// menghapus kiri kananya 
	fmt.Println(strings.Trim("    ferry    dwi zulkhifli ", " "))
	// menghapus kata yang kita mau
	fmt.Println(strings.ReplaceAll("Zul Zul Zul", "Zul", "Budi"))

}