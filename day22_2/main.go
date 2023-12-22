package main

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Brick struct {
	x1 int
	y1 int
	z1 int
	x2 int
	y2 int
	z2 int
}

func main() {
	content, err := utils.LoadTextFile("data/input22.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")

	bricks := make([]Brick, 0)
	for _, line := range lines {
		parts := strings.Split(line, "~")

		startCoords := strings.Split(parts[0], ",")
		x1, _ := strconv.Atoi(startCoords[0])
		y1, _ := strconv.Atoi(startCoords[1])
		z1, _ := strconv.Atoi(startCoords[2])

		endCoords := strings.Split(parts[1], ",")
		x2, _ := strconv.Atoi(endCoords[0])
		y2, _ := strconv.Atoi(endCoords[1])
		z2, _ := strconv.Atoi(endCoords[2])

		brick := Brick{x1, y1, z1, x2, y2, z2}

		bricks = append(bricks, brick)

	}

	slices.SortFunc(bricks, func(a, b Brick) int {
		if a.z1 < b.z1 {
			return -1
		} else if a.z1 > b.z1 {
			return 1
		} else {
			return 0
		}

	})

	fall(bricks)

	result := 0

	for idx := range bricks {

		tmpBricks := make([]Brick, len(bricks))
		copy(tmpBricks, bricks)
		tmpBricks = removeElement(tmpBricks, idx)

		count := fall(tmpBricks)

		result += count
	}

	fmt.Println(result)
}

func removeElement(slice []Brick, index int) []Brick {
	// Check if the index is valid
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}

	// Remove the element at the specified index
	return append(slice[:index], slice[index+1:]...)
}

func fall(bricks []Brick) int {

	count := 0
	for i := 0; i < len(bricks); i++ {
		brick := &bricks[i]
		gotCollision := false
		previousZ := brick.z1

		for brick.z1 > 1 {
			for j := i - 1; j >= 0; j-- {
				brickBelow := &bricks[j]

				if (brick.z2-1) >= brickBelow.z1 &&
					(brick.z1-1) <= brickBelow.z2 &&
					brick.x2 >= brickBelow.x1 &&
					brick.x1 <= brickBelow.x2 &&
					brick.y2 >= brickBelow.y1 &&
					brick.y1 <= brickBelow.y2 {
					gotCollision = true
				}

			}

			if gotCollision {
				break
			}

			brick.z1--
			brick.z2--

		}

		if previousZ != brick.z1 {
			count++
		}

	}
	return count
}
