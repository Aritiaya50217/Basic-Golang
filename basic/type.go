package main

import "fmt"

func main() {
	// int
	fmt.Println("1 + 1 = ", 1+1)
	// float
	fmt.Println("1.0 + 1.0 = ", 1.0+1.0)
	// นับความยาวอักษร
	// result = 11
	fmt.Println(len("Hello World"))
	// เข้าถึงข้อมูล
	// result = 101 คือ e ถูกแทนที่ด้วย byte
	fmt.Println("Hello World"[1])
	// ต่อคำ
	fmt.Println("Hello" + "World")
	// boolean
	// && = and , || = or ,! = not

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println((!true))
}
