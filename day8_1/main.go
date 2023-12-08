package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Direction struct {
	left  string
	right string
}

func main() {
	content, err := utils.LoadTextFile("data/input8.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\r\n")
	instructions := make([]string, 0)
	directions := make(map[string]Direction)

	steps := 0

	for idx, line := range lines {
		if idx == 0 {
			instructions = parseInstruction(line)
			continue
		}
		if line == "" {
			continue
		}
		key, direction := parseDirection(line)
		directions[key] = direction

	}

	currentPlace := "AAA"
	instructionsIdx := 0
	for currentPlace != "ZZZ" {
		steps++
		if instructionsIdx >= len(instructions) {
			instructionsIdx = 0
		}
		instruction := instructions[instructionsIdx]
		direction := directions[currentPlace]
		if instruction == "L" {
			currentPlace = direction.left
		} else {
			currentPlace = direction.right
		}
		currentPlace = strings.TrimSpace(currentPlace)
		instructionsIdx++
	}

	fmt.Println(directions)
	fmt.Println(instructions)

	fmt.Println(steps)
}

func parseInstruction(instruction string) []string {
	commands := strings.Split(instruction, "")

	result := make([]string, 0)
	result = append(result, commands...)
	return result

}

func parseDirection(direction string) (string, Direction) {
	splitted := strings.Split(direction, " = ")
	key := splitted[0]
	value := splitted[1]
	directions := strings.Split(value, ",")
	left := directions[0][1:]
	right := directions[1][1 : len(directions[1])-1]
	directionObj := Direction{left, right}
	return key, directionObj
}
