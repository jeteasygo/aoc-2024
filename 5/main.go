package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file1, _ := os.ReadFile("./input1.txt")
	file2, _ := os.ReadFile("./input2.txt")
	po := strings.Split(string(file1), "\n")
	updates := strings.Split(string(file2), "\n")
	var updatesInts [][]int
	for _, update := range updates {
		var updateInts []int
		for _, u := range strings.Split(update, ",") {
			num, _ := strconv.Atoi(u)
			updateInts = append(updateInts, num)
		}
		updatesInts = append(updatesInts, updateInts)
	}

	p1 := part1(po, updatesInts)
	p2 := part2(po, updatesInts)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func part1(pageOrders []string, updates [][]int) int {
	var total int
	pageOrder := make(map[int][]int)

	for _, po := range pageOrders {
		curr := strings.Split(po, "|")
		if len(curr) == 2 {
			y, _ := strconv.Atoi(curr[1])
			x, _ := strconv.Atoi(curr[0])
			pageOrder[x] = append(pageOrder[x], y)
		}
	}

	for _, update := range updates {
		valid := true
		for i, currentUpdateNum := range update {
			if orders, exists := pageOrder[currentUpdateNum]; exists {
				for _, order := range orders {
					for j := 0; j < i; j++ {
						if update[j] == order {
							valid = false
							break
						}
					}
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			total += update[len(update)/2]
		}
	}

	return total
}

func part2(pageOrders []string, updates [][]int) int {
	var total int
	pageOrder := make(map[int][]int)
	var failedUpdates [][]int

	for _, po := range pageOrders {
		curr := strings.Split(po, "|")
		if len(curr) == 2 {
			y, _ := strconv.Atoi(curr[1])
			x, _ := strconv.Atoi(curr[0])
			pageOrder[x] = append(pageOrder[x], y)
		}
	}

	for _, update := range updates {
		valid := true

		for i, currentUpdateNum := range update {
			if orders, exists := pageOrder[currentUpdateNum]; exists {
				for _, order := range orders {
					for j := 0; j < i; j++ {
						if update[j] == order {
							valid = false
							break
						}
					}
				}
			}
			if !valid {
				break
			}
		}

		if !valid {
			failedUpdates = append(failedUpdates, update)
		}
	}

	for _, update := range failedUpdates {
		tmpArr := make([]int, len(update))
		copy(tmpArr, update)
		changed := true

		for changed {
			changed = false
			for i := 0; i < len(tmpArr); i++ {
				currentUpdateNum := tmpArr[i]
				if orders, exists := pageOrder[currentUpdateNum]; exists {
					for _, order := range orders {
						for j := i + 1; j < len(tmpArr); j++ {
							if tmpArr[j] == order {
								tmpArr[i], tmpArr[j] = tmpArr[j], tmpArr[i]
								changed = true
							}
						}
					}
				}
			}
		}
		total += tmpArr[len(tmpArr)/2]
	}

	return total
}
