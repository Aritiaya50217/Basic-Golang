package main

import (
	"fmt"
	"math"
)

// structs เป็นประเภทข้อมูลที่ใช้เก็บฟิลด์ประเภทต่างๆ
type Circle struct {
	//ถ้าฟิลด์ประเภทเดียวกันสามารถรวมเป็นกลุ่มไว้ในบรรทัดเดียวได้
	x, y, r float64
}

// เราสามารถเข้าถึง fields ใน structs ได้ด้วยเครื่องหมาย .
/* อาร์กิวเมนท์ จะถูกส่งผ่านโดยการคัดลอกค่าเสมอ ถ้าเราต้องการเปลี่ยนค่าอะไรก็ตามใน function circleArea มันจะไม่กลับไปเปลี่ยนค่าต้นทาง จึงต้องส่งอาร์กิวเมนท์ใหม่ให้ส่งเป็น pointer ไปแทนที่ในกรณีที่เราต้องการแก้ไข
 */
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
func main() {
	// กำหนดค่าตั้งต้น (Initialization)
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))

}
