package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error){
	var err error
	err = nil
	if b == 0 {
		err = errors.New("my error")
	}
	return a/b, err
}

func divideTest() {
	var a int
	var b int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil || b == 0 {
		fmt.Println("ошибка")
		return
	}
	res, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(res)
	}
}

func main() {
	//err := errors.New("my error")
	//fmt.Println("test", err)
	//var a int
	//_, err := fmt.Scan(&a)
	//if err != nil {
	//	fmt.Println(a + 1)
	//} else {
	//	fmt.Println(a + 2)
	//}

	divideTest()
}
