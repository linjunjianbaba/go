package main

import "fmt"

//实现1+2+3....+100

func test01() (sum int) {
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return
}

func test02(i int) int {
	if i == 1 {
		return 1
	}
	return i + test02(i-1) //100 + 99 + 88
}

func test03(i int) int {
	if i == 100 {
		return 100
	}
	return i * test03(i+1)
}

func main() {
	sum := test01()
	fmt.Println("sum =", sum)

	max := test03(1)
	fmt.Println("max = ", max)
}
