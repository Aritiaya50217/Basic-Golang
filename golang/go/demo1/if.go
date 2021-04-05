package main

import "fmt"

func main() {

	fmt.Print("Input Your Number : ")
	var input float64
	fmt.Scanf("%f", &input)

	// สร้างตัวแปรเก็บค่า
	condition := input > 2
	// ถ้าเป็นจริง แสดงคำว่า Worked
	if condition {
		fmt.Println("Worked")
	} else {
		fmt.Println("Error")
	}
	// เรียกใช้ฟังก์ชัน
	andd()
	pipe()
}
func andd() {
	// && (and) คือ เงื่อนไขเป็นจริงทั้งคู่

	if 6 > 3 && 5 > 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("False")
	}
}
func pipe() {

	// || (หรือ) คือ เป็นเงื่อนไขจริง 1 อัน
	if 5 > 2 || 3 > 4 {
		fmt.Println("Yes")
	} else {
		fmt.Println("result error")
	}

}
