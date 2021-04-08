package main

import (
	"fmt"
	"net/http"
	"time"
)

// สร้างฐานข้อมูล
type Cookie struct {
	Name  string
	Value string
	// ใช้ time ที่ import มา
	Expires time.Time
}

func main() {
	http.HandleFunc("/", index)

	// port
	http.ListenAndServe(":8000", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	// ดึงเวลามาเก็บในตัวแปร expiration
	// กำหนดเวลาหมดอายุ คือ 1 ปี
	expiration := time.Now().Add(time.Hour * 24 * 365)
	// cookie เก็บข้อมูลคุกกี้จาก struct Cookie
	// Expires จะหมดเวลา ตามตัวแปร expiration ที่เรากำหนด
	cookie := http.Cookie{Name: "user", Value: "Artitaya Yaemjaraen", Expires: expiration}
	// เก็บคุ้กกี้
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Create Cookie")
}
