package main

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	content, err := utils.LoadTextFile("data/input18.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")
	vertices := make([]Point, 0)
	vertices = append(vertices, Point{0, 0})
	currentX := 0
	currentY := 0
	currentDirection := ""
	distance := 0
	border := 0
	for idx, line := range lines {
		parts := strings.Split(line, " ")
		prevDirection := currentDirection
		distance, currentDirection = parseHexStr(parts[2])
		if idx != 0 {
			if prevDirection == "R" || prevDirection == "L" {
				if currentDirection == "U" || currentDirection == "D" {
					vertices = append(vertices, Point{currentX, currentY})
				}
			} else if prevDirection == "D" || prevDirection == "U" {
				if currentDirection == "R" || currentDirection == "L" {
					vertices = append(vertices, Point{currentX, currentY})
				}
			}
		}
		if currentDirection == "R" {
			currentX += distance
			border += distance
		}
		if currentDirection == "L" {
			currentX -= distance
			border += distance
		}
		if currentDirection == "U" {
			currentY -= distance
			border += distance
		}
		if currentDirection == "D" {
			currentY += distance
			border += distance
		}
	}
	shoelace := shoelaceArea(vertices)
	area := shoelace + border/2 + 1
	fmt.Println(area)

}

func shoelaceArea(vertices []Point) int {
	firstValue := 0
	secondValue := 0
	for i := 0; i < len(vertices)-1; i++ {
		firstValue += vertices[i].x * vertices[i+1].y
		secondValue += vertices[i+1].x * vertices[i].y
	}
	firstValue += vertices[len(vertices)-1].x * vertices[0].y
	secondValue += vertices[0].x * vertices[len(vertices)-1].y
	area := math.Abs(float64(firstValue-secondValue)) * 0.5
	return int(area)
}

func parseHexStr(hexStr string) (int, string) {
	hexStr = hexStr[2 : len(hexStr)-1]
	directionStr := hexStr[len(hexStr)-1:]
	distanceStr := hexStr[:len(hexStr)-1]
	distance, _ := strconv.ParseInt(distanceStr, 16, 32)
	directionNumber, _ := strconv.ParseInt(directionStr, 16, 32)

	direction := ""

	switch directionNumber {
	case 0:
		direction = "R"
	case 1:
		direction = "D"
	case 2:
		direction = "L"
	case 3:
		direction = "U"
	}

	return int(distance), direction
}
