package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input.txt")
	vals := string(file)

	var leftArr, rightArr []int
	splt := strings.Split(vals, "\n")

	for _, v := range splt {
		j := strings.Split(v, "   ")

		lVal, _ := strconv.Atoi(j[0])
		rVal, _ := strconv.Atoi(j[1])

		leftArr = append(leftArr, lVal)
		rightArr = append(rightArr, rVal)
	}

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	x := part1(leftArr, rightArr)
	y := part2(leftArr, rightArr)

	fmt.Println("Part 1:", x)
	fmt.Println("Part 2:", y)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func part1(lArr []int, rArr []int) int {
	var output int
	for n := range lArr {
		lNum := lArr[n]
		rNum := rArr[n]

		h := diff(lNum, rNum)
		output += h
	}

	return output
}

func part2(lArr []int, rArr []int) int {
	var out int
	for _, currLeftArrNum := range lArr {
		var numOccurancesRighArr int
		for _, y := range rArr {
			if currLeftArrNum == y {
				numOccurancesRighArr += 1
			}
		}
		if numOccurancesRighArr > 0 {
			out += (currLeftArrNum * numOccurancesRighArr)
		}
	}

	return out
}
