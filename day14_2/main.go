package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Set map[int]bool

func (s *Set) Add(i int) {
	(*s)[i] = true
}

func (s *Set) Remove(i int) {
	delete(*s, i)
}

func (s *Set) Contains(i int) bool {
	_, ok := (*s)[i]
	return ok
}

func main() {
	content, err := utils.LoadTextFile("data/input14.txt")

	if err != nil {
		panic(err)
	}
	sum := 0
	cache := make(map[string]int)

	lines := strings.Split(content, "\n")
	idx := 1
	done := false
	for idx <= 1000000000 {
		for i := 0; i < len(lines[0]); i++ {
			lines = tiltVertical(lines, i, 'N')
		}

		for i := 0; i < len(lines); i++ {
			lines = tiltHorizontal(lines, i, 'W')
		}
		for i := 0; i < len(lines[0]); i++ {
			lines = tiltVertical(lines, i, 'S')
		}
		for i := 0; i < len(lines); i++ {
			lines = tiltHorizontal(lines, i, 'E')
		}

		cacheKey := strings.Join(lines, ";")
		cacheValue, ok := cache[cacheKey]
		if !ok {
			cache[cacheKey] = idx
		} else {
			cycle := idx - cacheValue
			newIdx := cacheValue + (1000000000-cacheValue)%cycle
			for k, v := range cache {
				if v == newIdx {
					lines = strings.Split(k, ";")
					sum += calculateScore(lines)
					done = true
					break
				}
			}
		}
		if done {
			break
		}
		idx++
	}
	printGrid(lines)
	fmt.Println(sum)
}

func printGrid(lines []string) {
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}

func calculateScore(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'O' {
				sum += len(lines) - i
			}
		}
	}
	return sum
}

func tiltVertical(lines []string, columnIdx int, direction rune) []string {
	ovalRockArr := make([]int, 0)
	ovalRockSet := make(Set)
	notMovableRockSet := make(Set)
	emptySpaceSet := make(Set)

	for i := 0; i < len(lines); i++ {
		if lines[i][columnIdx] == '.' {
			emptySpaceSet.Add(i)
		} else if lines[i][columnIdx] == 'O' {
			ovalRockArr = append(ovalRockArr, i)
			ovalRockSet.Add(i)
		} else if lines[i][columnIdx] == '#' {
			notMovableRockSet.Add(i)
		}
	}

	if direction == 'S' {
		for i, j := 0, len(ovalRockArr)-1; i < j; i, j = i+1, j-1 {
			ovalRockArr[i], ovalRockArr[j] = ovalRockArr[j], ovalRockArr[i]
		}

	}

	for _, rock := range ovalRockArr {

		position := rock
		if direction == 'N' {
			for i := rock - 1; i >= 0; i-- {
				if notMovableRockSet.Contains(i) || ovalRockSet.Contains(i) {
					break
				} else {
					if emptySpaceSet.Contains(i) {
						prevPosition := position
						position = i
						emptySpaceSet.Remove(i)
						emptySpaceSet.Add(prevPosition)
					}
				}
			}
		}

		if direction == 'S' {
			for i := rock + 1; i < len(lines); i++ {
				if notMovableRockSet.Contains(i) || ovalRockSet.Contains(i) {
					break
				} else {
					if emptySpaceSet.Contains(i) {
						prevPosition := position
						position = i
						emptySpaceSet.Remove(i)
						emptySpaceSet.Add(prevPosition)
					}
				}
			}
		}
		if position == rock {
			continue
		}
		ovalRockSet.Remove(rock)
		emptySpaceSet.Add(rock)
		ovalRockSet.Add(position)

	}

	lines = drawColumn(lines, columnIdx, emptySpaceSet, notMovableRockSet, ovalRockSet)
	return lines
}

func drawColumn(lines []string, columnIdx int, emptySpaceSet Set, notMovableRockSet Set, ovalRockSet Set) []string {
	for i := 0; i < len(lines); i++ {
		if emptySpaceSet.Contains(i) {
			lines[i] = replaceAtIndex(lines[i], columnIdx, '.')
		} else if notMovableRockSet.Contains(i) {
			lines[i] = replaceAtIndex(lines[i], columnIdx, '#')
		} else if ovalRockSet.Contains(i) {
			lines[i] = replaceAtIndex(lines[i], columnIdx, 'O')
		}
	}
	return lines
}

func replaceAtIndex(str string, index int, replacement rune) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func tiltHorizontal(lines []string, rowIdx int, direction rune) []string {
	ovalRockArr := make([]int, 0)
	ovalRockSet := make(Set)
	notMovableRockSet := make(Set)
	emptySpaceSet := make(Set)

	for i := 0; i < len(lines[rowIdx]); i++ {
		if lines[rowIdx][i] == '.' {
			emptySpaceSet.Add(i)
		} else if lines[rowIdx][i] == 'O' {
			ovalRockArr = append(ovalRockArr, i)
			ovalRockSet.Add(i)
		} else if lines[rowIdx][i] == '#' {
			notMovableRockSet.Add(i)
		}
	}

	if direction == 'E' {
		for i, j := 0, len(ovalRockArr)-1; i < j; i, j = i+1, j-1 {
			ovalRockArr[i], ovalRockArr[j] = ovalRockArr[j], ovalRockArr[i]
		}

	}

	for _, rock := range ovalRockArr {

		position := rock
		if direction == 'W' {
			for i := rock - 1; i >= 0; i-- {
				if notMovableRockSet.Contains(i) || ovalRockSet.Contains(i) {
					break
				} else {
					if emptySpaceSet.Contains(i) {
						prevPosition := position
						position = i
						emptySpaceSet.Remove(i)
						emptySpaceSet.Add(prevPosition)
					}
				}
			}
		}

		if direction == 'E' {
			for i := rock + 1; i < len(lines[rowIdx]); i++ {
				if notMovableRockSet.Contains(i) || ovalRockSet.Contains(i) {
					break
				} else {
					if emptySpaceSet.Contains(i) {
						prevPosition := position
						position = i
						emptySpaceSet.Remove(i)
						emptySpaceSet.Add(prevPosition)
					}
				}
			}
		}

		if position == rock {
			continue
		}

		ovalRockSet.Remove(rock)
		emptySpaceSet.Add(rock)
		ovalRockSet.Add(position)
	}

	lines = drawRow(lines, rowIdx, emptySpaceSet, notMovableRockSet, ovalRockSet)
	return lines
}

func drawRow(lines []string, rowIdx int, emptySpaceSet Set, notMovableRockSet Set, ovalRockSet Set) []string {
	for i := 0; i < len(lines[rowIdx]); i++ {
		if emptySpaceSet.Contains(i) {
			lines[rowIdx] = replaceAtIndex(lines[rowIdx], i, '.')
		} else if notMovableRockSet.Contains(i) {
			lines[rowIdx] = replaceAtIndex(lines[rowIdx], i, '#')
		} else if ovalRockSet.Contains(i) {
			lines[rowIdx] = replaceAtIndex(lines[rowIdx], i, 'O')
		}
	}
	return lines
}
