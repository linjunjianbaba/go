package main

import "fmt"

// type xs struct {
// 	name string
// 	sex  byte
// 	age  int
// }

// func (p xs) printInfo() {
// 	fmt.Println(p.name, p.sex, p.age)
// }

// func main() {
// 	p := xs{"jun", 'm', 1}
// 	p.printInfo()
// }
type person struct {
	name string
	sex  byte
	age  int
}

type student struct {
	person
	id   int
	addr string
}

func main() {
	//顺序初始化
	s := student{person{"jun", 'm', 18}, 1, "gz"}
	fmt.Printf("s = %+v\n", s)
	//部分成员初始化
	d := student{person: person{"biao", 'n', 25}, id: 2}
	fmt.Printf("s = %+v\n", d)
	//成员操作
	var f student
	f.name = "jing"
	f.sex = 'm'
	f.age = 25
	f.id = 3
	f.addr = "gz"
	fmt.Println(f)
}
