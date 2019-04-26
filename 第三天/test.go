package main

import (
	"fmt"
)

// import "fmt"

type Str struct {
	text string
}

type err interface {
	Add() string
}

func New(a string) err {
	return &Str{a}
}

func (w *Str) Add() string {
	return w.text
}

func main() {
	t := "this is error"
	var i err
	s := &Str{"ni 是谁"}
	i = s
	fmt.Println(i.Add())
	fmt.Println(New(t).Add())
}
