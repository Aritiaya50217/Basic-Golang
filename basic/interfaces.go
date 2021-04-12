package main
import "fmt"
		"math"

type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

/* interface จะประกาศ method แทนชุดของ field
ชุดของ method เหล่านี้จะเป็นตัวบังคับว่าไทป์ใดๆ ก็ตามที่ต้องการ implementation (นำไปใช้งาน) จะต้องเขียนรายละเอียดให้ method เหล่านี้ */

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64{
	var area float64
	for _, s:= range shapes{
		area +=s.area()
	}
	return area
}

/* interface สามารถเขียนเป็น field ได้*/

type MultiShape struct{
	shapes[]Shape
}
/* สามารถเปลี่ยน MultiShape ให้เป็น Shape ได้เพื่อดึงเอา area ออกมาด้วยการเรียก method */
func (m*MultiShape) area() float64{
	var area float64
	for _, s:= rang m.shapes {
		area+=s.aarea()
	}
	return area
}

func main() {

	// เรียกใช้ method
	fmt.Println(totalArea(&c,&r))
}