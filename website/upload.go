package main

import (
	"fmt"
	"io"
	"net/http"

	// os ใช้ในการเปิดไฟล์ , อัพโหลด
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", uploadHandle)
	// port
	http.ListenAndServe(":8000", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	// เรียกใช้งาน upload.html
	http.ServeFile(w, r, "upload.html")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method : ", r.Method)
	r.ParseForm()
	fmt.Println("Username : ", r.Form["username"])
	fmt.Println("Password : ", r.Form["password"])
}
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	// สร้างตัวแปรการรับไฟล์
	// file รับค่าไฟล์
	// handle รายละเอียด
	// err ค่าความผิดพลาด
	file, handle, err := r.FormFile("file")
	// อัพโหลดไม่สำเร็จให้สั่งปิด
	defer file.Close()
	//nil = ว่างเปล่า (null)
	if err != nil {
		fmt.Println(err)
		return
	}
	// เอาไฟล์ไปเก็บใน server
	fmt.Fprintf(w, "%v", handle.Header)
	// f กรณีเราอัพเดทข้อมูลสำเร็จ
	// err ข้อผิดพลาด
	// ./image/ คือ การอัพโหลดไปที่โฟลเดอร์ image
	// 0666 คือ status การอัพโหลด ถ้าอัพโหลดซ้ำจะเรียกข้อมูลได้อันเดียว
	f, err := os.OpenFile("./image/"+handle.Filename, os.O_CREATE, 0666)
	// กรณีเกิดข้อผิดพลาด
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// ทำการย้ายไฟล์ไปยังตำแหน่งปลายทาง โดยการ copy
	io.Copy(f, file)
	fmt.Fprintf(w, "Success Upload Complete")

}
