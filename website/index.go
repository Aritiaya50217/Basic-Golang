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

	http.ListenAndServe(":8080", nil)

}
func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Resquest")
}
func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Request")
}
