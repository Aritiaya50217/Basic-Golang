package main

import "fmt"

func main() {
	// สร้างตัวแปร แบบที่ 1 ต้องระบุ type ทุกครั้ง
	var x string = "Hello World"
	fmt.Println(x)
	x = "oil"
	fmt.Println(x)
	var y string = "hello"
	var z string = "world"
	fmt.Println(y == z)

	// สร้างตัวแปร แบบที่ 2
	o := "hello world"
	fmt.Println(o)
}
