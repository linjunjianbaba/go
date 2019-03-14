package main

import "fmt"

func MyFunc1(args ...int) {
	//全部元素传递给其他函数
	MyFunc2(args...)

	//只想把后个参数传递给另外一个函数使用
	MyFunc2(args[:2]...) //args[0]-args[2](不包括args[2]本身)，传递过去
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")
	MyFunc2(args[2:]...) //从args[2]开始（包括本身），把后面所用元素传递过去
}

func MyFunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func main() {
	MyFunc1(1, 2, 3, 4, 5, 6, 7)
}
