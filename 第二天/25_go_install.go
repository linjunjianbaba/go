package main

import io "fmt"

// func main() {
// 	yuwen := 98
// 	shuxue := 99
// 	yingyu := 80
// 	io.Printf("总分是： %d, 平均分是%f\n", yuwen+shuxue+yingyu, float64(yingyu+shuxue+yingyu)/3)
// 	num1 := 35
// 	num2 := 40
// 	// num3 := 2
// 	// var avg float64
// 	// avg := (num1 + num2) / num3
// 	io.Printf("avg = %f", float64(num1+num2)/2)
// }
// func main() {
// 	io.Println("请输入考试成绩：")
// 	var score int
// 	io.Scanf("%d", &score)
// 	for i := 1; i <= score; i++ {
// 		io.Printf("%d\n", i)
// 	}
// }
// func main() {
// 	var sum int
// 	for i := 1; i <= 100; i++ {
// 		sum = sum + i
// 	}
// 	io.Println("sum = ", sum)

// }
// func main() {
// 	var b int
// 	var s int
// 	var g int
// 	for i := 100; i <= 999; i++ {
// 		b = i / 100
// 		s = i % 100 / 10
// 		g = i % 10
// 		if b*b*b+s*s*s+g*g*g == i {
// 			io.Println("水仙数有", i)
// 		}

// 	}
// }
// func main() {
// 	var (
// 		userName string
// 		userPwd  string
// 		count    int
// 	)
// 	for {
// 		io.Println("请输入用户名:")
// 		io.Scanf("%s\n", &userName)
// 		io.Println("请输入密码:")
// 		io.Scanf("%s\n", &userPwd)
// 		if userName == "admin" && userPwd == "888888" {
// 			io.Println("登陆成功")
// 			break
// 		} else {
// 			count++
// 			if count >= 3 {
// 				io.Println("你输入的次数过多")
// 				break
// 			}
// 			io.Println("用户名密码输入错误，请重新输入！！")
// 		}
// 	}
// }
// func main() {
// 	var i int
// 	var j int

// 	for i = 1; i <= 9; i++ {
// 		for j = 1; j <= i; j++ {
// 			io.Printf("%d*%d=%d\t", i, j, i*j)
// 		}
// 		io.Println("")
// 	}
// }

func test(a, b int) (sum, max int) {
	sum = a + b
	max = b - a
	return
}

type funcType func(a, b int) (d, c int)

func main() {
	var s int
	var result funcType

	result = test
	s, _ = result(10, 100)
	io.Println(s)
}
