package main

import "fmt"

func main() {
	//goto可以用在任何地方，但是不能跨函数使用
	fmt.Println("1111111111")
	fmt.Println("2222222222")
	goto Biao //goto关键字，Biao用户起的名字，也可以叫标签

	fmt.Println("2222222222") //使用了goto此行不运行
Biao:
	fmt.Println("标标标标标标")
}
