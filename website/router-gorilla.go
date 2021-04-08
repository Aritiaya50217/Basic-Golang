package main

import (
	"fmt"
	"net/http"

	/* import libary from github */
	"github.com/gorilla/mux"
)

func main() {

	userDB := map[string]int{
		"Artitaya": 20,
		"oil":      23,
		"A":        12,
	}

	// สร้าง router
	router := mux.NewRouter()
	// เปลี่ยนจาก http.HandleFunc เป็น router.HandleFunc
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	// ส่งพารามิเตอร์เข้ามาผ่านทาง url => ("/user/{name}")
	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Response) {
		// ใช้ method get ในการรับค่า แล้วเอามาเก็บในตัวแปร vars
		// r = request
		vars := mux.Vars(r)
		// name = params
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "My name is %s %d years old", name, age)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
