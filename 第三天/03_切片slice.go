package main

import "fmt"

func main() {
	array := []int{10, 20, 30, 0, 0}

	s := array[0:3:5]
	fmt.Println("s =", s)
}
