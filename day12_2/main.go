package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	pattern     string
	numbersCode string
}

var cache = make(map[Input]int)

func setCache(pattern string, numbers []int, result int) int {

	numbersCode := ""
	for idx, n := range numbers {
		numbersCode += strconv.Itoa(n) + strconv.Itoa(idx)
	}
	cache[Input{pattern, numbersCode}] = result
	return result
}

func main() {
	content, err := utils.LoadTextFile("data/input12.txt")

	if err != nil {
		panic(err)
	}
	sum := 0
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		strNumbers := strings.Split(lineParts[1], ",")
		damagedArr := make([]int, len(strNumbers))
		for i, strNumber := range strNumbers {
			number, _ := strconv.Atoi(strNumber)
			damagedArr[i] = number
		}
		pattern := lineParts[0]
		newDamagedArr := make([]int, len(damagedArr))
		newPattern := pattern
		copy(newDamagedArr, damagedArr)
		for i := 0; i < 4; i++ {
			newPattern = newPattern + "?" + pattern
			newDamagedArr = append(newDamagedArr, damagedArr...)
		}
		sum += count(newPattern, newDamagedArr)

	}

	fmt.Println(sum)
}

func count(pattern string, numbers []int) int {
	if len(pattern) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(pattern) == 0 {
		return 0
	}

	numbersCode := ""
	for idx, n := range numbers {
		numbersCode += strconv.Itoa(n) + strconv.Itoa(idx)
	}
	if val, ok := cache[Input{pattern, numbersCode}]; ok {
		return val
	}

	if pattern[0] == '.' {
		result := count(pattern[1:], numbers)
		return setCache(pattern, numbers, result)
	}

	// cut branches
	var sum int
	for _, n := range numbers {
		sum += int(n)
	}
	if len(pattern) < sum {
		result := 0
		return setCache(pattern, numbers, result)
	}

	if pattern[0] == '?' {
		result := count(pattern[1:], numbers) + count("#"+pattern[1:], numbers)
		return setCache(pattern, numbers, result)
	}

	if pattern[0] == '#' {
		if len(numbers) == 0 {
			result := 0
			return setCache(pattern, numbers, result)
		}

		n := numbers[0]
		indexDot := strings.Index(pattern, ".")
		if indexDot == -1 {
			indexDot = len(pattern)
		}
		if indexDot < int(n) {
			// not enough # or ?
			result := 0
			return setCache(pattern, numbers, result)
		}

		// eat n # or ?
		remaining := pattern[n:]
		if len(remaining) == 0 {
			result := count(remaining, numbers[1:])
			return setCache(pattern, numbers, result)
		}

		if remaining[0] == '#' {
			// fail
			result := 0
			return setCache(pattern, numbers, result)
		}
		// remaining[0] == '.' || remaining[0] == '?'
		// eat first ? since it should be a .
		result := count(remaining[1:], numbers[1:])
		return setCache(pattern, numbers, result)
	}
	panic("unreachable")
}
