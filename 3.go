package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getMul(line string) int {
	total := 0
	expression := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := expression.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		firstNumber, _ := strconv.Atoi(match[1])
		secondNumber, _ := strconv.Atoi(match[2])
		total += firstNumber * secondNumber
	}

	return total
}

func part1(scanner *bufio.Scanner) {
	expression := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	var total int
	for scanner.Scan() {
		matches := expression.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			firstNumber, _ := strconv.Atoi(match[1])
			secondNumber, _ := strconv.Atoi(match[2])
			total += firstNumber * secondNumber
		}

	}
	fmt.Printf("total: %v\n", total)
}

func part2(scanner *bufio.Scanner) {
	total := 0
	do := "do()"
	dont := "don't()"
	var fullLine string
	for scanner.Scan() {
		line := scanner.Text()
		fullLine += line
	}

	untilDo := true
	ok := true

	for ok {
		if untilDo {
			before, after, found := strings.Cut(fullLine, dont)
			if !found {
				ok = false
			}
			total += getMul(before)
			fullLine = after
			untilDo = false
		} else {
			_, after, found := strings.Cut(fullLine, do)
			if !found {
				ok = false
			}
			fullLine = after
			untilDo = true
		}

	}
	fmt.Printf("total: %v\n", total)
}

func main() {
	file, err := os.Open("3.txt")
	if err != nil {
		fmt.Printf("error open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	//part1(scanner)
	part2(scanner)
}
