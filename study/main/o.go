package main

import "fmt"

func test1() {
	fmt.Println(str7)
}

var str7 = "qweqw"

func main() {
	var str1 = "hhh"
	var str2 = "hhwww"
	var str3 = "Code = %s, DD = %s"
	var str4 = fmt.Sprintf(str3, str1, str2)
	fmt.Println(str4)
	var int1 int = 1
	fmt.Println(int1)
	var str5 string
	fmt.Println(str5)
	str6 := "ww"
	fmt.Println(str6)
	test1()
	str7 = "qweqwsss"
	fmt.Println(str7)
	var int2 int
	int2 = 1
	fmt.Println(int2)
	var int3, str8 = 6, '2'
	var int4 = int3 * int2
	fmt.Println(int3, str8, int4)
	var str9 = "wqeqweqweqwewqeq"
	fmt.Println(len(str9))
}
