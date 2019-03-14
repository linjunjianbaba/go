package main  

//方法1
import "fmt"  //导入包，必须使用，否则使用不通过
import "os"
//方法2
import (
	fmt
	os
)

//给包名其别名
import io "fmt"  //io为别名，直接用io.Printf("1.%T\n", a)


var a byte

func main() {
	var a int //局部变量

	//1.不同作用域，允许定义同名变量
	//2.使用变量的原则，就近原则

	fmt.Printf("1.%T\n", a) //int
	{
		var a float64
		fmt.Printf("2.%T\n", a)  //float64
	test()
}

func test() {
	fmt.Printf("3.%T\n", a)  //uint8,就是byte
}
