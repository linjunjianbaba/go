package main

import "fmt"

func main() {
	fmt.Println("hello word")
	//自动推导类型：
	a := 5
	fmt.Printf("a type is %T\n", a)
}

/*
hello word
a type is int
*/