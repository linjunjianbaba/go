package main

import "fmt"

func MyFunc01(a int, b int) { //固定参数

}

//...int类型这样的类型，...type不定参数类型
//
func MyFunc2(args ...int) { //传递的实参可以是0或多个
	fmt.Println("len(arges) =", len(args)) //获取用户传递参数个数
}

func main() {
	MyFunc2(0)
	MyFunc2(1, 2, 3)
	MyFunc2(2, 5, 6, 7)
}

//固定参数一定要传参，不定参数根据需求传递
func MyFunc3(a int, args ...int) {

}

//注意：不定参数，一定（只能）放在形参中的最后一个参数
//func Myfunc4(args ...int, a int){
// 这是错误的写法
//}
