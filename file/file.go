//Package file implements file utils useful for advent of code solutions
//generally functions panic instead of returning an error
package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Read reads the content of a file into a string
func Read(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

//ReadInts reads the content of a file as newline seperated ints
func ReadInts(file string) []int {
	lines := ReadLines(file)
	var r []int
	for _, l := range lines {
		i, _ := strconv.Atoi(l)
		r = append(r, i)
	}
	return r
}

//ReadLines reads all lines of a file into a slice of strings
func ReadLines(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	var r []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r = append(r, scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return r
}
