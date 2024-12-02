package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncrementing(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return false
		}
	}
	return true
}

func isDecrementing(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			return false
		}
	}
	return true
}

func isIncrementingOrDecrementing(input []int) bool {
	return isIncrementing(input) || isDecrementing(input)
}

func main() {
	file, _ := os.ReadFile("./input.txt")
	vals := strings.Split(string(file), "\n")

	p1 := part1(vals)
	p2 := part2(vals)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func part1(vals []string) int {
	safeReports := 0
mainLoop:
	for _, row := range vals {
		rowVals := strings.Split(row, " ")

		// Step 1 - check incrementing or decrementing
		// convert row to int array
		var tmpRow []int
		for _, rowVal := range rowVals {
			x, _ := strconv.Atoi(rowVal)
			tmpRow = append(tmpRow, x)
		}
		// not safe if not incrementing or decrementing
		if !isIncrementingOrDecrementing(tmpRow) {
			continue
		}

		// Step 2 - check if left and right of number is within +-3
		for i := 0; i < len(rowVals)-1; i++ {
			curr, _ := strconv.Atoi(rowVals[i])
			nextNum, _ := strconv.Atoi(rowVals[i+1])
			if curr == nextNum {
				continue mainLoop
			}

			if curr-nextNum > 3 || nextNum-curr > 3 {
				continue mainLoop
			}
		}
		safeReports++
	}

	return safeReports
}

func isSafeReport(input []int) bool {
	if !isIncrementingOrDecrementing(input) {
		return false
	}

	// Step 2 - check if left and right of number is within +-3
	for i := 0; i < len(input)-1; i++ {
		curr := input[i]
		nextNum := input[i+1]
		if curr == nextNum {
			return false
		}

		if curr-nextNum > 3 || nextNum-curr > 3 {
			return false
		}
	}

	return true
}

func part2(vals []string) int {
	safeReports := 0
	var unsafeArr [][]int

mainLoop:
	for _, row := range vals {
		rowVals := strings.Split(row, " ")

		// Step 1 - check incrementing or decrementing
		// convert row to int array
		var tmpRow []int
		for _, rowVal := range rowVals {
			x, _ := strconv.Atoi(rowVal)
			tmpRow = append(tmpRow, x)
		}
		// not safe if not incrementing or decrementing
		if !isIncrementingOrDecrementing(tmpRow) {
			unsafeArr = append(unsafeArr, tmpRow)
			continue
		}

		// Step 2 - check if left and right of number is within +-3
		for i := 0; i < len(rowVals)-1; i++ {
			curr, _ := strconv.Atoi(rowVals[i])
			nextNum, _ := strconv.Atoi(rowVals[i+1])

			if curr == nextNum {
				unsafeArr = append(unsafeArr, tmpRow)
				continue mainLoop
			}

			if curr-nextNum > 3 || nextNum-curr > 3 {
				unsafeArr = append(unsafeArr, tmpRow)
				continue mainLoop
			}
		}
		safeReports++
	}

	fmt.Println("unsafeArr length:", len(unsafeArr))
outer:
	for _, unsafeRow := range unsafeArr {

		tmp := make([]int, len(unsafeRow))
		copy(tmp, unsafeRow)

		for n := range unsafeRow {
			modifiedTmp := make([]int, len(tmp))
			copy(modifiedTmp, tmp)

			tmpArr := append(modifiedTmp[:n], modifiedTmp[n+1:]...)

			if isSafeReport(tmpArr) {
				safeReports++
				continue outer
			}
		}
	}

	return safeReports
}
