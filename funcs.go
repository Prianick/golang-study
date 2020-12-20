package main

import "fmt"

func vote(x int, y int, z int) int {
	arr := [3]int{x, y, z}
	var positiveVotes, negativeVotes int

	for _, elem := range arr {
		if elem == 1 {
			positiveVotes++
		} else {
			negativeVotes++
		}
	}

	if positiveVotes > negativeVotes {
		return 1
	}

	return 0
}

func fibonacci(n int) int {
	var i int
	f := 1
	var numbers [11]int
	for i = 0; i <= n; i++ {
		if i == 0 || i == 1 {
			numbers[i] = f
		} else {
			numbers[i] = numbers[i-2] + numbers[i-1]
		}
	}
	return numbers[n-1]
}

func sumInt(a ...int) (int, int) {
	count := 0
	sum := 0
	for _, elem := range a {
		count++
		sum = sum + elem
	}
	return count, sum
}

func main() {
	// fmt.Println(vote(0,0,1))
	// fmt.Println(fibonacci(10))
	a, b := sumInt(1, 0, 3, 4)
	fmt.Print(a, b)

}
