package main

import (
	"fmt"
	"strconv"
)

func main() {
	lower, upper := 246515, 739105
	c := 0
	fmt.Println(isValid2(112233), isValid2(123444), isValid2(111122))
	for i := lower; i <= upper; i++ {
		if isValid2(i) {
			c++
		}
	}
	fmt.Println(c)
}

func isValid(num []byte) bool {
	//check: there are two adjacent digits that are the same
	//check: digits are monotonically increasing
	prevDigit := num[0]
	adjacent := false
	for _, b := range num[1:] {
		if prevDigit == b {
			adjacent = true
		}
		if prevDigit > b {
			return false
		}
		prevDigit = b
	}
	return adjacent
}

func isValid2(n int) bool {
	//check: there are two adjacent digits that are the same
	//check: digits are monotonically increasing
	num := []byte(strconv.Itoa(n))
	prevDigit := num[0]
	adjacent := false
	c := 0
	//lowest := 100
	for _, b := range num[1:] {
		if b == prevDigit {
			c++
		} else {
			if c == 1 {
				adjacent = true
			}
			c = 0
		}
		if prevDigit > b {
			return false
		}
		prevDigit = b
	}
	return adjacent || c == 1
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
