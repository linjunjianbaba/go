package main

import "fmt"

//无参无返回值函数定义//
// func sum() {
// 	num := 0
// 	for i := 0; i <= 100; i++ {
// 		num = num + i
// 	}
// 	fmt.Println("num = ", num)
// }

// func main() {
// 	//无参无返回值函数的调用： 函数名（）
// 	// sum()
// }

//有参无返回值函数定义，普通参数列表//
//定义函数时，在函数名后面（）定义的参数叫形参
//参数传递，只能由实参传递给形参，不能反过来，单向传递
//1
func MyFunc1(a int) {
	fmt.Println("a =", a)
}

//2
func MyFunc2(a int, b int) { //等同于a, b int
	fmt.Printf("a =% d, b = %d\n", a, b)
}
func main() {
	//有参无返回值函数调用：函数名（所需参数）
	//调用函数传递的参数叫实参
	MyFunc1(6666)
	MyFunc2(7777, 8888)
}
