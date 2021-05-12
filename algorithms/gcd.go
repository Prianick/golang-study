package main

import "fmt"

func main() {
	var a, b uint64
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a == 0 {
		fmt.Printf("%d", a)
	}
	if b == 0 {
		fmt.Printf("%d", b)
	}
	for ; a != 0 && b !=0;  {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a == 0 {
		fmt.Printf("%d", b)
	} else {
		fmt.Printf("%d", a)
	}
}
