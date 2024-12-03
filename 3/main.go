package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input.txt")
	vals := strings.Split(string(file), "\n")

	p1 := part1(vals)
	p2 := part2(vals)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func part1(vals []string) int {
	var total int

	for _, i := range vals {
		match, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
		out := match.FindAllStringSubmatch(i, -1)
		for _, j := range out {
			x1, _ := strconv.Atoi(j[1])
			x2, _ := strconv.Atoi(j[2])
			total += (x1 * x2)
		}
	}
	return total
}

func part2(vals []string) int {
	var total int
	var locked bool = false

	for _, i := range vals {
		mul, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		out := mul.FindAllStringSubmatch(i, -1)

		for _, j := range out {
			curr := j[0]

			switch curr {
			case "do()":
				locked = false
			case "don't()":
				locked = true
			}

			if !locked && strings.HasPrefix(curr, "mul(") {
				parts := strings.TrimPrefix(strings.TrimSuffix(curr, ")"), "mul(")
				numbers := strings.Split(parts, ",")
				x1, _ := strconv.Atoi(numbers[0])
				x2, _ := strconv.Atoi(numbers[1])
				total += (x1 * x2)
			}
		}
	}
	return total
}
