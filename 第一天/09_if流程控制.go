package main

import "fmt"

//一个分支//

// func main() {
// 	s := "林志标"
// 	//if和{就是条件，条件通常是关系运算符
// 	if s == "林志标" { //左括号和if在同一行
// 		fmt.Println("你是大神")
// 	}
// 	//if支持一个初始化语句，初始化语句和判断条件以分号分隔
// 	if a := 1; a == 1 { //	条件为真指向{}语句
// 		fmt.Println("a == 1")
// 	}
// }

//多分支//
func main() {
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else { //else后面没有条件
		fmt.Println("a != 10")
	}

	//2
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else { //else后面没有条件
		fmt.Println("a != 10")
	}

	//3
	if a := 20; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}
}
