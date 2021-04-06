package main

import "fmt"

// global ประกาศนอก function สามารถเรียกใช้งานได้ทุกฟังก์ชัน
var gVariable int = 500

// function หลัก
func main() {

	localVariable := 40
	fmt.Println("Global =", gVariable)
	fmt.Println("Local =", localVariable)
	// เรียกใช้งาน function
	anotherFunction()
}
func anotherFunction() {

	localVariable := 22
	fmt.Println("Global =", gVariable)
	fmt.Println("Local =", localVariable)

}
