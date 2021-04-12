package main

import "fmt"

func main() {
	fmt.Printf("Fahrenheit is : ")
	var F float64
	fmt.Scanf("%f", &F)

	C := ((F - 32) * 5 / 9)
	fmt.Println("Celsius is : ", C)

}
