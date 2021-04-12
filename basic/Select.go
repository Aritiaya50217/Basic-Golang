package main

import (
	"fmt"
	"time"
)

// go มี statement พิเ ศษชื่อว่า selectทำงานคล้ายกับ switch แต่ใช้กับ channel

func main() {
	// channel จะทำงานแบบ synchronize
	c1 := make(chan string)
	c2 := make(chan string)
	/* buffered channel คือ การส่งค่าพารามิเตอร์ตัวที่ 2 เข้าไปใน func make
	buffered เป็น Asynchonousคือรันได้เลยไม่ต้องรอ
	c3 := make(chan int, 1) จะได้ buffered channel ที่มีความจุเป็น 1
	*/
	c3 := make(chan int, 1)
	go func() {
		for {
			c1 <- "from 1"
			// พิมพ์ from 1 ทุก 2 วินาที
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			// พิมพ์ from 2 ทุก 3 วินาที
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			c3 <- 00
			time.Sleep(time.Second * 4)
		}
	}()
	go func() {
		for {
			select {
			// ส่งค่าที่เก็บใน channel c1 ไปยัง msg1
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case msg3 := <-c3:
				fmt.Println("Channel : ", msg3)
			// select ถูกใช้บ่อยในการทำ timeout
			case <-time.After(time.Second):
				fmt.Println("timeout")
			/* time.After จะสร้าง channell ขึ้นมาและส่งไปตอนที่ครบกำหนดเวลา
			สามารถระบุกรณี default ได้ด้วย
			กรณี default จะเกิดทันทีเมื่อไม่มี channel ใดพร้อม */
			default:
				fmt.Println("nothing ready")
			}
		}
	}()
	// เอาค่าทั้งหมดมาแสดง
	var input string
	fmt.Scanln(&input)

}
