package main

import (
	"fmt"
	"strconv"
	"strings"

	e "github.com/drak3/aoc19/errors"
	"github.com/drak3/aoc19/file"
)

func main() {
	in := file.ReadLines("day3/in.txt")
	w1 := strings.Trim(in[0], "\n")
	w2 := strings.Trim(in[1], "\n")
	fmt.Println(solve2(w1, w2))
}

func solve2(w1, w2 string) int {
	wire1 := strings.Split(w1, ",")
	wire2 := strings.Split(w2, ",")
	seen1 := walk(wire1)
	seen2 := walk(wire2)
	var crosses []point
	for k := range seen1 {
		if _, ok := seen2[k]; ok {
			crosses = append(crosses, k)
		}
	}
	fmt.Println(crosses)
	mindist := 100000
	minp := point{0, 0}
	for _, c := range crosses {
		if seen1[c]+seen2[c] < mindist {
			minp = c
			mindist = seen1[c] + seen2[c]
		}
	}
	return seen1[minp] + seen2[minp]
}

func walk(directions []string) map[point]int {
	var dx = map[rune]int{
		'U': 0,
		'D': 0,
		'L': -1,
		'R': 1,
	}

	var dy = map[rune]int{
		'U': 1,
		'D': -1,
		'L': 0,
		'R': 0,
	}
	steps := 0
	p := point{0, 0}
	m := make(map[point]int)
	for _, w := range directions {
		d1, x1 := splitDirection(w)
		for i := 0; i < x1; i++ {
			steps++
			p.x += dx[d1]
			p.y += dy[d1]
			if _, ok := m[p]; !ok {
				m[p] = steps
			}
		}
	}
	return m
}

func manhatten(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

type point struct {
	x, y int
}

func splitDirection(s string) (rune, int) {
	x, err := strconv.Atoi(s[1:])
	e.Ck(err)
	return rune(s[0]), x
}
