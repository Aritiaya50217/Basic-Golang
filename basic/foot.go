package main

import "fmt"

func main() {
	// รับค่า
	fmt.Printf("Length (ft) : ")
	var foot float64
	fmt.Scanf("%f", &foot)

	// m เป็น ft จะหาได้จาก m = ft/3.2808
	// แปลงฟุตเป็นเมตร
	m := (foot * 0.3048)
	fmt.Println("Length (m) : ", m)
}
