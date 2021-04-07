package main

import "fmt"

// structure
// create structure book
// ต้องระบุ type

type Books struct {
	// ประกอบด้วย filed
	title    string
	author   string
	subtitle string
	price    float64
}

func main() {

	// เรียกใช้งาน structure
	// Books คือ type ตามที่ระบุด้านบน
	// Book1 ตัวแปรเก็บข้อมูลของ Books
	var Book1 Books
	Book1.title = "Go Programming"
	Book1.author = "Artisss"
	Book1.subtitle = "Golang language"
	Book1.price = 199
	fmt.Println(Book1)
	fmt.Println(Book1.title)

	// แบบที่ 2
	golang := Books{title: "Go Programming", author: "Artisss", subtitle: "Golang language", price: 199}

	fmt.Println(golang)
}
