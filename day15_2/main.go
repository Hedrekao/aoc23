package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

var hashMap = make(map[int]Inside)

type Inside struct {
	lenses []Lens
	boxNum int
}

type Lens struct {
	label string
	power int
}

func main() {
	content, err := utils.LoadTextFile("data/input15.txt")

	if err != nil {
		panic(err)
	}
	sum := 0

	line := strings.Split(content, "\n")

	for _, value := range strings.Split(line[0], ",") {
		if value[len(value)-1] == '-' {
			hashedKey := doAlgo(value[:len(value)-1])
			hashedValue, ok := hashMap[hashedKey]
			if ok {
				if len(hashedValue.lenses) > 0 {
					for idx, lens := range hashedValue.lenses {
						if lens.label == value[:len(value)-1] {
							hashMap[hashedKey] = Inside{
								lenses: append(hashedValue.lenses[:idx], hashedValue.lenses[idx+1:]...),
								boxNum: hashedValue.boxNum,
							}
							break
						}
					}
				}

			}
		} else {
			splitValue := strings.Split(value, "=")
			hashedKey := doAlgo(splitValue[0])
			hashedValue, ok := hashMap[hashedKey]
			if ok {
				power, _ := strconv.Atoi(splitValue[1])
				found := false
				for idx, lens := range hashedValue.lenses {
					if lens.label == splitValue[0] {
						lenses := append(hashedValue.lenses[:idx], Lens{
							label: splitValue[0],
							power: power,
						})
						lenses = append(lenses, hashedValue.lenses[idx+1:]...)
						hashMap[hashedKey] = Inside{
							lenses: lenses,
							boxNum: hashedValue.boxNum,
						}
						found = true
					}
				}
				if !found {
					hashMap[hashedKey] = Inside{
						lenses: append(hashedValue.lenses, Lens{
							label: splitValue[0],
							power: power,
						}),
						boxNum: hashedValue.boxNum,
					}
				}
			} else {
				power, _ := strconv.Atoi(splitValue[1])
				hashMap[hashedKey] = Inside{
					lenses: []Lens{
						{
							label: splitValue[0],
							power: power,
						},
					},
					boxNum: hashedKey + 1,
				}

			}

		}
	}

	for _, value := range hashMap {
		if len(value.lenses) > 0 {
			for idx, lens := range value.lenses {
				sum += lens.power * (idx + 1) * value.boxNum
			}
		}
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
