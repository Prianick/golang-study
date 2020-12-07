package main

import "fmt"

func main() {
	var n int
	max := 0
	positiveCount := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if max < arr[i] {
			max = arr[i]
		}
	}
	for index, elm := range arr {
		if index%2 == 0 {
			fmt.Print(elm, " ")
		}
		if elm > 0 {
			positiveCount++
		}
	}
	fmt.Println(positiveCount)
}

func main1() {
	var n int
	max := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Print(max)
}
