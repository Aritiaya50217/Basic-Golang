package main

import "fmt"

func main() {
	// แสดงข้อความเพื่อรับค่า
	fmt.Printf("Enter a number : ")
	// เอาค่าที่รับมา เก็บไว้ในตัวแปร input
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

}
