package main

import "fmt"

//多个返回值
func MyFunc1() (int, int, int) {
	return 1, 2, 3
}

//go官方推荐写法
func MyFunc2() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

func main() {
	//函数调用
	a, b, c := MyFunc2()
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}


函数定义说明：
 func：函数有关键字func开始声明
 FuncName：函数名称，根据约定，函数名首字母小写即为private（私有），大写即为public（公有）
 参数列表：函数可以有0或者多个参数，参数格式为：变量名 类型func MyFunc3(a int, args ...int)，
如果有多个参数通过逗号分隔，不支持默认参数
 返回类型：如果有返回值，那么必须在函数的内部添加return语句