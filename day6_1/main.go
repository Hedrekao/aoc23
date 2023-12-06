package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content, err := utils.LoadTextFile("data/input6.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\r\n")

	product := 0

	timeLine := lines[0]
	distLine := lines[1]

	timesStr := strings.Split(strings.TrimSpace(strings.Split(timeLine, ":")[1]), " ")
	filteredTimesStr := make([]string, 0)
	for _, str := range timesStr {
		if str != "" {
			filteredTimesStr = append(filteredTimesStr, str)
		}
	}

	times := make([]int, len(filteredTimesStr))

	distsStr := strings.Split(strings.TrimSpace(strings.Split(distLine, ":")[1]), " ")
	filteredDistsStr := make([]string, 0)
	for _, str := range distsStr {
		if str != "" {
			filteredDistsStr = append(filteredDistsStr, str)
		}
	}

	dists := make([]int, len(filteredDistsStr))
	for idx, timeStr := range filteredTimesStr {
		time, _ := strconv.Atoi(timeStr)
		times = append(times, time)
		dist, _ := strconv.Atoi(filteredDistsStr[idx])
		dists = append(dists, dist)
	}

	for i := 0; i < len(times); i++ {
		time := times[i]
		dist := dists[i]
		n_wins := 0
		for j := 0; j <= time; j++ {
			speed := j
			left_time := time - j
			achived_dist := speed * left_time
			if achived_dist > dist {
				n_wins++
			}
		}
		if product == 0 {
			product = n_wins
		} else {
			product *= n_wins
		}
	}

	fmt.Println(product)
}
