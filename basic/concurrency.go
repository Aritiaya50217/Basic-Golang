package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* goroutines เป็นฟังก์ชันที่สามารถทำงานได้ในเวลาเดียวกันกับฟังก์ชันอื่นได้
วิธีการสร้าง goroutines คือใส่ go นำหน้าการเรียกใช้งานฟังก์ชัน
*/
func f(index string) {
	for i := 0; i < 10; i++ {
		fmt.Println(index, ":", i+1)
	}
}

// หน่วงเวลาในการรัน goroutines ด้วย time.Sleep and rand.Intn

func testTime(n int) {
	/* พิมพ์ตัวเลขจาก 0 ถึง 10 แล้วรอ 0-250 มิลลิวินาทีในแต่ละรอบ
	 */
	for m := 0; m < 10; m++ {
		fmt.Println(n, ":", m)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// goroutines
	/*
		go f("index")
		var input string
		fmt.Scanln(&input)
	*/

	// รัน 10 routines
	/*
		\for x := 0; x < 10; x++ {
			go f("x")

		}
		var input string
		fmt.Scanln(&input)
	*/

	// หน่วงเวลาในการรัน goroutines
	for m := 0; m < 10; m++ {
		go testTime(m)
	}
	var input string
	fmt.Scanln(&input)
}
