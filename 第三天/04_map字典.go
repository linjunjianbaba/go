package main

import "fmt"

func main() {
	dict := make(map[int]string)
	dict[1] = "id"
	dict[2] = "name"
	fmt.Println(dict)
}
