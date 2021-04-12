package main

import (
	"html/template"
	"net/http"
)

// create stucture เก็บข้อมูล
type Product struct {
	Name  string
	Price int
}

func main() {
	// สร้างตัวแปร var ให้เก็บค่า template ไปแสดงยัง index
	var templates = template.Must(template.ParseFiles("index.html"))
	// แสดงหน้าแรก
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		// ตัวแปร myProduct รับข้อมูลมาจาก Product struct
		myProduct := Product{"นมสด", 500}
		// การใช้งาน template
		templates.ExecuteTemplate(w, "index.html", myProduct)
	})

	// แสดงหน้า login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	// แสดงหน้า signup
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "signup.html")
	})

	// แสดงไฟล์ txt ในหน้าเว็บ
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "file.txt")
	})

	// port 8080
	http.ListenAndServe(":8080", nil)
}
