package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content, err := utils.LoadTextFile("data/input9.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\r\n")
	sum := 0
	for _, line := range lines {
		sum += handleLine(line)
	}
	fmt.Println(sum)
}

func handleLine(line string) int {
	numbers := strings.Split(line, " ")
	allDifferences := make([][]int, 0)
	allZero := true
	baseNumbers := make([]int, 0)
	baseDifferences := make([]int, 0)
	for i := 0; i < len(numbers)-1; i++ {
		firstNumber, _ := strconv.Atoi(numbers[i])
		secondNumber, _ := strconv.Atoi(numbers[i+1])
		baseNumbers = append(baseNumbers, firstNumber)
		if i == len(numbers)-2 {
			baseNumbers = append(baseNumbers, secondNumber)
		}
		difference := secondNumber - firstNumber
		if difference != 0 {
			allZero = false
		}
		baseDifferences = append(baseDifferences, difference)
	}
	allDifferences = append(allDifferences, baseNumbers)
	allDifferences = append(allDifferences, baseDifferences)

	differencesIdx := 1
	for !allZero {
		differences := make([]int, 0)
		previousDifferences := allDifferences[differencesIdx]
		allZero = true
		for i := 0; i < len(previousDifferences)-1; i++ {
			firstNumber := previousDifferences[i]
			secondNumber := previousDifferences[i+1]
			difference := secondNumber - firstNumber
			if difference != 0 {
				allZero = false
			}
			differences = append(differences, difference)
		}
		allDifferences = append(allDifferences, differences)
		differencesIdx++
	}

	prediction := getPrediction(allDifferences)
	return prediction
}

func getPrediction(allDifferences [][]int) int {
	prediction := 0
	for i := len(allDifferences) - 2; i >= 0; i-- {
		differences := allDifferences[i]
		prediction = differences[0] - prediction
	}
	return prediction
}
