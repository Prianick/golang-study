package main

import (
	"fmt"
)

func findFib1() {
	var n int
	fmt.Scan(&n)
	numbers := make([]uint64, n)
	numbers[0] = 0
	numbers[1] = 1
	for i := 2; i < n; i++ {
		numbers[i] = numbers[i-1] + numbers[i-2]
	}
	fmt.Print(numbers[n - 1])
}

func findFib2() {
	var n int
	fmt.Scan(&n)
	var current, prev1, prev2 uint64
	current = 0
	prev1 = 0
	prev2 = 1
	for i := 2; i <= n + 1; i++ {
		current = prev1%10 + prev2%10
		prev2 = prev1
		prev1 = current
	}
	fmt.Println(current%10)
}

func findModMN(){
	var n, m, i uint64
	fmt.Scan(&n)
	fmt.Scan(&m)
	var current, prev1, prev2 uint64
	current = 0
	prev1 = 0
	prev2 = 1
	for i = 2; i <= n + 1; i++ {
		current = (prev1%10 + prev2%10)%10
		prev2 = prev1
		prev1 = current
	}
	fmt.Println(current)
}

func main() {
	findModMN()
}

/*

100
3736710778780434371

400
2121778230729308891

500
2171430676560690477

1000
817770325994397771

2000
11629652114648610021


*/