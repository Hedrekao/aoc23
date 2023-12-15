package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

var cache = make(map[string]int)

func main() {
	content, err := utils.LoadTextFile("data/input15.txt")

	if err != nil {
		panic(err)
	}
	sum := 0

	line := strings.Split(content, "\n")

	for _, value := range strings.Split(line[0], ",") {
		sum += doAlgo(value)
	}

	fmt.Println(sum)
}

func doAlgo(value string) int {
	result := 0
	cacheKey := ""

	for _, char := range value {
		cacheKey += string(char)
		if val, ok := cache[cacheKey]; ok {
			result = val
			continue
		}
		result += int(char)
		result *= 17
		result %= 256
		cache[cacheKey] = result
	}

	return result
}
