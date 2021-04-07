package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	// การตรวจสอบว่า test ขึ้นต้นด้วย te หรือไม่
	p("HasPerfix: ", s.HasPrefix("test", "te"))
	// การตรวจสอบว่า test ลงท้ายด้วย e หรือไม่
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	// ค้นหาข้อมูลว่า e อยู่ตำแหน่งไหน
	p("Index: ", s.Index("test", "e"))
	// การนำ string มาต่อกัน แล้วเก็บใน [] โดยมี - คั่น
	p("Join : ", s.Join([]string{"a", "b"}, "-"))
	// ทำซ้ำ
	p("Repeat: ", s.Repeat("a", 5))
	// แทนตำแหน่ง ตัว o เป็นเลข 0
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	// แทนตำแหน่ง ตัว o เป็นเลข 0 แค่ 2 ตัว
	p("Replace: ", s.Replace("fooo", "o", "0", 2))
	// การแยกช่วงของข้อมูล ระหว่าง a-b-c-d-e เป็น a b c d e แล้วเก็บใน array[]
	p("Spilt: ", s.Split("a-b-c-d-e", "-"))
	// เปลี่ยนพิมพ์เล็กเป็นพิมพ์ใหญ่
	p("ToLower: ", s.ToLower("TEST"))
	// เปลี่ยนพิมพ์ใหญ่เป็นพิมพ์เล็ก
	p("ToUpper : ", s.ToUpper("test"))
	p("Len: ", len("Hello"))

}
