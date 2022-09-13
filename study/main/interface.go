package main

import "fmt"

type Phone interface {
	call()
}

// struct_name

type Iphone struct {
}

// struct_name_variable struct_name

func (iphone Iphone) call() {
	fmt.Println("Hello I'm iphone")
}

func main() {
	var phone Phone
	phone = new(Iphone)
	phone.call()
}
