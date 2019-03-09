package main

import "fmt"

func main() {
	//浮点类型
	var f float32
	f = 3.156
	fmt.Println("f =", f)

	f1 := 3.1456
	fmt.Printf("f1 is type %T\n", f1) //f1 type is float64
	//float64存储小数比float32更准确

	//字符类型//
	var ch byte //声明字符类型
	ch = 97
	fmt.Println("ch = ", ch)
	//格式化输出，%c已字符方式打印，%d已挣型方式打印
	fmt.Printf("%c, %d\n", ch, ch)
	ch = 'a' //字符，单引号
	fmt.Printf("%c, %d\n", ch, ch)

	//大写转小写，小写转大写，大小相差32，小写大
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写: %c\n", 'A'+32)
	fmt.Printf("小写转大写: %c\n", 'a'-32)
	//\n代表换行

	//字符串类型//
	var str1 string
	str1 = "abc"
	fmt.Println("str1 =", str1)

	//自动推导类型
	str2 := "mike"
	fmt.Printf("str2类型是: %T\n", str2)

	//内建函数，len()可以测字符串大的长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))

	//字符和字符串的区别//
	//字符 1.单引号 2.字符，往往只有一个字符，转义字符除外'\n'

	var zh byte
	var str string
	zh = 'a'
	fmt.Println("zh = ", zh)
	//字符串 1.双引号 2.字符串有1个或多个字符组成，3.字符串都是隐藏了一个结束符，'\0'
	str = "a" //有'a'和'\0'组成一个字符串
	fmt.Println("str = ", str)

	str = "hello go"
	//操作字符串中的某个字符，从0开始操作
	fmt.Printf("str[0] = %c, str[2] = %c\n", str[0], str[2])

	//复数类型//
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 2.5 + 3.56i
	fmt.Printf("t2 is type %T\n", t2)

	//通过内部函数，取实部和虚部
	fmt.Println("real(t2) =", real(t2), ", imag(t2) =", imag(t2))

	//格式化//
	//%T操作变量所需类型 %d整形格式 %s字符串格式 %c字符格式 %f浮点格式 %v自动匹配格式输出智能不是太好

	//变量的输入//

	var e string
	fmt.Printf("请输入e:")
	//阻塞用户输入
	//fmt.Scanf("%d", &e) //别忘了&
	fmt.Scan(&e)
	fmt.Println("e =", e)

	//类型转换//
	//不能转换的类型，叫不兼容类型
	//bool不能转换为int
	//整形也不能转换为bool

	var x byte

	x = 'a' //字符类型本质上就是整型
	var y int
	y = int(x) //类型转换，把x的值取出来后，转成int再给t赋值
	fmt.Println("t =", y)

	// 类型别名 //
	// 给int64起一个别名bigint
	type bigint int64
	var a bigint //等价于var a int64
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)
	var b long = 1
	var c char = 'a'
	fmt.Printf("b = %d, c = %c\n", b, c)

	var xing string
	var mima string
	fmt.Printf("请输入姓名:")
	fmt.Scan(&xing)

	fmt.Printf("请输入密码:")
	fmt.Scan(&mima)
	fmt.Printf("姓名: %v ,密码: %v\n", xing, mima)

}
