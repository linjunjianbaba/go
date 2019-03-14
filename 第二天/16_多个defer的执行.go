package main

import "fmt"

func test(x int) {
	s := 100 / x //x为0时产生异常
	fmt.Println("s = ", s)
}

func main() {
	//defer延迟调用，main函数结束前调用
	defer fmt.Println("aaaaaaaaaaa")

	defer test(0)

	defer fmt.Println("bbbbbbbbbbbb")
	defer fmt.Println("cccccccccccccccc")
	//如果一个函数有多个defer语句，他们会以LIFO（后进先出）的顺序执行，哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行
}

/*
输出结果
cccccccccccccccc
bbbbbbbbbbbb
aaaaaaaaaaa
panic: runtime error: integer divide by zero
*/
