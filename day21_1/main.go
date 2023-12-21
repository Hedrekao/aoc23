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

	for i := 0; i < len(lines); i++ {
		row := lines[i]
		for j := 0; j < len(row); j++ {
			if row[j] == '#' {
				rocksMap[getCellId(j, i)] = true
			}
			if row[j] == 'S' {
				start = Point{j, i, 0}
			}
		}
	}

	rowLen := len(lines[0])
	colLen := len(lines)

	fmt.Println(bfs(start, rowLen, colLen))

}

func getCellId(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func getExtendedCellId(x int, y int, distance int) string {
	return fmt.Sprintf("%d-%d-%d", x, y, distance)
}

func bfs(startPoint Point, rowLen int, colLen int) int {
	queue := make(Queue, 0)
	queue.Push(startPoint)

	count := 0
	maxDistance := 64
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
