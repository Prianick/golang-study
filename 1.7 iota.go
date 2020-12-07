package main

import (
	"fmt"
)


const (
	a = iota + 1
	_
	b
	c
	d = c + 2
	t
	i
	i2 = iota + 2
)

func main() {
	fmt.Print("c")
	fmt.Println(c)
	fmt.Print("b")
	fmt.Println(b)
	fmt.Print("d")
	fmt.Println(d)
	fmt.Print("t")
	fmt.Println(t)
	fmt.Print("i")
	fmt.Println(i)
}