package main

import (
	"fmt"
	"net/http"
)

func main() {
	// หน้าแรก เรียกใช้งาน func index
	http.HandleFunc("/", index)
	// หน้า login
	http.HandleFunc("/login", login)
	// create port
	http.ListenAndServe(":8000", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method : ", r.Method)
	r.ParseForm()
	// r.Form["username"] คือ แสดง username ที่เรากรอกในฟอร์ม login.html แต่ไม่แสดงหน้าเพจ
	fmt.Println("UserName: ", r.Form["username"])
	fmt.Println("Password: ", r.Form["password"])
}
