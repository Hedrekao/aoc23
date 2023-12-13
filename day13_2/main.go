package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

func main() {
	content, err := utils.LoadTextFile("data/input13.txt")

	if err != nil {
		panic(err)
	}
	sum := 0
	lines := strings.Split(content, "\n\n")

	for _, grid := range lines {
		possibleRowSymetries := make(map[int]int)
		lines := strings.Split(grid, "\n")
		for i := 0; i < len(lines); i++ {
			for idx, line := range lines {
				if idx == 0 {
					possibleRowSymetries = possibleRowSymetriesIdx(strings.Split(lines[i], ""))
				}
				for key := range possibleRowSymetries {
					possibleRowSymetries[key] += checkRow(strings.Split(line, ""), key)

					if possibleRowSymetries[key] > 1 {
						delete(possibleRowSymetries, key)
					}

				}
				if len(possibleRowSymetries) == 0 {
					break
				}
			}
			if len(possibleRowSymetries) > 0 {
				foundSomething := false
				for key := range possibleRowSymetries {
					if possibleRowSymetries[key] == 1 {
						sum += key + 1
						foundSomething = true
					}
				}
				if foundSomething {
					break
				}
			}

		}
		for idx := 0; idx < len(lines[0]); idx++ {
			possibleColumnSymetries := possibleColumnSymetriesIdx(grid, idx)

			for i := 0; i < len(lines[0]); i++ {
				for key := range possibleColumnSymetries {
					possibleColumnSymetries[key] += checkColumns(strings.Split(grid, "\n"), i, key)
					if possibleColumnSymetries[key] > 1 {
						delete(possibleColumnSymetries, key)
					}
				}

				if len(possibleColumnSymetries) == 0 {
					break
				}

			}
			if len(possibleColumnSymetries) > 0 {
				foundSomething := false
				for key, val := range possibleColumnSymetries {
					if val == 1 {
						sum += (key + 1) * 100
						foundSomething = true
					}
				}
				if foundSomething {
					break
				}
			}
		}

	}
	fmt.Println(sum)
}

func checkColumns(lines []string, columnIdx int, symetryIdx int) int {
	remainingForward := len(lines) - symetryIdx - 2
	remainingBackward := symetryIdx
	min := min(remainingForward, remainingBackward)

	if symetryIdx == 0 {
		if lines[symetryIdx][columnIdx] == lines[symetryIdx+1][columnIdx] {
			return 0
		}
		return 1
	}

	differences := 0
	for i := 0; i <= min; i++ {
		if lines[symetryIdx-i][columnIdx] != lines[symetryIdx+1+i][columnIdx] {
			differences++
		}
	}
	return differences
}

func checkRow(row []string, symetryIdx int) int {
	remainingForward := len(row) - symetryIdx - 2
	remainingBackward := symetryIdx
	min := min(remainingForward, remainingBackward)

	if symetryIdx == 0 {
		if row[symetryIdx] == row[symetryIdx+1] {
			return 0
		}
		return 1
	}

	differences := 0
	for i := 0; i <= min; i++ {
		if row[symetryIdx-i] != row[symetryIdx+1+i] {
			differences++
		}
	}
	return differences
}

func possibleColumnSymetriesIdx(grid string, columnIdx int) map[int]int {
	var idx = make(map[int]int)
	lines := strings.Split(grid, "\n")
	for i := 0; i < len(lines)-1; i++ {
		j := 0
		hasSymetry := true
		for i-j >= 0 && i+j+1 <= len(lines)-1 {
			if lines[i-j][columnIdx] != lines[i+j+1][columnIdx] {
				hasSymetry = false
				break
			}
			j++
		}
		if hasSymetry {
			idx[i] = 0
		}
	}
	return idx
}

func possibleRowSymetriesIdx(line []string) map[int]int {
	var idx = make(map[int]int)
	length := len(line)
	for i := 0; i < length-1; i++ {
		j := 0
		hasSymetry := true
		for i-j >= 0 && i+j+1 <= length-1 {
			if line[i-j] != line[i+1+j] {
				hasSymetry = false
				break
			}
			j++
		}
		if hasSymetry {
			idx[i] = 0
		}
	}
	return idx
}
