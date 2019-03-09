package main

import "fmt"

func main() {
	//for 初始化条件；判断条件；条件变化 {
	//}
	//1+2+3...100累加
	// sum := 0
	//1.初始化条件 i： =1
	//2.判断条件是否为真，i<=100,如果为真执行循环体，如果为假跳出循环
	//3.条件变化 i++
	//4.重复2，3，4
	// for i := 1; i <= 100; i++ {
	// 	sum = sum + i
	// }
	// fmt.Println("sum =", sum)
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
	//迭代打印每个元素，默认返回两个值：一个元素的位置，一个是元素的本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	for i := range str { //第二个返回值，默认丢弃，返回元素的位置（下标）
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for i, _ := range str { //第二个返回值，默认丢弃，返回元素的位置（下标）
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
}
