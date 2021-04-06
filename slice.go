package main

import "fmt"

func main() {
	// สร้าง slice ในการจัดเก็บข้อมูล โดยไม่กำหนดขนา
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// การเอา slice มาต่อกัน คล้าย list ใน python
	slice2 := append(slice1, 9, 10)

	fmt.Println(slice2)

	// เรียกใช้ฟังก์ชัน
	array()
	copy1()

}
func array() {
	arr := [5]float64{1, 2, 3, 4, 5}
	// x เก็บค่าที่ 0 to 4
	x := arr[0:4]
	fmt.Println(x)
}
func copy1() {
	// เก็บข้อมูลใน  []
	a_slice := []int{1, 2, 3}
	// กำหนดขนาดข้อมูล แค่ 2 ตำแหน่ง
	b_slice := make([]int, 2)
	// copy ข้อมูลของ a_slice ลง b_slice
	copy(b_slice, a_slice)
	// result [1,2]
	fmt.Println(b_slice)
}
