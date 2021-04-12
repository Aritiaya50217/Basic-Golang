package main

import "fmt"

func main() {
	// ตัวแปร x เก็บข้อมูลไว้ในอาร์เรย์ 5 หน่วย
	var x [5]int
	// กำหนดให้สมาชิกลำดับที่ 5 ของอาร์เรย์ x เป็น 100
	x[4] = 100
	fmt.Println(x)

	//
	test()
}
func test() {
	// สร้างอาร์เรย์ความยาว 5 หน่วย แล้วกำหนดค่าของสมาชิกแต่ละตัว
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	// i = 0 คือ ค่าเริ่มต้น
	// len(y) = นับตามจำนวนค่าในอาร์เรย์
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	// แปลง len() ให้เป็น float64  เหมือนกับ total
	fmt.Println(total / float64(len(y)))
}

// แบบที่ 2
func test2() {
	var z [5]float64
	z[0] = 98
	z[1] = 93
	z[2] = 77
	z[3] = 82
	z[4] = 83

	var total float64 = 0
	// value เปรียบเสมือน z[0]
	// _ ใช้เป็นตัวแปรเพื่อบอกคอมไพเลอร์ว่า ไม่ต้องการตัวแปรนั้น
	for _, value := range z {
		total += value
	}
	fmt.Println(total / float64(len(z)))
}
