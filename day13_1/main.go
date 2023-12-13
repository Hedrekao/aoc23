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
		possibleRowSymetries := make(map[int]bool)
		for idx, line := range strings.Split(grid, "\n") {
			if idx == 0 {
				possibleRowSymetries = possibleRowSymetriesIdx(strings.Split(line, ""))
				continue
			}
			for key, value := range possibleRowSymetries {
				if value {
					if !checkRow(strings.Split(line, ""), key) {
						delete(possibleRowSymetries, key)
					}
				}
			}
			if len(possibleRowSymetries) == 0 {
				break
			}
		}
		if len(possibleRowSymetries) > 0 {
			for key := range possibleRowSymetries {
				sum += key + 1
			}
		}
		possibleColumnSymetries := possibleColumnSymetriesIdx(grid)

		lines := strings.Split(grid, "\n")
		for i := 0; i < len(lines[0]); i++ {
			for key := range possibleColumnSymetries {
				if !checkColumns(strings.Split(grid, "\n"), i, key) {
					delete(possibleColumnSymetries, key)
				}
			}

			if len(possibleColumnSymetries) == 0 {
				break
			}

		}
		if len(possibleColumnSymetries) == 1 {
			for key := range possibleColumnSymetries {
				sum += (key + 1) * 100
			}
		}

	}
	fmt.Println(sum)
}

func checkColumns(lines []string, columnIdx int, symetryIdx int) bool {
	remainingForward := len(lines) - symetryIdx - 2
	remainingBackward := symetryIdx
	min := min(remainingForward, remainingBackward)

	if symetryIdx == 0 {
		return lines[symetryIdx][columnIdx] == lines[symetryIdx+1][columnIdx]
	}

	for i := 0; i <= min; i++ {
		if lines[symetryIdx-i][columnIdx] != lines[symetryIdx+1+i][columnIdx] {
			return false
		}
	}
	return true
}

func checkRow(row []string, symetryIdx int) bool {
	remainingForward := len(row) - symetryIdx - 2
	remainingBackward := symetryIdx
	min := min(remainingForward, remainingBackward)

	if symetryIdx == 0 {
		return row[symetryIdx] == row[symetryIdx+1]
	}

	for i := 0; i <= min; i++ {
		if row[symetryIdx-i] != row[symetryIdx+1+i] {
			return false
		}
	}
	return true
}

func possibleColumnSymetriesIdx(grid string) map[int]bool {
	var idx = make(map[int]bool)
	lines := strings.Split(grid, "\n")
	for i := 0; i < len(lines)-1; i++ {
		j := 0
		hasSymetry := true
		for i-j >= 0 && i+j+1 <= len(lines)-1 {
			if lines[i-j][0] != lines[i+j+1][0] {
				hasSymetry = false
				break
			}
			j++
		}
		if hasSymetry {
			idx[i] = true
		}
	}
	return idx
}

func possibleRowSymetriesIdx(line []string) map[int]bool {
	var idx = make(map[int]bool)
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
			idx[i] = true
		}
	}
	return idx
}
