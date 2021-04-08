package main

import (
	"fmt"
	"net/http"
)

func main() {
	// HandleFunc ถ้ามีการส่ง request มาจะให้ไปทำงานที่ function ไหน
	// w คือ ค่าที่ส่งมา Response
	// r คือ ส่งค่ากลับ request
	// หน้าแรก "/" แสดงคำว่า Hello My website
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ต้องใส่ w เมื่อเป็นการ response
		fmt.Fprintf(w, "Hello My website")

	})
	// หน้า /product จะไปเรียกใช้ funs product
	http.HandleFunc("/product", product)
	// หน้า /user จะไปเรียกใช้ funs user
	http.HandleFunc("/user", user)

	// ตัวอย่าง สร้าง map เก็บข้อมูล
	userDB := map[string]int{
		"Artitaya": 20,
		"oil":      23,
		"A":        12,
	}

	http.HandleFunc("/userdb/", func(w http.ResponseWriter, r *http.Request) {
		// รับ request มาเก็บในตัวแปร name
		name := r.URL.Path[len("/userdb/"):]
		// เก็บค่า userDB โดยการระบุผ่าน name
		age := userDB[name]
		// แสดงชื่อ และอายุ
		fmt.Fprintf(w, "My name is %s %d years old", name, age)

	})

	http.ListenAndServe(":8080", nil)

}
func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Resquest")
}
func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Request")
}
