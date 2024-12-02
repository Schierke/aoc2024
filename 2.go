package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Order int

const (
	decrease Order = 0
	increase Order = 1
)

func isSafe(line []string, order Order) bool {
	before, _ := strconv.Atoi(line[0])

	for i := range line {
		if i == 0 {
			continue
		}
		number, _ := strconv.Atoi(line[i])
		if order == increase {
			if number <= before || number-before > 3 {
				return false
			}
		} else {
			if number >= before || before-number > 3 {
				return false
			}
		}
		before = number
	}

	return true
}

func part1(scanner *bufio.Scanner) {
	var total int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if isSafe(line, increase) || isSafe(line, decrease) {
			fmt.Println(line)
			total++
		}
	}
	fmt.Printf("total: %v\n", total)
}

func part2(scanner *bufio.Scanner) {
	total := 0

	index := func(line []string, order Order) int {
		before, _ := strconv.Atoi(line[0])
		badLevel := 0

		for i := range line {
			if i == 0 {
				continue
			}
			number, _ := strconv.Atoi(line[i])
			if order == increase {
				if number <= before || number-before > 3 {
					return i
				}
			} else {
				if number >= before || before-number > 3 {
					return i
				}
			}
			before = number
		}

		return badLevel
	}

	newSlice := func(line []string, idx int) []string {
		var s []string
		for i, v := range line {
			if i != idx {
				s = append(s, v)
			}
		}

		return s
	}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if isSafe(line, increase) || isSafe(line, decrease) {
			total++
		} else {
			idx := index(line, increase)
			if idx != 0 {
				s1 := newSlice(line, idx)
				s2 := newSlice(line, idx-1)

				if isSafe(s1, increase) || isSafe(s2, increase) {
					total++
				}
			}

			idx = index(line, decrease)
			if idx != 0 {
				s1 := newSlice(line, idx)
				s2 := newSlice(line, idx-1)

				if isSafe(s1, decrease) || isSafe(s2, decrease) {
					total++
				}
			}
		}
	}

	fmt.Printf("total: %v\n", total)
}

func main() {
	file, err := os.Open("2.txt")
	if err != nil {
		fmt.Printf("error open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	//part1(scanner)
	part2(scanner)
}
