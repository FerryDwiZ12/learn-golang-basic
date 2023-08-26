package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error dengan massage : ", message)
	}
	
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	defer endApp()
	if error {
		panic("APLIKASI EROR")
	}
	
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
}