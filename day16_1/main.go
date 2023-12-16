package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Position struct {
	x  int
	y  int
	dx int
	dy int
}

type Queue []Position

func (q *Queue) Push(n Position) {
	*q = append(*q, n)
}

func (q *Queue) Pop() Position {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

var positionMap = make(map[Position]bool)

func main() {
	content, err := utils.LoadTextFile("data/input16.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")
	grid := make([][]rune, len(lines))

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]rune, len(line))
		for j := 0; j < len(line); j++ {
			row[j] = rune(line[j])
		}
		grid[i] = row
	}

	start := Position{0, 0, 1, 0}
	bfs(grid, start)
	energized := countEnergized(grid)
	fmt.Println(energized)

}

func countEnergized(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			if positionMap[Position{j, i, 0, 0}] ||
				positionMap[Position{j, i, 1, 0}] ||
				positionMap[Position{j, i, -1, 0}] ||
				positionMap[Position{j, i, 0, 1}] ||
				positionMap[Position{j, i, 0, -1}] {
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]rune, start Position) {
	queue := make(Queue, 0)
	queue = append(queue, start)

	for !queue.IsEmpty() {
		current := queue.Pop()

		if positionMap[current] {
			continue
		}

		positionMap[current] = true

		currentChar := grid[current.y][current.x]
		if currentChar == '.' {
			nextX := current.x + current.dx
			nextY := current.y + current.dy
			if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
				nextPosition := Position{nextX, nextY, current.dx, current.dy}
				queue.Push(nextPosition)
			}
		} else if currentChar == '-' {
			if current.dx == 1 || current.dx == -1 {
				nextX := current.x + current.dx
				nextY := current.y + current.dy
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					nextPosition := Position{nextX, nextY, current.dx, current.dy}
					queue.Push(nextPosition)
				}
			} else {
				nextX := current.x + 1
				nextY := current.y
				nextPosition := Position{nextX, nextY, 1, 0}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}

				nextX = current.x - 1
				nextY = current.y
				nextPosition = Position{nextX, nextY, -1, 0}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			}
		} else if currentChar == '|' {
			if current.dy == 1 || current.dy == -1 {
				nextX := current.x + current.dx
				nextY := current.y + current.dy
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					nextPosition := Position{nextX, nextY, current.dx, current.dy}
					queue.Push(nextPosition)
				}
			} else {
				nextX := current.x
				nextY := current.y + 1
				nextPosition := Position{nextX, nextY, 0, 1}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}

				nextX = current.x
				nextY = current.y - 1
				nextPosition = Position{nextX, nextY, 0, -1}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			}
		} else if currentChar == '/' {
			if current.dx == 1 {
				nextX := current.x
				nextY := current.y - 1
				nextPosition := Position{nextX, nextY, 0, -1}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			} else if current.dx == -1 {
				nextX := current.x
				nextY := current.y + 1
				nextPosition := Position{nextX, nextY, 0, 1}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			} else if current.dy == 1 {
				nextX := current.x - 1
				nextY := current.y
				nextPosition := Position{nextX, nextY, -1, 0}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			} else if current.dy == -1 {
				nextX := current.x + 1
				nextY := current.y
				nextPosition := Position{nextX, nextY, 1, 0}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			}
		} else if currentChar == '\\' {
			if current.dx == 1 {
				nextX := current.x
				nextY := current.y + 1
				nextPosition := Position{nextX, nextY, 0, 1}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			} else if current.dx == -1 {
				nextX := current.x
				nextY := current.y - 1
				nextPosition := Position{nextX, nextY, 0, -1}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			} else if current.dy == 1 {
				nextX := current.x + 1
				nextY := current.y
				nextPosition := Position{nextX, nextY, 1, 0}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			} else if current.dy == -1 {
				nextX := current.x - 1
				nextY := current.y
				nextPosition := Position{nextX, nextY, -1, 0}
				if nextX >= 0 && nextX < len(grid[0]) && nextY >= 0 && nextY < len(grid) {
					queue.Push(nextPosition)
				}
			}
		}

	}

}
