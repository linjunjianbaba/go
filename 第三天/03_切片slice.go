package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := array[:6:8]
	fmt.Println(s)
}
