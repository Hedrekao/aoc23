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

	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines[0]); i++ {
		sum += calculateColumn(lines, i)
	}

	fmt.Println(sum)
}

func calculateColumn(lines []string, columnIdx int) int {
	sum := 0
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

	for _, rock := range ovalRockArr {

		minPosition := rock
		for i := rock - 1; i >= 0; i-- {
			if notMovableRockSet.Contains(i) || ovalRockSet.Contains(i) {
				break
			} else {
				if emptySpaceSet.Contains(i) {
					prevPosition := minPosition
					minPosition = i
					emptySpaceSet.Remove(i)
					emptySpaceSet.Add(prevPosition)
				}
			}
		}
		ovalRockSet.Remove(rock)
		emptySpaceSet.Add(rock)
		ovalRockSet.Add(minPosition)
		sum += len(lines) - minPosition

	}

	return sum
}
