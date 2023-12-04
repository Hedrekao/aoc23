package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

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
	n_lines := len(lines)
	sum := 0
	queue := Queue{}
	visited := Set{}
	for i := 0; i < n_lines; i++ {
		queue.Push(i)
	}

	for !queue.IsEmpty() {
		sum++
		lineIdx := queue.Pop()
		if visited.Contains(lineIdx) {
			continue
		}
		line := lines[lineIdx]
		n_of_winning_numbers := handleLine(line)
		if n_of_winning_numbers == 0 {
			visited.Add(lineIdx)
			continue
		}
		for i := 1; i <= n_of_winning_numbers; i++ {
			queue.Push(lineIdx + i)
		}
	}

	fmt.Println(sum)
}

func handleLine(line string) int {
	winningNumbersStr := strings.Split(strings.Split(line, "|")[0], ":")[1]
	winningNumbers := strings.Split(strings.TrimSpace(winningNumbersStr), " ")
	n_of_winning_numbers := 0
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
		if guessedNumber == "" {
			continue
		}
		num, _ := strconv.Atoi(strings.TrimSpace(guessedNumber))
		if resultSet.Contains(num) {
			n_of_winning_numbers++
		}
	}

	return n_of_winning_numbers
}
