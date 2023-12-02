package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var maxMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	content, err := utils.LoadTextFile("data/input2.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	sumOfIds := 0
	for _, line := range lines {
		isValid := true
		parts := strings.Split(line, ":")
		idPart, combinations := parts[0], strings.Split(parts[1], ";")
		gameId, err := strconv.Atoi(strings.TrimSpace(strings.Split(idPart, " ")[1]))
		if err != nil {
			panic(err)
		}
		for _, combination := range combinations {
			if !isValid {
				break
			}
			combination = strings.Trim(combination, " ")
			cubes := strings.Split(combination, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				colorValue := strings.Split(cube, " ")
				value, _ := strconv.Atoi(colorValue[0])
				color := strings.TrimSpace(colorValue[1])
				if value > maxMap[color] {
					isValid = false
					break
				}

			}
		}

		if isValid {
			sumOfIds += gameId
		}
	}

	fmt.Println(sumOfIds)
}
