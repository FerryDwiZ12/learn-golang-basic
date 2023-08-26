package main

import "fmt"


// baris 8 - 14 tidak menggunakan reqursive func 

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialReqrusive(value int) int {
	if value == 1 {
		return 1
	}else{
		return value * factorialReqrusive(value-1)
	}
}


func main() {
	loop := factorialLoop(3)
	fmt.Println(loop)
	fmt.Println(3 * 2 * 1)


	reqursive := factorialReqrusive(3)
	fmt.Println(reqursive)
}