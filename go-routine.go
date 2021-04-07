package main

import "fmt"

func main() {
	// n = 0
	TestRoutine(0)
	// สร้างการทำงานซ้อน ฟังก์ชัน TestRoutine(n int)
	// สร้างตัวแปร input เก็บค่า string
	var input string
	fmt.Scanln(&input)
}

// Routine คือ ฟังก์ชันที่สามารถเรียกใช้งาน ฟังก์ชันอื่นได้ในขณะที่มีการทำงานอยู่
func TestRoutine(n int) {
	// n คือ พารามิเตอร์
	// i = 1 แต่ไม่เกิน 10
	for i := 1; i <= 10; i++ {
		fmt.Println("index", ":", i)
	}
}
