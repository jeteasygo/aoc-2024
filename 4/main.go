package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input.txt")
	arr := make([][]string, len(strings.Split(string(file), "\n")))

	// make 2d array
	for i, line := range strings.Split(string(file), "\n") {
		arr[i] = strings.Split(line, "")
	}

	p1 := part1(arr)
	p2 := part2(arr)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func part1(vals [][]string) int {
	var total int
	for x := range vals {
		for y := range vals[x] {
			if vals[x][y] == "X" {
				// check right
				if y+3 < len(vals[0]) && vals[x][y+1] == "M" && vals[x][y+2] == "A" && vals[x][y+3] == "S" {
					total++
				}
				// check left
				if y-3 >= 0 && vals[x][y-1] == "M" && vals[x][y-2] == "A" && vals[x][y-3] == "S" {
					total++
				}
				// check down
				if x+3 < len(vals) && vals[x+1][y] == "M" && vals[x+2][y] == "A" && vals[x+3][y] == "S" {
					total++
				}
				// check up
				if x-3 >= 0 && vals[x-1][y] == "M" && vals[x-2][y] == "A" && vals[x-3][y] == "S" {
					total++
				}
				// check diagonal down right
				if x+3 < len(vals) && y+3 < len(vals[0]) && vals[x+1][y+1] == "M" && vals[x+2][y+2] == "A" && vals[x+3][y+3] == "S" {
					total++
				}
				// check diagonal down left
				if x+3 < len(vals) && y-3 >= 0 && vals[x+1][y-1] == "M" && vals[x+2][y-2] == "A" && vals[x+3][y-3] == "S" {
					total++
				}
				// check diagonal up right
				if x-3 >= 0 && y+3 < len(vals[0]) && vals[x-1][y+1] == "M" && vals[x-2][y+2] == "A" && vals[x-3][y+3] == "S" {
					total++
				}
				// check diagonal up left
				if x-3 >= 0 && y-3 >= 0 && vals[x-1][y-1] == "M" && vals[x-2][y-2] == "A" && vals[x-3][y-3] == "S" {
					total++
				}
			}
		}
	}

	return total
}

func part2(vals [][]string) int {
	var total int
	for x := range vals {
		for y := range vals[x] {
			if vals[x][y] == "A" {
				if x-1 >= 0 && y-1 >= 0 && x+1 < len(vals) && y+1 < len(vals[0]) &&
					x+1 < len(vals) && y-1 >= 0 && x-1 >= 0 && y+1 < len(vals[0]) {

					upperLeft := vals[x-1][y-1]
					bottomRight := vals[x+1][y+1]
					validDiagonal1 := (upperLeft == "M" && bottomRight == "S") || (upperLeft == "S" && bottomRight == "M")

					upperRight := vals[x-1][y+1]
					bottomLeft := vals[x+1][y-1]
					validDiagonal2 := (upperRight == "M" && bottomLeft == "S") || (upperRight == "S" && bottomLeft == "M")

					// only if both diagonals are valid
					if validDiagonal1 && validDiagonal2 {
						total++
					}
				}
			}
		}
	}

	return total
}
