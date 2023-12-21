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
}
type Queue []Point

func (q *Queue) Push(n Point) {
	*q = append(*q, n)
}

func (q *Queue) Pop() Point {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

var rocksMap = make(map[string]bool)

func main() {
	content, err := utils.LoadTextFile("data/input21.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")
	start := Point{0, 0, 0}
	grid := make([][]byte, len(lines))
	orgLen := len(lines[0])
	halfLen := orgLen / 2

	for i := 0; i < len(lines); i++ {
		grid[i] = make([]byte, len(lines[i]))
		row := lines[i]
		for j := 0; j < len(row); j++ {
			grid[i][j] = row[j]
			if row[j] == 'S' {
				start.x = j
				start.y = i
			}
		}
	}

	start.x += 2 * len(grid[0])
	start.y += 2 * len(grid)
	grid = expandGrid(grid)
	fmt.Println(start)

	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			if row[j] == '#' {
				rocksMap[getCellId(j, i)] = true
			}
		}
	}

	firstValue := bfs(start, len(grid), len(grid[0]), halfLen)
	secondValue := bfs(start, len(grid), len(grid[0]), halfLen+orgLen)
	thirdValue := bfs(start, len(grid), len(grid[0]), halfLen+orgLen*2)

	a, b, c := getQuadraticParams(firstValue, secondValue, thirdValue)

	finalX := 26501365 / orgLen

	result := a*finalX*finalX + b*finalX + c

	fmt.Println(result)

}

func getQuadraticParams(y1 int, y2 int, y3 int) (int, int, int) {

	a := (y3 + y1 - 2*y2) / 2
	b := y2 - y1 - a
	c := y1

	return a, b, c
}

func expandGrid(grid [][]byte) [][]byte {
	rowLen := len(grid)
	colLen := len(grid[0])

	for i := 0; i < rowLen; i++ {
		row := make([]byte, colLen)
		copy(row, grid[i])

		tmpRow := append(row, row...)
		tmpRow = append(tmpRow, tmpRow...)
		grid[i] = append(grid[i], tmpRow...)
	}

	orgGrid := make([][]byte, len(grid))
	copy(orgGrid, grid)
	grid = append(grid, orgGrid...)
	grid = append(grid, orgGrid...)
	grid = append(grid, orgGrid...)
	grid = append(grid, orgGrid...)

	fmt.Println(len(grid), len(grid[0]))

	return grid
}

func getCellId(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func getExtendedCellId(x int, y int, distance int) string {
	return fmt.Sprintf("%d-%d-%d", x, y, distance)
}

func bfs(startPoint Point, rowLen int, colLen int, maxDistance int) int {
	queue := make(Queue, 0)
	queue.Push(startPoint)

	count := 0
	penultimateMap := make(map[string]bool)
	visitedMap := make(map[string]bool)

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.distance == maxDistance {
			if !penultimateMap[getCellId(current.x, current.y)] {
				count++
				penultimateMap[getCellId(current.x, current.y)] = true
			}
			continue
		}

		if visitedMap[getExtendedCellId(current.x, current.y, current.distance)] {
			continue
		}

		visitedMap[getExtendedCellId(current.x, current.y, current.distance)] = true

		if current.x < rowLen-1 && !rocksMap[getCellId(current.x+1, current.y)] {
			queue.Push(Point{current.x + 1, current.y, current.distance + 1})
		}
		if current.x > 0 && !rocksMap[getCellId(current.x-1, current.y)] {
			queue.Push(Point{current.x - 1, current.y, current.distance + 1})
		}
		if current.y < colLen-1 && !rocksMap[getCellId(current.x, current.y+1)] {
			queue.Push(Point{current.x, current.y + 1, current.distance + 1})
		}
		if current.y > 0 && !rocksMap[getCellId(current.x, current.y-1)] {
			queue.Push(Point{current.x, current.y - 1, current.distance + 1})
		}

	}

	return count

}
