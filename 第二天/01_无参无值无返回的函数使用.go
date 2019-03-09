package main

import "fmt"

func sum() {
	num := 0
	for i := 0; i <= 100; i++ {
		num = num + i
	}
	fmt.Println("num = ", num)
}

func main() {
	sum()
}
