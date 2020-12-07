package main

import "fmt"

func main() {
	var workArray [10]uint8
	var n uint8
	var i uint8
	for i = 0; i < 10; i++ {
		fmt.Scan(&n)
		workArray[i] = n
	}
	var n2 uint8
	var tmp uint8
	for i = 0; i < 3; i++ {
		fmt.Scan(&n)
		fmt.Scan(&n2)
		tmp = workArray[n2]
		workArray[n2] = workArray[n]
		workArray[n] = tmp
	}

	for _, elem := range workArray{
		fmt.Print(elem)
		fmt.Print(" ")
	}
}
