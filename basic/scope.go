package main

import "fmt"

// การใช้งานแบบ scope
// ประกาศตัวแปรไว้ด้านนอก function อื่นเรียกใช้ได้
var x string = "Hello World"

func main() {
	fmt.Println(x)
	// เรียกใช้ function
	f()
}
func f() {
	fmt.Println(x)
}
