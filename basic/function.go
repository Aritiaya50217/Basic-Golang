package main

import "fmt"

func main() {

	xs := []float64{1, 2, 3, 4, 55}
	fmt.Println(average(xs))
	x, y := f()

}

// สร้าง function ในการหาค่าเฉลี่ย
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	// return เพื่อส่งค่าไปยังฟังก์ชันที่เรียกใช้
	return total / float64(len(xs))

}

// การส่งค่ากลับหลายค่า
func f() (int, int) {
	return 5, 6
}
