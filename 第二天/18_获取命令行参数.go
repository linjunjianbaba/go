package main

import "fmt"
import "os"

func main() {
	//接收用户传递的参数，都是以字符串方式传递
	list := os.Args

	n := len(list)
	fmt.Printf("n = %d\n", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}
}

/*
n = 1
list[0] = C:\Users\bill\AppData\Local\Temp\go-build557886074\b001\exe\18_获取命令行参数.exe

*/
