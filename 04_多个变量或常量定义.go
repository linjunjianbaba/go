package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）
	// var a int =1
	// var b float64 = 2.0
	var ( //于上面的定义方式相同
		a int     = 1
		b float64 = 3.14
	)
	a, b = 10, 2.99
	fmt.Println("a = ", a)
	fmt.Println("b =", b)
	//选中代码，按ctrl+/ ,注释和取消注释代码的快捷键
	// const i int = 10
	// const j float64 = 4.8
	// const (
	// 	i int     = 10
	// 	j float64 = 8.88
	// )
	//可以自动推导数据类型
	const (
		i = 10
		j = 1.33
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}
