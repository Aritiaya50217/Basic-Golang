package main

import "fmt"

func main() {
	for x := 1; x <= 10; x++ {
		if x%2 == 0 {
			fmt.Println(x, "Even")
		} else {
			fmt.Println(x, "ODD")
		}

	}

	test()
}
func test() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Println(x, "Fizz")
		} else if x%5 == 0 {
			fmt.Println(x, "Buzz")
		} else if x%3 == 0 && x%5 == 0 {
			fmt.Println(x, "FizzBuzz")
		}
	}
}
