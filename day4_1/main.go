package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Set map[int]bool

func (s Set) Add(item int) {
	s[item] = true
}

func (s Set) Remove(item int) {
	delete(s, item)
}

func (s Set) Contains(item int) bool {
	return s[item]
}
func (s Set) Keys() []int {
	keys := make([]int, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	content, err := utils.LoadTextFile("data/input4.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	sum := 0
	for _, line := range lines {
		winningNumbersStr := strings.Split(strings.Split(line, "|")[0], ":")[1]
		line_score := 0.5
		winningNumbers := strings.Split(strings.TrimSpace(winningNumbersStr), " ")
		resultSet := Set{}
		for _, num := range winningNumbers {
			result, _ := strconv.Atoi(strings.TrimSpace(num))
			if result == 0 {
				continue
			}
			resultSet.Add(result)
		}
		guessedNumbersStr := strings.Split(line, "|")[1]
		guessedNumbers := strings.Split(guessedNumbersStr, " ")
		for _, guessedNumber := range guessedNumbers {
			fmt.Println(guessedNumber)
			if guessedNumber == "" {
				continue
			}
			num, _ := strconv.Atoi(strings.TrimSpace(guessedNumber))
			if resultSet.Contains(num) {
				fmt.Printf("winning number: %v \n", resultSet.Keys())
				line_score *= 2
			}
		}
		if line_score >= 1 {
			sum += int(line_score)
		}

	}

	fmt.Println(sum)
}
