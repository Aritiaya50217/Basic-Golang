package main

import (
	"fmt"
	"time"
)

/* channel ประกาศโดยใช้คำว่า chan
ประเภทของสิ่งที่จะส่งเข้าไปยัง channel
สามารถระบุได้ว่า channel นี้จะให้ทำเฉพาะรับหรือส่ง
func pinger(c chan<-string) คือ ทำได้แค่ส่งข้อความไปที่ c เท่านั้น
func pinger(c <-chan string) คือ channel ที่ไม่ได้ระบุทิศทาง
สามารถใช้ได้ทั้ง 2 ทิศทางทั้งรับและส่ง
*/
func pinger(c chan string) {
	for i := 0; ; i++ {
		// <- หมายถึง ส่งและรับข้อความบน channel
		// c <- "ping" คือ ส่ง ping เข้าไปใน channel
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		/* msg := <-c รับข้อความจาก channel และเก็บข้อความไว้ใน msg
		 */
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	go ponger(c)

	var input string
	fmt.Scanln(&input)
}
