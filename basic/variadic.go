package main

import "fmt"

/* ประเภทของพารามิเตอร์ จะใส่ ...เข้าไปหน้าประเภทข้อมูล พารามิเตอร์แบบนี้ต้องเป็นลำดับสุดท้ายของฟังก์ชันเท่านั้น ตอนใช้งานสามารถส่งค่าพารามิเตอร์หลายๆค่า หรือไม่ใส่ก็ได้ */

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3))
	// เรียกใช้ข้อมูลแบบ slice
	xs := []int{1, 2, 3}
	// เอาค่าใน slice xs ไปทำงานในฟังก์ชัน add
	fmt.Println(add(xs...))
}
