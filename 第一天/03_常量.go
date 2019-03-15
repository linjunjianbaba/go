package main

import "fmt"

func main() {
	//变量：程序运行期间，可以改变的量，变量声明需要var
	//常量：程序运行期间，不可以改变的量，常量声明需要const
	const a int = 10
	//a = 20 会出现错误，常量不允许修改
	fmt.Printf("a = %d\n", a)

	const b = 11.2 //不能使用:=
	fmt.Printf("b type is %T\n", b)
	fmt.Printf("b = %v", b)
}

/*
a = 10
b type is float64
b = 11.2
*/
