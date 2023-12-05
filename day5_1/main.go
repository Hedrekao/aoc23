package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	start     int
	end       int
	startDest int
}

func main() {
	content, err := utils.LoadTextFile("data/input5.txt")

	if err != nil {
		panic(err)
	}
	items := []string{"seeds", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

	lines := strings.Split(content, "\n\r\n")

	seeds := make([]int, 0)
	seedsValues := strings.Split(strings.Split(strings.TrimSpace(lines[0]), ": ")[1], " ")

	for _, seed := range seedsValues {
		seedId, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedId)
	}

	dictionary := make(map[string][]Range)
	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		dictType := items[idx]
		rangeArr := handleLine(line)
		dictionary[dictType] = rangeArr
	}

	locations := make([]int, 0)
	for _, seed := range seeds {
		currentId := seed
		for _, item := range items {
			if item == "seeds" {
				continue
			}
			ranges := dictionary[item]
			for _, rangeItem := range ranges {
				fmt.Printf("%s: %d - %d\n", item, rangeItem.start, rangeItem.end)
				if rangeItem.start <= currentId && rangeItem.end >= currentId {
					currentId = rangeItem.startDest + (currentId - rangeItem.start)
					break
				}
			}
			fmt.Println(currentId)
		}
		fmt.Println("------")
		locations = append(locations, currentId)
	}
	minLocation := locations[0]
	for _, location := range locations {
		if location < minLocation {
			minLocation = location
		}
	}
	fmt.Println(minLocation)
}

func handleLine(line string) []Range {
	rows := strings.Split(line, "\r\n")
	result := make([]Range, 0)
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		values := strings.Split(row, " ")
		start, _ := strconv.Atoi(values[1])
		end, _ := strconv.Atoi(strings.TrimSpace(values[2]))
		end = start - 1 + end
		startDest, _ := strconv.Atoi(values[0])
		rangeItem := Range{
			start:     start,
			end:       end,
			startDest: startDest,
		}
		result = append(result, rangeItem)
	}
	return result
}
