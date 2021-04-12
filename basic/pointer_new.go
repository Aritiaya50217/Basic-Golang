package main

import "fmt"

func one(xPrt *int) {
	*xPrt = 1
}
func main() {
	/* new รับ type เป็นอาร์กิวเมนท์และทำการจองหน่วยความจำให้พอดีกับค่าของข้อมูลประเภทนั้นๆ และส่ง pointer กลับออกมา
	 */
	xPrt := new(int)
	one(xPrt)
	fmt.Println(*xPrt)
}
