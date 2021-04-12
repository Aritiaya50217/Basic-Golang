package main

import "fmt"

// Embedded Types ไทป์แบบฝังตัว
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi,my name is", p.Name)
}

// สร้าง Android struct ที่มีความสัมพันธ์กับ Person
type Android struct {
	/*
		Person Person
		Medel string
	*/
	// ไทป์ฝังอีกแบบคือ
	// ใช้ไทป์ Person โดยไม่ประกาศชื่อ
	Person
	Model string
}

func main() {
	// เรียก Person struct ได้ด้วยการเรียกชื่อไทป์ได้เลย
	a := new(Android)
	a.Person.Talk()

}
