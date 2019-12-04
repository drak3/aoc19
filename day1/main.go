package main

import (
	"fmt"

	"github.com/drak3/aoc19/file"
)

func main() {
	in := file.ReadInts("day1/in.txt")
	mass := 0
	for _, i := range in {
		mass += i/3 - 2
		rem := i/3 - 2
		for rem != 0 {
			rem = rem/3 - 2
			if rem < 0 {
				rem = 0
			}
			mass += rem
		}
	}

	fmt.Println(mass)
}
