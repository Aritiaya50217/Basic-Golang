package main

import "fmt"

func main() {
	// factorial(num int)
	fmt.Println(factorial(5))
}
func factorial(num int) int {
	// num-1 ไปเรื่อย ๆ จนมีค่าเท่ากับ 1
	// 5-1 = 4 => 4-1 =3 =>3-1=2 => 2-1 = 1
	if num == 1 {
		return 1
	}
	// 5 * ((5-1)*(4-1)*(3-1)*(2-1))
	// result 120
	return num * factorial((num - 1))
}
