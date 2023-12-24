package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
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
	start := Point{1, 0}
	end := Point{len(lines[0]) - 2, len(lines) - 1}
	junctions := getJunctions(lines)
	junctions[start] = true
	junctions[end] = true

	paths := getPaths(lines, junctions)
	visited := make([]bool, len(junctions))
	visited[paths[start][0].index] = true
	result := findLongestPath(lines, paths, start, end, 0, visited)

	fmt.Println(result)

}

func getJunctions(grid []string) map[Point]bool {
	junctions := map[Point]bool{}
	for row, line := range grid {
		for col, char := range line {
			if char == '#' {
				continue
			}
			point := Point{col, row}
			neighbours := 0
			for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				next := Point{point.x + dir.x, point.y + dir.y}
				if insideGrid(grid, next) && grid[next.y][next.x] != '#' {
					neighbours++
				}
			}
			if neighbours > 2 {
				junctions[point] = true
			}
		}
	}
	return junctions
}

type PathTo struct {
	end           Point
	length, index int
}

func getPaths(grid []string, junctions map[Point]bool) map[Point][]PathTo {
	paths := map[Point][]PathTo{}
	junctionIndex := 0
	for junctionPoint := range junctions {
		for _, startDir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			currentPoint := Point{junctionPoint.x + startDir.x, junctionPoint.y + startDir.y}
			if insideGrid(grid, currentPoint) && grid[currentPoint.y][currentPoint.x] != '#' {
				path := getPath(grid, junctionPoint, currentPoint, startDir, 1, junctions)
				path.index = junctionIndex
				paths[junctionPoint] = append(paths[junctionPoint], path)
			}
		}
		junctionIndex++
	}
	return paths
}

func getPath(grid []string, pathStart, currentPoint, currentDir Point, pathLength int, junctions map[Point]bool) PathTo {
	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{currentPoint.x + dir.x, currentPoint.y + dir.y}
		if grid[next.y][next.x] != '#' {
			if _, found := junctions[next]; found {
				return PathTo{next, pathLength + 1, 0}
			} else {
				return getPath(grid, pathStart, next, dir, pathLength+1, junctions)
			}
		}
	}
	return PathTo{Point{-1, -1}, 0, 0}
}

func findLongestPath(grid []string, paths map[Point][]PathTo, start, end Point, step int, visited []bool) int {
	maxStep := 0
	for _, path := range paths[start] {
		index := paths[path.end][0].index
		if !visited[index] {
			if path.end == end {
				return step + path.length
			}
			visited[index] = true
			maxStep = max(maxStep, findLongestPath(grid, paths, path.end, end, step+path.length, visited))
			visited[index] = false
		}
	}
	return maxStep
}

func dirLeft(p Point) Point {
	return Point{p.y, -p.x}
}

func dirRight(p Point) Point {
	return Point{-p.y, p.x}
}

func insideGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}
