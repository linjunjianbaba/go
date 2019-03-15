package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}
func main() {
	//赋值前必须声明变量
	var a int
	a = 10 //赋值，可以使用N次
	a = 30
	fmt.Println("a =", a)
	//:= 自动推导类型，先声明变量b，再给b赋值20
	//自动推导，同一个变量名只能使用一次，用于初始化那次
	b := 20
	fmt.Println("b = ", b)
	//b := 40
	//b :=40 前面已经有变量b，不能再新建一个变量b
	b = 40                 //只是赋值是可以的
	fmt.Println("b = ", b) //Println()可以自动换行Print（）需要添加\n换行
	//_匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势

	var e, f, g int
	e, f, g = test() //ruturn 1，2，3
	fmt.Printf("e = %d, f = %d, g = %d\n", e, f, g)
	_, f, _ = test()
	fmt.Printf("f = %d\n", f)
}

/*
a = 30
b =  20
b =  40
e = 1, f = 2, g = 3
f = 2
*/
