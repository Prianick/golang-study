package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

/*
	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n){
		fmt.Println(n)
	}vsfjk j j
*/

func isPalindrom() bool {
	rs := getStr()
	lastIndex := len(rs) - 1
	if lastIndex < 0 {
		return false
	}
	for i := range rs {
		if i == lastIndex-i || i > lastIndex-i {
			break
		}
		if rs[i] != rs[lastIndex-i] {
			return false
		}
	}
	return true
}

func getStr() []rune {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	return []rune(text)
}

func strpos() int {
	x, s := "", ""
	fmt.Scan(&x)
	fmt.Scan(&s)
	haystack := []rune(x)
	needle := []rune(s)
	match := false
	position := 0
	lastPoint := 0
	for i := 0; i < len(haystack); i++ {
		for j := lastPoint; j < len(needle); j++ {
			k := i + j
			if needle[j] != haystack[k] {
				match = false
				lastPoint = 0
				break
			} else {
				match = true
				if j+1 == len(needle) {
					position = i
					i = len(haystack)
				}
			}
		}
	}

	if match {
		return position
	} else {
		return -1
	}
}
func strpos2() {
	x, s := "", ""
	fmt.Scan(&x)
	fmt.Scan(&s)
	res := strings.Index(x, s)
	fmt.Println(res)
}

func isRightStr() {
	rs := getStr()
	len := len(rs)
	if rs[len-1] == '.' && unicode.IsUpper(rs[0]) {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}

func removeOddSymbols() {
	rs := getStr()
	newStr := []rune("")
	for i := range rs {
		if (i+1)%2 != 0 {
			newStr = append(newStr, rs[i])
		}
	}
	fmt.Printf("%s", string(newStr))
}

func removeADuplicateCharacters() {
	rs := getStr()
	var next rune
	var prev rune
	cleanStr := []rune("")
	for i := range rs {
		if i-1 >= 0 {
			prev = rs[i-1]
			if prev == rs[i] {
				continue
			}
		}
		if i+1 < len(rs) {
			next = rs[i+1]
			if next == rs[i] {
				continue
			}
		}
		cleanStr = append(cleanStr, rs[i])
	}
	fmt.Printf("%s", string(cleanStr))
}

func leaveOnlyUniqueChars() {
	rs := getStr()
	cleanStr := []rune("")
	isUniq := true
	for i := range rs {
		isUniq = true
		for j := range rs {
			if rs[i] == rs[j] && i != j {
				isUniq = false
			}
		}
		if isUniq == true {
			cleanStr = append(cleanStr, rs[i])
		}
	}
	fmt.Printf("%s", string(cleanStr))
}

func isPswdValid() bool {
	rs := getStr()
	if len(rs) < 5 {
		return false
	}
	for i := range rs {
		if !unicode.Is(unicode.Latin, rs[i]) && !unicode.IsNumber(rs[i]) {
			return false
		}
	}
	return true
}

func isPswdValid2() {
	var (
		src string
		rex *regexp.Regexp
	)

	rex = regexp.MustCompile(`(?m)^[0-9a-zA-Z]{5,}$`)
	fmt.Scan(&src)
	if rex.MatchString(src) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func main() {
	//res := isPalindrom()
	//if res {
	//	fmt.Print("Палиндром")
	//} else {
	//	fmt.Print("Нет")
	//}
	//fmt.Print(strpos())
	//removeOddSymbols()

	//leaveOnlyUniqueChars()

	//if isPswdValid() {
	//	fmt.Print("Ok")
	//} else {
	//	fmt.Print("Wrong password")
	//}

	isPswdValid2()
}
