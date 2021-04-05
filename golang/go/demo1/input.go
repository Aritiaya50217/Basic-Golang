package main

import "fmt"

func main() {
	// รับค่า
	fmt.Print("Input Your Number : ")
	// ตัวแปร ชื่อ input รับค่าแบบทศนิยม
	var input float64
	// แสดงผลเป็นทศนิยม
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

}
