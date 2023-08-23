package main

import "fmt"

func main() {
	var months = [...]string{
		"january",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	//  [low : high]
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println("pointer :")
	// months[5] = "diubah"
	// fmt.Println(slice1)

	// slice1[0] = "ini mei baru"
	// fmt.Println(months)

	// append = membuat array baru jika dia tidak muat kapasitasnya

	var slice2 = months[10:]
	fmt.Println(slice2)


	var slice3 = append(slice2, "zul")
	fmt.Println("print slice3 :",slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println("rubah nilai slice3 :", slice3)

	fmt.Println("slice2 : ", slice2)
	fmt.Println("months :", months)


	// make 
	// buat array dengan panjang 5 : dan 2 itu panjangnya
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Zul"
	newSlice[1] = "ferry"

	var newSlice2 = newSlice[1:]
	fmt.Println("newSlice2 :", newSlice2)

	var appendSlice2 = append(newSlice2, "hai")
	fmt.Println("appendSlice2 :", appendSlice2)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))


	// copy slice 

	// variable baru, buat slice array baru, panjang dan capacitinya
	copyNewSlice := make([]string, len(newSlice), cap(newSlice))
	// mengambil data dari newSlice masukan ke copyNewSlice
	copy(copyNewSlice, newSlice)
	fmt.Println(copyNewSlice)


	// hati hati membuat array atau slice
	// ini membuat array
	iniArray := [...]int {1,2,3,4,5}
	fmt.Println("ini Array :",iniArray)
	// ini membuat slice
	iniSlice := []int{1,2,3,4,5}
	fmt.Println("ini Slice :",iniSlice)
}