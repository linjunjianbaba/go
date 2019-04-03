package main

import "fmt"

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用一个接口，不同表现，可以实现不同表现，加减乘除
func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	//result = Add（a，b）//Add（）必须定义后才能调用
	return
}
func main() {
	a := Calc(23, 24, Minus)
	fmt.Println("a = ", a)

}
