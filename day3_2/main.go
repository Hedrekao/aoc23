package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type GearData struct {
	numbersAround []int
}

type GearDataMap map[int]GearData

func main() {
	content, err := utils.LoadTextFile("data/input3.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	currentNumber := ""
	gearDataMap := make(GearDataMap)
	currentGearId := -1
	isNumberValid := false
	gearRow := -1
	gearCol := -1
	sum := 0
	for rowIdx, line := range lines {
		for columnIdx, charNum := range line {
			char := string(charNum)
			isDigit := unicode.IsDigit(charNum)
			if !isDigit {
				if currentNumber != "" && isNumberValid {
					number, _ := strconv.Atoi(currentNumber)
					gearDataMap[currentGearId] = GearData{append(gearDataMap[currentGearId].numbersAround, number)}
					if len(gearDataMap[currentGearId].numbersAround) == 2 {
						sum += gearDataMap[currentGearId].numbersAround[0] * gearDataMap[currentGearId].numbersAround[1]
						delete(gearDataMap, currentGearId)
					}
					isNumberValid = false

				}

				currentNumber = ""
			} else {
				if !isNumberValid {
					isNumberValid, gearRow, gearCol = checkIfValidNumber(rowIdx, columnIdx, lines)
					if isNumberValid {
						cellId := gearRow*len(lines[0]) + gearCol
						if _, exists := gearDataMap[cellId]; !exists {
							gearDataMap[cellId] = GearData{[]int{}}
						}
						currentGearId = cellId
					}
				}

				currentNumber += char

			}
		}

	}

	fmt.Println(sum)
}

func checkIfValidNumber(rowIdx, columnIdx int, lines []string) (bool, int, int) {
	if rowIdx > 0 {
		char := lines[rowIdx-1][columnIdx]
		if char == '*' {
			return true, rowIdx - 1, columnIdx
		}
		if columnIdx > 0 {
			char := lines[rowIdx-1][columnIdx-1]
			if char == '*' {
				return true, rowIdx - 1, columnIdx - 1
			}
		}
		if columnIdx < len(lines[0])-2 {
			char := lines[rowIdx-1][columnIdx+1]
			if char == '*' {
				return true, rowIdx - 1, columnIdx + 1
			}
		}
	}
	if rowIdx < len(lines)-1 {
		char := lines[rowIdx+1][columnIdx]
		if char == '*' {
			return true, rowIdx + 1, columnIdx
		}
		if columnIdx > 0 {
			char := lines[rowIdx+1][columnIdx-1]
			if char == '*' {
				return true, rowIdx + 1, columnIdx - 1
			}
		}
		if columnIdx < len(lines[0])-2 {
			char := lines[rowIdx+1][columnIdx+1]
			if char == '*' {
				return true, rowIdx + 1, columnIdx + 1
			}
		}
	}

	if columnIdx > 0 {
		char := lines[rowIdx][columnIdx-1]
		if char == '*' {
			return true, rowIdx, columnIdx - 1
		}
	}
	if columnIdx < len(lines[0])-2 {
		char := lines[rowIdx][columnIdx+1]
		if char == '*' {
			return true, rowIdx, columnIdx + 1
		}
	}

	return false, -1, -1
}
