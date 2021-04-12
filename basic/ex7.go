/* โจทย์ข้อที่ 1
sum คือ ฟังก์ชันที่รับค่าข้อมูลแบบ สไลช์ของตัวเลขและหาผลรวมทั้งหมดของตัวเลขที่อยู่ในสไลซ์ จงเขียนฟังก์ชัน signature ของ sum ที่รองรับการทำงานแบบนี้
*/

/*
package main

import "fmt"

func main() {
	number := []float64{10, 20, 30, 40}
	fmt.Println(sum(number))
}
func sum(number []float64) float64 {
	total := 0.0
	for _, v := range number {
		total += v
	}
	return total

}
*/

/* โจทย์ข้อที่ 2
จงเขียนฟังก์ชันที่รับเลขจำนวนเต็มและส่งค่ากลับสองค่า คือ ค่าหารสองของจำนวนนั้น และ true ถ้าจำนวนนั้นเป็นเลขคู่ หรือ false ถ้าจำนวนนั้นเป็นเลขคี่ เช่น half(1) ให้ส่งค่า (0,false)และ half(2) ควรจะส่งค่า (1,true)
*/
/*
package main

func main() {

}


*/

/* โจทย์ข้อที่ 3
จงเขียนฟังก์ชันที่รับค่าพารามิ้ตอร์ 1 ค่าแต่เป็นแบบไม่จำกัด(variadic parameter) เพื่อหาค่าที่มากที่สุดของลิสต์ตัวเลขที่ส่งมา
*/
package main

import "fmt"

func maxNumber(args ...int) int {
	total := 0
	for _, n := range args {
		total = n
	}
	return total
}
func main() {
	fmt.Println()
}
