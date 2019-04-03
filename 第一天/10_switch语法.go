package main

import "fmt"

// func main() {
// 	var num int
// 	fmt.Printf("请按楼层：")
// 	fmt.Scan(&num)

// 	switch num { //switch后面写的是变量的本身
// 	case 1:
// 		fmt.Println("按下的是1楼")
// 	case 2:
// 		fmt.Println("按下的是2楼")
// 	case 3:
// 		fmt.Println("按下的是3楼")
// 	case 4:
// 		fmt.Println("按下的是4楼")
// 	default:
// 		fmt.Println("请按下的楼层")
// 	}

// }

func main() {
	//支持一个初始化语句，初始化语句变量本身，用分号分隔
	switch num := 4; num {

	case 1:
		fmt.Println("按下的是1楼")
		//break //go语言保留了break关键字，跳出switch语句，不写，默认就包含
		fallthrough //不跳出switch语句，后面无条件执行
	case 2:
		fmt.Println("按下的是2楼")
	case 3, 4, 5:
		fmt.Println("按下的是yyyy楼") //可以写多个
	case 6:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("请按下的楼层")
	}
	soure := 85
	switch { //可以没有条件
	case soure > 90:
		fmt.Println("优秀")
	case soure > 80:
		fmt.Println("良好")
	case soure > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其他")
	}

}
