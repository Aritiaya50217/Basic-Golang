package main

import "fmt"

func main() {
	// map คือการเก็บข้อมูลแบบพจนานุกรม
	// x คือ map ของ string
	x := make(map[string]string)
	x["key"] = "oil"
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])

	// run function
	testMap()
}
func testMap() {
	//map ไม่มีการเรียงลำดับ ผลลัพธ์จะออกมาแบบใดก็ได้
	// ตัวแปร elements คือ map
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["Ne"] = "Neon"
	fmt.Println(elements)

	// ถ้าค่าที่เชื่อมโยงกับคีย์มีอยู่จะรันโค้ดที่อยู่ภายใน
	// name = ผลลัพธ์ของการค้นหา
	// ok = บ่งบอกว่าการค้นหาสำเร็จ
	if name, ok := elements["Ne"]; ok {
		fmt.Println(name, ok)
	}
}
func testMap2() {
	// ใช้ map เก็บข้อมูลทั่วไป
	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name": "Beryllium",
		},
	}
	fmt.Println(elements2)
}
