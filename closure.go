package main

import "fmt"

func main() {

	// closure สามารถ return function ได้เลย
	// โดยไม่ต้องเรียกใช้งาน function

	num := func(x, y int) int {
		return x + y
	}
	fmt.Println(num(3, 4))

	// ตัวแปร x มีค่าเท่ากับ 0
	x := 0
	// ตัวแปร increment เก็บค่าจาก func() โดย x จะบวกค่าไปเรื่อยๆ ตามจำนวนที่เราเรียกใช้งาน
	increment := func() int {
		x++
		return x
	}
	// result = 1
	fmt.Println(increment())
	// result = 2
	fmt.Println(increment())

}
