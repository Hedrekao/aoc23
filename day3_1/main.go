package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := utils.LoadTextFile("data/input3.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	currentNumber := ""
	isNumberValid := false
	sum := 0
	for rowIdx, line := range lines {
		for columnIdx, charNum := range line {
			char := string(charNum)
			isDigit := unicode.IsDigit(charNum)
			if !isDigit {
				if currentNumber != "" {
					number, _ := strconv.Atoi(currentNumber)
					if isNumberValid {
						sum += number
						isNumberValid = false
					}
					currentNumber = ""
				}
			} else {
				currentNumber += char
				if !isNumberValid {
					isNumberValid = checkIfValidNumber(rowIdx, columnIdx, lines)
				}
			}
		}

	}

	fmt.Println(sum)
}

func checkIfValidNumber(rowIdx, columnIdx int, lines []string) bool {
	if rowIdx > 0 {
		char := lines[rowIdx-1][columnIdx]
		if char != '.' && !unicode.IsDigit(rune(char)) {
			return true
		}
		if columnIdx > 0 {
			char := lines[rowIdx-1][columnIdx-1]
			if char != '.' && !unicode.IsDigit(rune(char)) {
				return true
			}
		}
		if columnIdx < len(lines[0])-2 {
			char := lines[rowIdx-1][columnIdx+1]
			if char != '.' && !unicode.IsDigit(rune(char)) {
				return true
			}
		}
	}
	if rowIdx < len(lines)-1 {
		char := lines[rowIdx+1][columnIdx]
		if char != '.' && !unicode.IsDigit(rune(char)) {
			return true
		}
		if columnIdx > 0 {
			char := lines[rowIdx+1][columnIdx-1]
			if char != '.' && !unicode.IsDigit(rune(char)) {
				return true
			}
		}
		if columnIdx < len(lines[0])-2 {
			char := lines[rowIdx+1][columnIdx+1]
			if char != '.' && !unicode.IsDigit(rune(char)) {
				return true
			}
		}
	}

	if columnIdx > 0 {
		char := lines[rowIdx][columnIdx-1]
		if char != '.' && !unicode.IsDigit(rune(char)) {
			return true
		}
	}
	if columnIdx < len(lines[0])-2 {
		char := lines[rowIdx][columnIdx+1]
		if char != '.' && !unicode.IsDigit(rune(char)) {
			return true
		}
	}

	return false
}
