package main

import "fmt"

func main() {
	m := make(map[float32]int)
	m[0] = 1
	m[1] = 2
	fmt.Println(m)
	//var s1 []map[int]string{}
	//s2 := []map[int]int[uint18]
	var s3 []map[int]int
	var s4 = []map[int]string{
		map[int]string{1: "string1", 2: "string2"},
	}
	var users [2]string
	//fmt.Println(s1)
	//fmt.Println(s2)
	fmt.Println(s3)
	s4[0] = make(map[int]string)
	s4[0][0] = "string"
	users[0] = "string"
	fmt.Println(s4)
	fmt.Println(users)

}
