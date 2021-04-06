package main

import "fmt"

func main() {
	// function
	map1()
	map2()
}
func map1() {
	// map คล้าย dict ใน python
	x := make(map[string]string)
	x["TH"] = "Thailand"
	x["JP"] = "Japan"
	x["EN"] = "England"

	fmt.Println(x["TH"])
}
func map2() {
	y := map[string]string{
		"TH": "Thailand",
		"JP": "Japan",
	}
	fmt.Println(y)
}
