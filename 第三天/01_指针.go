package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("&a = %p\n", &a)
	var p *int
	p = &a
	fmt.Printf("p = %p\n", p)
	fmt.Printf("a = %d, *p = %d", a, *p)

	*p = 111
	fmt.Printf("a = %d, *p = %d", a, *p)
}
