package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// ฟังก์ชันแบบพิเศษ
/* สร้าง method ให้ structs ด้วยการประกาศ func ขึ้นมาแล้วเพิ่ม receiver เข้าไประหว่างค่า func ในที่นี้เราใช้ (c*Circle) เราสามารถเรียกใช้ func ได้ด้วยการใช้ .
 */
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r

}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
