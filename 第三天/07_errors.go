package main

import (
	"fmt"
)

type error interface {
	Error() string
}
type errorString struct {
	text string
}

func New(text string) error { //var &errorString{text} err
	return &errorString{text}
}

func (e *errorString) Error() string { //var e *errorString
	return e.text
}

func main() {
	t := &errorString{"this is error"}
	var e error
	e = t
	fmt.Println(e.Error())
}
