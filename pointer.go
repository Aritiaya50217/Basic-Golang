package main

import "fmt"

func main() {
	pointer()
	CreatePointer()

}
func pointer() {
	x := 10
	fmt.Printf("value is %d\n", x)
	// &x การเข้าถึงที่อยู่ของ x
	// ผลลัพธ์  c000014078 คือเลขฐาน 2 ที่เก็บข้อมูลชุดนี้
	fmt.Printf("Address x variable %x\n", &x)
}

// สร้าง Pointer
func CreatePointer() {
	x := 10
	// p จัดเก็บ address ของตัวแปร x
	var p *int
	// ชี้ไปยัง address ที่ตัวแปร x อยู่
	p = &x
	fmt.Printf("Pointer P = %x\n", p)
	// เปลี่ยนข้อมูลใน x โดยใช้ address(p) มา * 20
	*p = 20
	// แสดงค่า x ที่เปลี่ยนไป
	fmt.Printf("value is  %d\n", x)
}
