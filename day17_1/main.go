package main

import (
	"aoc2023/utils"
	"container/heap"
	"fmt"
	"strings"
)

type Position struct {
	x        int
	y        int
	dx       int
	dy       int
	counter  int
	distance int
}

type PriorityQueue []*Position

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(p interface{}) {
	position := p.(*Position)
	*pq = append(*pq, position)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	position := old[n-1]
	*pq = old[0 : n-1]
	return position
}

func main() {
	content, err := utils.LoadTextFile("data/input17.txt")

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

	value := dijkstra(grid)
	fmt.Println(value)

}

func dijkstra(grid [][]rune) int {
	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	visited := make(map[string]bool)

	start := Position{0, 0, 1, 0, 1, 0}
	endX := len(grid[0]) - 1
	endY := len(grid) - 1

	heap.Push(&priorityQueue, &start)

	for priorityQueue.Len() > 0 {

		current := heap.Pop(&priorityQueue).(*Position)
		if current.x == endX && current.y == endY {
			return current.distance
		}
		cacheKey := getCachedKey(current.x, current.y, current.counter, current.dx, current.dy)

		if visited[cacheKey] {
			continue
		}

		visited[cacheKey] = true
		if current.counter < 3 {
			nextX := current.x + current.dx
			nextY := current.y + current.dy
			nextDistance := getDistanceValue(grid, nextX, nextY)
			if nextDistance > 0 {
				next := Position{nextX, nextY, current.dx, current.dy, current.counter + 1, current.distance + nextDistance}
				heap.Push(&priorityQueue, &next)
			}
		}

		if current.dx != 0 {
			nextX := current.x
			nextY := current.y + 1
			nextDistance := getDistanceValue(grid, nextX, nextY)
			if nextDistance > 0 {
				next := Position{nextX, nextY, 0, 1, 1, current.distance + nextDistance}
				heap.Push(&priorityQueue, &next)
			}
			nextX = current.x
			nextY = current.y - 1
			nextDistance = getDistanceValue(grid, nextX, nextY)
			if nextDistance > 0 {
				next := Position{nextX, nextY, 0, -1, 1, current.distance + nextDistance}
				heap.Push(&priorityQueue, &next)
			}
		} else if current.dy != 0 {
			nextX := current.x + 1
			nextY := current.y
			nextDistance := getDistanceValue(grid, nextX, nextY)
			if nextDistance > 0 {
				next := Position{nextX, nextY, 1, 0, 1, current.distance + nextDistance}
				heap.Push(&priorityQueue, &next)
			}
			nextX = current.x - 1
			nextY = current.y
			nextDistance = getDistanceValue(grid, nextX, nextY)
			if nextDistance > 0 {
				next := Position{nextX, nextY, -1, 0, 1, current.distance + nextDistance}
				heap.Push(&priorityQueue, &next)
			}
		}
	}

	return -1
}

func getDistanceValue(grid [][]rune, x int, y int) int {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return -1
	}

	return int(grid[y][x] - '0')
}

func getCachedKey(x int, y int, dx int, dy int, counter int) string {
	return fmt.Sprintf("%d-%d-%d-%d-%d", x, y, counter, dx, dy)
}
