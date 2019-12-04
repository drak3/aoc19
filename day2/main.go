package main

import (
	"fmt"
	"strconv"
	"strings"

	e "github.com/drak3/aoc19/errors"
	"github.com/drak3/aoc19/file"
	"github.com/drak3/aoc19/intprog"
)

func main() {
	in := file.Read("day2/in.txt")
	in = strings.Trim(in, "\n")
	strs := strings.Split(in, ",")
	nums := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(s)
		e.Ck(err)
		nums[i] = n
	}
	fmt.Println(nums)
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if r, err := intprog.Run(nums, i, j); err == nil && r[0] == 19690720 {
				fmt.Println(100*i + j)
				return
			}
		}
	}
	fmt.Println("not found")
}
