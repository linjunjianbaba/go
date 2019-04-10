package main

import "fmt"

type xs struct {
	name string
	sex  byte
	age  int
}

func (p xs) printInfo() {
	fmt.Println(p.name, p.sex, p.age)
}

func main() {
	p := xs{"jun", 'm', 1}
	p.printInfo()
}
