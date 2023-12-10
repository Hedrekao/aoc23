package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Point struct {
	x        int
	y        int
	distance int
	char     rune
}

type PointQueue []Point

type Set map[int]bool

func (s Set) Add(i int) {
	s[i] = true
}

func (s Set) Contains(i int) bool {
	_, ok := s[i]
	return ok
}

func (pq *PointQueue) Push(p Point) {
	*pq = append(*pq, p)
}

func (pq *PointQueue) Pop() Point {
	point := (*pq)[0]
	*pq = (*pq)[1:]
	return point
}

func (pq *PointQueue) Len() int {
	return len(*pq)
}

func (pq *PointQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func main() {
	content, err := utils.LoadTextFile("data/input10.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")
	maze := make([][]rune, 0)
	startPoint := Point{x: 0, y: 0, distance: 0}
	for rowIdx, line := range lines {
		row := make([]rune, 0)
		for colIdx, char := range line {
			if char == 'S' {
				startPoint.x = colIdx
				startPoint.y = rowIdx
				startPoint.char = char
			}
			row = append(row, char)
		}
		maze = append(maze, row)
	}
	queue := make(PointQueue, 0)
	visited := make(Set)
	mazeWidth := len(maze[0])
	maxDistance := -1
	queue.Push(startPoint)
	for !queue.IsEmpty() {
		currentPoint := queue.Pop()
		cellId := calculateCellId(currentPoint.x, currentPoint.y, mazeWidth)
		visited.Add(cellId)
		if currentPoint.distance > maxDistance {
			maxDistance = currentPoint.distance
		}
		if currentPoint.x > 0 {
			leftCellId := calculateCellId(currentPoint.x-1, currentPoint.y, mazeWidth)
			if !visited.Contains(leftCellId) {
				if currentPoint.char == 'S' || currentPoint.char == '-' || currentPoint.char == '7' || currentPoint.char == 'J' {
					nextChar := maze[currentPoint.y][currentPoint.x-1]
					if isValidPipe(nextChar, "east") {
						leftPoint := Point{x: currentPoint.x - 1, y: currentPoint.y, distance: currentPoint.distance + 1, char: nextChar}
						queue.Push(leftPoint)
					}
				}
			}
		}
		if currentPoint.x < mazeWidth-1 {
			rightCellId := calculateCellId(currentPoint.x+1, currentPoint.y, mazeWidth)
			if !visited.Contains(rightCellId) {
				if currentPoint.char == 'S' || currentPoint.char == '-' || currentPoint.char == 'L' || currentPoint.char == 'F' {
					nextChar := maze[currentPoint.y][currentPoint.x+1]
					if isValidPipe(nextChar, "west") {
						rightPoint := Point{x: currentPoint.x + 1, y: currentPoint.y, distance: currentPoint.distance + 1, char: nextChar}
						queue.Push(rightPoint)
					}
				}
			}
		}
		if currentPoint.y > 0 {
			topCellId := calculateCellId(currentPoint.x, currentPoint.y-1, mazeWidth)
			if !visited.Contains(topCellId) {
				if currentPoint.char == 'S' || currentPoint.char == '|' || currentPoint.char == 'L' || currentPoint.char == 'J' {
					nextChar := maze[currentPoint.y-1][currentPoint.x]
					if isValidPipe(nextChar, "south") {
						topPoint := Point{x: currentPoint.x, y: currentPoint.y - 1, distance: currentPoint.distance + 1, char: nextChar}
						queue.Push(topPoint)
					}
				}
			}
		}
		if currentPoint.y < len(maze)-1 {
			bottomCellId := calculateCellId(currentPoint.x, currentPoint.y+1, mazeWidth)
			if !visited.Contains(bottomCellId) {
				if currentPoint.char == 'S' || currentPoint.char == '|' || currentPoint.char == '7' || currentPoint.char == 'F' {
					nextChar := maze[currentPoint.y+1][currentPoint.x]
					if isValidPipe(nextChar, "north") {
						bottomPoint := Point{x: currentPoint.x, y: currentPoint.y + 1, distance: currentPoint.distance + 1, char: nextChar}
						queue.Push(bottomPoint)
					}
				}
			}
		}
	}
	fmt.Println(maxDistance)
}

func calculateCellId(x int, y int, width int) int {
	return y*width + x
}

func isValidPipe(char rune, inDir string) bool {

	if char == '|' && (inDir == "north" || inDir == "south") {
		return true
	}

	if char == '-' && (inDir == "east" || inDir == "west") {
		return true
	}

	if char == 'L' && (inDir == "north" || inDir == "east") {
		return true
	}

	if char == 'J' && (inDir == "north" || inDir == "west") {
		return true
	}

	if char == '7' && (inDir == "south" || inDir == "west") {
		return true
	}

	if char == 'F' && (inDir == "south" || inDir == "east") {
		return true
	}

	return false
}
