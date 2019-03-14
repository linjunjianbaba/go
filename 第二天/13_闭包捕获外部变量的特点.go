package main

import "fmt"

func main() {
	a := 10
	str := "lin"

	func() {
		//闭包已引用方式捕获外部变量
		a = 8888
		str = "jun"
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() //()代表直接调用
	fmt.Printf("外部：a = %d, str = %s\n", a, str)
}
