package main

import (
	"fmt"
)

func main() {
	// ชุดข้อความ
	firstname := "Artitaya "
	lastname := "Yaemjaraen"
	sum := firstname + lastname

	fmt.Println(sum)
	// เข้าถึงข้อมูล
	fmt.Println(sum[1:3]) // ผลลัพธ์ rt = [1:3]

}
