package main

import "fmt"

func main() {
	// i := 1 คือ เริ่มต้นที่ 1
	// i <= 10 คือ ผลลัพธ์จะไม่เกิน 10
	// i++ คือ บวกเพิ่มไปเรื่อยๆ จนครบ 10
	for i := 1; i <= 10; i++ {
		fmt.Println("Artitaya", i)
	}
	// เรียกใช้ function
	modulate()

}

// หาเลขคู่ เลขคี่
func modulate() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		} else {
			fmt.Println("ODD", i)
		}
	}
}
