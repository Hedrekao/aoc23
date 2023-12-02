package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content, err := utils.LoadTextFile("data/input2.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		_, combinations := parts[0], strings.Split(parts[1], ";")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		if err != nil {
			panic(err)
		}
		for _, combination := range combinations {
			combination = strings.Trim(combination, " ")
			cubes := strings.Split(combination, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				colorValue := strings.Split(cube, " ")
				value, _ := strconv.Atoi(colorValue[0])
				color := strings.TrimSpace(colorValue[1])
				switch color {
				case "red":
					if value > maxRed {
						maxRed = value
					}
				case "green":
					if value > maxGreen {
						maxGreen = value
					}
				case "blue":
					if value > maxBlue {
						maxBlue = value
					}
				}
			}
		}

		sum += maxBlue * maxGreen * maxRed
	}

	fmt.Println(sum)
}
