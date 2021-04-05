// การประกาศตัวแปร
package main

import "fmt"

func main() {
	// ตัวแปร myAge รับค่าเป็น int
	var myAge int = 200
	// ถ้าใช้ := ไม่ต้องประกาศ var
	// := คือ รับค่าโดยไม่ต้องกำหนดประเภท
	myAge2 := 500
	var something bool = true
	// age1 = 35 and age2 = 44
	age1, age2 := 35, 44

	fmt.Println("Value Variable = ", myAge)
	fmt.Println("Value Variable = ", myAge2)
	fmt.Println("Someting", something)
	fmt.Println(age1, age2)
}
