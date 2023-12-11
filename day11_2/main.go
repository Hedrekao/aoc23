package main

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x float64
	y float64
}

type Pair struct {
	a Point
	b Point
}

type Set map[int]bool

func (s Set) Add(i int) {
	s[i] = true
}

func (s Set) Contains(i int) bool {
	_, ok := s[i]
	return ok
}

func main() {
	content, err := utils.LoadTextFile("data/input11.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")
	galaxies := make([]Point, 0)
	emptyRows := make(map[int]bool)
	emptyCols := make(map[int]bool)
	sum := 0
	for rowIdx, line := range lines {
		emptyRows[rowIdx] = true
		for colIdx, char := range line {
			if rowIdx == 0 {
				emptyCols[colIdx] = true
			}
			if char == '#' {
				galaxies = append(galaxies, Point{x: float64(colIdx), y: float64(rowIdx)})
				emptyRows[rowIdx] = false
				emptyCols[colIdx] = false
			}
		}
	}
	emptyColsArr := make([]int, 0)
	for colIdx, isEmpty := range emptyCols {
		if isEmpty {
			emptyColsArr = append(emptyColsArr, colIdx)
		}
	}
	emptyRowsArr := make([]int, 0)
	for rowIdx, isEmpty := range emptyRows {
		if isEmpty {
			emptyRowsArr = append(emptyRowsArr, rowIdx)
		}
	}

	pairs := make([]Pair, 0)
	for i := 0; i < len(galaxies); i++ {

		firstGalaxy := galaxies[i]

		for j := i + 1; j < len(galaxies); j++ {
			secondGalaxy := galaxies[j]
			pairs = append(pairs, Pair{a: firstGalaxy, b: secondGalaxy})
		}

	}

	for _, pair := range pairs {
		manhattanDist := math.Abs(pair.a.x-pair.b.x) + math.Abs(pair.a.y-pair.b.y)
		emptyRowsInBetween := 0
		for _, emptyRow := range emptyRowsArr {
			if (float64(emptyRow) > pair.a.y && float64(emptyRow) < pair.b.y) || (float64(emptyRow) < pair.a.y && float64(emptyRow) > pair.b.y) {
				emptyRowsInBetween += 999999
			}
		}
		emptyColsInBetween := 0
		for _, emptyCol := range emptyColsArr {
			if (float64(emptyCol) > pair.a.x && float64(emptyCol) < pair.b.x) || (float64(emptyCol) < pair.a.x && float64(emptyCol) > pair.b.x) {
				emptyColsInBetween += 999999
			}

		}
		sum += int(manhattanDist) + emptyRowsInBetween + emptyColsInBetween
	}

	fmt.Println(sum)
}
