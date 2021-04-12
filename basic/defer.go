package main

import (
	"fmt"
	"os"
)

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	/* defer จะใช้กำหนดฟังก์ชันที่ถูกเรียกใช้งาน
	เมื่อฟังก์ชันหลักที่ครอบ defer อยู่ทำงานเสร็จ*/
	// ผลลัพธ์จะได้ 1st ตามด้วย 2nd
	defer second()
	first()
	// run function
	testDefer()
}

// defer มักเอาไปใช้คืนค่าของทรัพยากรระบบเมื่อใช้งานเสร็จแล้ว
func testDefer() {
	// ทำการเปิดไฟล์
	f, _ := os.Open(filename)
	// ปิดเมื่อเลิกใช้
	defer f.Close()
}
