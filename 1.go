package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	var list1, list2 []int
	var sum int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		firstNumber, _ := strconv.Atoi(line[0])
		secondNumber, _ := strconv.Atoi(line[1])

		list1 = append(list1, firstNumber)
		list2 = append(list2, secondNumber)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := range list1 {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Printf("sum: %v\n", sum)
}

func part2(scanner *bufio.Scanner) {
	occurs := make(map[int]int)
	var list []int
	var sum int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		firstNumber, _ := strconv.Atoi(line[0])
		secondNumber, _ := strconv.Atoi(line[1])

		list = append(list, firstNumber)
		occurs[secondNumber]++
	}

	for _, value := range list {

		sum += value * occurs[value]
	}
	fmt.Printf("sum: %v\n", sum)
}

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Printf("error open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	//part1(scanner)
	part2(scanner)
}
