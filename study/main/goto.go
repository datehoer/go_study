package main

import "fmt"

func main() {
	var s = 10
	for s < 20 {
		if s == 15 {
			s++
			goto OO
		}
		s++
		fmt.Println(s)
	}
OO:
	fmt.Println("ok")
}
