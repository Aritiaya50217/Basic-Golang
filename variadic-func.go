package main

import "fmt"

func main() {

	summation(1, 1)
	summation2(2, 2, 4)
	// ใส่พารามิเตอร์กี่ตัวก็ได้
	summation_variadic(1, 1, 1, 1)
}
func summation(a int, b int) {
	fmt.Println("total = ", a+b)
}
func summation2(a int, b int, c int) {
	fmt.Println("total = ", a+b+c)
}

// variadic func = การรับค่าแบบไม่จำกัดจำนวน
// ส่งค่าเข้ามาจัดเก็บในตัวแปร nums...
func summation_variadic(nums ...int) {
	var total int
	// _ , คือ ค่าแรกในกรณีที่เราไม่ได้กำหนด
	// range nums คือ วนลูปตามตัวเลขที่ใส่
	for _, n := range nums {
		// เอา n มาบวกเพิ่ม ไปเรื่อยๆจนครบ ตามจำนวน
		total += n
		fmt.Println("index :", total)
	}
}
