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

	seedsValues := strings.Split(strings.Split(strings.TrimSpace(lines[0]), ": ")[1], " ")

	dictionary := make(map[string][]Range)
	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		dictType := items[idx]
		rangeArr := handleLine(line)
		dictionary[dictType] = rangeArr
	}

	locationId := 0
	for {
		currentId := locationId
		for i := len(items) - 1; i > 0; i-- {
			item := items[i]
			if item == "seeds" {
				continue
			}
			ranges := dictionary[item]
			for _, rangeItem := range ranges {
				if rangeItem.start <= currentId && rangeItem.end >= currentId {
					currentId = rangeItem.startDest + (currentId - rangeItem.start)
					break
				}
			}

		}

		if validSeed(seedsValues, currentId) {
			break
		}
		locationId++
	}
	fmt.Println(locationId)
}

func handleLine(line string) []Range {
	rows := strings.Split(line, "\r\n")
	result := make([]Range, 0)
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		values := strings.Split(row, " ")
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(strings.TrimSpace(values[2]))
		end = start - 1 + end
		startDest, _ := strconv.Atoi(values[1])
		rangeItem := Range{
			start:     start,
			end:       end,
			startDest: startDest,
		}
		result = append(result, rangeItem)
	}
	return result
}

func validSeed(seeds []string, locationId int) bool {
	for i := 0; i < len(seeds); i += 2 {
		startSeedRangeStr := seeds[i]
		endSeedRangeStr := seeds[i+1]

		startSeedRange, _ := strconv.Atoi(startSeedRangeStr)
		endSeedRange, _ := strconv.Atoi(endSeedRangeStr)

		endSeedRange = startSeedRange - 1 + endSeedRange

		if locationId >= startSeedRange && locationId <= endSeedRange {
			return true
		}
	}
	return false
}
