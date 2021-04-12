package main

import "fmt"

func main() {
	// slice ไม่มีการระบุค่าใน []
	// ใช้ฟังก์ชัน make ในการสร้าง slice
	arr := [5]float64{1, 2, 3, 4, 5}
	// [low:high]
	// low = ลำดับเริ่มต้น , hight = ลำดับสุดท้าย (ไม่รวมลำดับที่กล่าวมา)
	x := arr[0:5]
	fmt.Println(x)

	// run function
	func_append()
	func_copy()
}

// function slice มี 2 แบบ คือ append , copy
func func_append() {
	slice1 := []int{1, 2, 3}
	// เพิ่ม 4 ,5 ใน slice1
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}
func func_copy() {
	slice3 := []int{1, 2, 3}
	// slice4 มีขนาด 2 หน่วย
	slice4 := make([]int, 2)
	// slice4 copy slice3
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
