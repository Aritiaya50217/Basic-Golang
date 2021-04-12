package main

import "fmt"

func main() {
	/* panic คือ การส่งข้อความบอกข้อผิดพลาดออกมา สามารถจัดการกับ panic ที่เกิดขึ้นระหว่างโปรแกรมทำงานได้ โดยใช้ฟังก์ชัน recover ซึ่ง recover จะหยุดการ panic แล้วส่งค่ากลับไปที่จุดเกิด panic   */

	// เรียก recover ผ่าน defer
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
