package main

import "fmt"

//函数定义
func MaxMinFunc(a, b int) (Max, Min int) {
	if a > b {
		Max = a
		Min = b
	} else {
		Max = b
		Min = a
	}
	return //有返回值的函数，必须通过return返回
}

func main() {
	Max, Min := MaxMinFunc(20, 10)
	fmt.Printf("max = %d, min = %d\n", Max, Min)

	//通过匿名函数丢弃某个返回值
	a, _ := MaxMinFunc(88, 99)
	fmt.Printf("a = %d", a)
}
