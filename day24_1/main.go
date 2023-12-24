package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Hailstone struct {
	startX    float64
	startY    float64
	startZ    float64
	velocityX float64
	velocityY float64
	velocityZ float64
}

const (
	MIN = 200000000000000
	MAX = 400000000000000
)

func main() {
	content, err := utils.LoadTextFile("data/input24.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")
	hailstones := make([]Hailstone, 0)

	for _, line := range lines {
		parts := strings.Split(line, " @ ")

		position := strings.Split(parts[0], ", ")
		startX, _ := strconv.ParseFloat(strings.TrimSpace(position[0]), 64)
		startY, _ := strconv.ParseFloat(strings.TrimSpace(position[1]), 64)
		startZ, _ := strconv.ParseFloat(strings.TrimSpace(position[2]), 64)

		velocities := strings.Split(parts[1], ", ")
		velocityX, _ := strconv.ParseFloat(strings.TrimSpace(velocities[0]), 64)
		velocityY, _ := strconv.ParseFloat(strings.TrimSpace(velocities[1]), 64)
		velocityZ, _ := strconv.ParseFloat(strings.TrimSpace(velocities[2]), 64)

		hailstone := Hailstone{startX, startY, startZ, velocityX, velocityY, velocityZ}
		hailstones = append(hailstones, hailstone)

	}

	result := 0

	for i := 0; i < len(hailstones); i++ {
		hailstone := &hailstones[i]
		for j := i + 1; j < len(hailstones); j++ {
			hailstone2 := &hailstones[j]

			if checkCollision(hailstone, hailstone2) {
				result++
			}
		}
	}

	fmt.Println(result)
}

func checkCollision(hailstone1 *Hailstone, hailstone2 *Hailstone) bool {
	a1 := hailstone1.velocityY / hailstone1.velocityX
	b1 := hailstone1.startY - a1*hailstone1.startX

	a2 := hailstone2.velocityY / hailstone2.velocityX
	b2 := hailstone2.startY - a2*hailstone2.startX

	if a1 == a2 && b1 == b2 {
		return true
	}

	if a1 == a2 {
		return false
	}

	cx := (b2 - b1) / (a1 - a2)
	cy := a1*cx + b1

	if cx >= MIN && cx <= MAX && cy >= MIN && cy <= MAX {

		if hailstone1.velocityX > 0 && cx <= hailstone1.startX {
			return false
		}

		if hailstone1.velocityX < 0 && cx >= hailstone1.startX {
			return false
		}

		if hailstone1.velocityY > 0 && cy <= hailstone1.startY {
			return false
		}

		if hailstone1.velocityY < 0 && cy >= hailstone1.startY {
			return false
		}

		if hailstone2.velocityX > 0 && cx <= hailstone2.startX {
			return false
		}

		if hailstone2.velocityX < 0 && cx >= hailstone2.startX {

			return false
		}

		if hailstone2.velocityY > 0 && cy <= hailstone2.startY {
			return false
		}

		if hailstone2.velocityY < 0 && cy >= hailstone2.startY {
			return false
		}

		return true

	}

	return false
}
