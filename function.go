package main

import "fmt"

func main() {
	// เรียกใช้ฟังก์ชัน
	// x คือ ข้อความที่เก็บไว้ในพารามิเตอร์(str)
	showString("x")
	// 10,20 คือ ข้อมูลที่เก็บไว้ในพารามิเตอร์ (a,b)
	addition(10, 20)

	empty()
	// ตัวแปร result เก็บค่าจาก function addition2(5, 5)
	// 5+5 = 10 การบวกมาจากใน function
	// 10 * 10 = 100

	result := addition2(5, 5)
	fmt.Println(result * 10)
}

// สร้าง function โดยมีค่าพารามิเตอร์ คือ str เก็บข้อมูลแบบ string
func showString(str string) {
	fmt.Println(str)
}

// สร้าง function
// โดยมีค่าพารามิเตอร์ คือ a,b เก็บข้อมูลแบบ integer
func addition(a int, b int) {
	fmt.Println("Sum total =", a+b)
	fmt.Println("Minus total =", a-b)
	fmt.Println("Multiply total =", a*b)
	fmt.Println("Divide total =", a/b)

}

// function ที่ไม่มีการรับค่า
func empty() {
	fmt.Println("OK")
}

// return ค่า
func addition2(a int, b int) int {
	// ต้องระบุว่า return ออกไปประเภทไหน
	output := a + b
	return output
}
