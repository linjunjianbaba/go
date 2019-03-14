package main

import "fmt"

//所谓闭包就是一个函数“捕获”了和它在同一个作用域的其他常量和变量。这意味着当闭包被调用的时候
//不管程序在什么地方调用，闭包能够使用这些常量或者变量。它不关心这些捕获了的变量和常量是否已超
//出了作用域，所以只有闭包还在使用它，这些变量就还会存在。

//函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int {
	var x int //没有初始值
	return func() int {
		x++
		return x * x
	}
}

func test01() int {
	//函数被调用时，x才分配空间，才初始化为0
	var x int //没有初始化，值为0
	x++
	return x * x //函数调用完毕，X自动释放
}

func main() {
	//返回值为一个匿名函数，返回一个函数类型，通过f来调用闭包函数
	//它不关心这些捕获了的变量和常量是否已经超出了作用域
	//所以只有闭包还在使用它，这些变量就还会存在

	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
