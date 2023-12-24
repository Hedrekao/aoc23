package main

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strings"
)

type Point struct {
	x       int
	y       int
	steps   int
	visited []string
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

func main() {
	content, err := utils.LoadTextFile("data/input23.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")
	grid := make([][]byte, len(lines))

	for i := 0; i < len(lines); i++ {
		grid[i] = make([]byte, len(lines[i]))
		row := lines[i]
		for j := 0; j < len(row); j++ {
			grid[i][j] = row[j]
		}
	}

	result := bfs(grid)
	fmt.Println(result)

}

func makeKey(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func bfs(grid [][]byte) int {
	queue := make(Queue, 0)
	startPoint := Point{0, 1, 0, make([]string, 0)}
	queue.Push(startPoint)
	endY := len(grid) - 1
	endX := len(grid[0]) - 2

	results := make([]int, 0)
	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.x == endX && current.y == endY {
			results = append(results, current.steps)
			continue
		}

		currentKey := makeKey(current.x, current.y)

		if slices.Contains(current.visited, currentKey) {
			continue
		}

		current.visited = append(current.visited, currentKey)

		currentTile := grid[current.y][current.x]

		if currentTile == '>' {
			newX := current.x + 1
			newY := current.y
			if newX < len(grid[0]) && grid[newY][newX] != '#' {
				newVisited := make([]string, len(current.visited))
				copy(newVisited, current.visited)
				queue.Push(Point{newX, newY, current.steps + 1, newVisited})
			}
		} else if currentTile == '<' {
			newX := current.x - 1
			newY := current.y
			if newX >= 0 && grid[newY][newX] != '#' {
				newVisited := make([]string, len(current.visited))
				copy(newVisited, current.visited)
				queue.Push(Point{newX, newY, current.steps + 1, newVisited})
			}
		} else if currentTile == '^' {
			newX := current.x
			newY := current.y - 1
			if newY >= 0 && grid[newY][newX] != '#' {
				newVisited := make([]string, len(current.visited))
				copy(newVisited, current.visited)
				queue.Push(Point{newX, newY, current.steps + 1, newVisited})
			}
		} else if currentTile == 'v' {
			newX := current.x
			newY := current.y + 1
			if newY < len(grid) && grid[newY][newX] != '#' {
				newVisited := make([]string, len(current.visited))
				copy(newVisited, current.visited)
				queue.Push(Point{newX, newY, current.steps + 1, newVisited})
			}
		} else {
			neighbours := getValidNeighbours(current.x, current.y, grid, current.steps)

			for _, neighbour := range neighbours {
				newVisited := make([]string, len(current.visited))
				copy(newVisited, current.visited)
				neighbour.visited = newVisited
				queue.Push(neighbour)
			}

		}
	}

	maxResult := 0

	for _, result := range results {
		if result > maxResult {
			maxResult = result
		}
	}

	return maxResult

}

func getValidNeighbours(x int, y int, grid [][]byte, steps int) []Point {

	neighbours := make([]Point, 0)

	if x > 0 && grid[y][x-1] != '#' {
		nextTile := grid[y][x-1]
		if nextTile != '>' {
			neighbours = append(neighbours, Point{x - 1, y, steps + 1, make([]string, 0)})
		}
	}

	if y > 0 && grid[y-1][x] != '#' {
		nextTile := grid[y-1][x]
		if nextTile != 'v' {
			neighbours = append(neighbours, Point{x, y - 1, steps + 1, make([]string, 0)})
		}
	}

	if x < len(grid[0])-1 && grid[y][x+1] != '#' {
		nextTile := grid[y][x+1]
		if nextTile != '<' {
			neighbours = append(neighbours, Point{x + 1, y, steps + 1, make([]string, 0)})
		}
	}

	if y < len(grid)-1 && grid[y+1][x] != '#' {
		nextTile := grid[y+1][x]
		if nextTile != '^' {
			neighbours = append(neighbours, Point{x, y + 1, steps + 1, make([]string, 0)})
		}
	}

	return neighbours
}
