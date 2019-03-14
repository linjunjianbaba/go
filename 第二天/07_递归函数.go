package main

import "fmt"

func MyFunc1(a int) {
	if a == 1 {
		fmt.Println("a = ", a)
		return
	}
	//函数调用自身
	MyFunc1(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	MyFunc1(3)

	fmt.Println("main")
}
