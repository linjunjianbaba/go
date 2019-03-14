package main

import "fmt"

//无参有返回值，只用一个返回值
//有返回值的函数需要通过return中断函数，通过return返回
func MyFunc1() int {
	return 6666
}

//给返回值起一个变量名，go推荐写法
func MyFunc2() (result int) {
	return 7777
}

//给返回值起一个变量名，go推荐写法
//这个是最常用写法
func MyFunc3() (result int) {
	result = 8888888
	return
}
func main() {
	//无参有返回值函数调用
	var a int
	a = MyFunc1()
	fmt.Println("a = ", a)
	b := MyFunc2()
	fmt.Println("b = ", b)

	c := MyFunc3()
	fmt.Println("c = ", c)
}
