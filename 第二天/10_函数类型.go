package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

//函数也是以种类型，通过type给一个函数类型起名
//FuncType它是一个函数类型
type FuncType func(int, int) int //没有函数名，没有{}

func main() {
	var result int
	result = Add(3, 4) //传统调用方式
	// result := Add(2, 2)
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名问fTest
	var fTest FuncType
	fTest = Add //是变量就可以赋值

	result = fTest(10, 90) //等价于Add（10,90）
	fmt.Println("result1 = ", result)
}
