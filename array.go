package main

import "fmt"

func main() {
	// ประกาศตัวแปร x เก็บข้อมูล 5 ตัว เป็นตัวเลข
	var x [5]int
	// index 0 มีค่าเท่ากับ 1
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	// แสดงผลลัพธ์ทั้งหมด [1 2 3 4 5]
	fmt.Println("All Array : ", x)
	// แสดงข้อมูล index 4 คือ 5
	fmt.Println("Array index 4 = ", x[4])

	arry()

}
func arry() {
	// ตัวแปร x เก็บข้อมูล 5 ตัว แบบทศนิยม
	x := [5]float64{1, 2, 3, 4, 7}
	// total = 0
	var total float64
	// ตัวแปร value เก็บค่าตามจำนวน ตัวแปร x
	for _, value := range x {
		// บวกเพิ่มไปเรื่อย ๆ จนครบทุกตัว
		// 0 + 1 + 2 + 3 + 4 +7
		total += value
	}
	fmt.Println("Total =", total)
	// หาค่าเฉลี่ย
	// ใส่ len() เพื่อนับจำนวนใน array
	fmt.Println("Avg =", (total / float64(len(x))))

}
