package main

import "fmt"

func main() {
	fmt.Print("Enter your Number : ")
	// ตัวแปร
	var number int
	fmt.Scanf("%d", &number)
	// switch การตรวจสอบในแต่ละ case
	// ถ้า number ที่ป้อนเข้ามาเป็น...ให้ทำ ...
	switch number {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	// default คล้ายกับ else ถ้าไม่ตรงกับ case ใดให้ แสดง...
	default:
		fmt.Println("Error")

	}
}
