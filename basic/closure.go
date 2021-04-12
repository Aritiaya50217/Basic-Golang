package main

import "fmt"

func main() {
	// closure คือ ฟังก์ชันที่ไม่ต้องมีชื่อเรียก
	add := func(x, y int) int {
		return x * y
	}
	fmt.Println(add(1, 1))

	// return function
	nextEven := makeEvenGenerator()
	// result = 0
	fmt.Println(nextEven())
	// result = 2
	fmt.Println(nextEven())
	// result = 4
	fmt.Println(nextEven())
	// function เรียกตัวเอง
	fmt.Println(factorial(5))

}

// เขียนฟังก์ชันที่ส่งค่าออกมาเป็นฟังก์ชันอื่น
func makeEvenGenerator() func() uint {
	// ค่าเริ่มต้น = 0
	i := uint(0)
	return func() (ret uint) {
		// สร้างตัวแปร ret
		ret = i
		// บวกเพิ่มทีละ 2
		i += 2
		return
	}
}

// function เรียกตัวเอง
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	// มีการเรียกใช้ชื่อฟังก์ชันตัวเอง
	return x * factorial(x-1)
}
