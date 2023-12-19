package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Condition struct {
	identifier string
	condition  string
	number     int
	ifTrue     string
}

type ExtendedCondition struct {
	conditions []Condition
	final      string
}

type Interval struct {
	min int
	max int
}

func getValue(interval Interval) int {
	return interval.max - interval.min + 1
}

var conditions = make(map[string]ExtendedCondition)

func main() {
	content, err := utils.LoadTextFile("data/input19.txt")

	if err != nil {
		panic(err)
	}

	sum := 0
	parts := strings.Split(content, "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		items := strings.Split(line, "{")
		name := strings.TrimSpace(items[0])
		conditionStr := strings.TrimSpace(items[1][:len(items[1])-1])
		conditionsArr := make([]Condition, 0)

		for _, condition := range strings.Split(conditionStr, ",")[:len(strings.Split(conditionStr, ","))-1] {
			conditionData := parseCondition(condition)
			conditionsArr = append(conditionsArr, conditionData)
		}

		final := strings.Split(conditionStr, ",")[len(strings.Split(conditionStr, ","))-1]

		extended := ExtendedCondition{
			conditions: conditionsArr,
			final:      final,
		}
		conditions[name] = extended

	}

	intervals := []Interval{
		{min: 1, max: 4000},
		{min: 1, max: 4000},
		{min: 1, max: 4000},
		{min: 1, max: 4000},
	}

	sum = getAllPossibilities("in", intervals)

	fmt.Println(sum)
}

func getAllPossibilities(name string, intervals []Interval) int {
	if name == "A" {
		sum := 1
		for _, interval := range intervals {
			sum *= getValue(interval)
		}
		return sum
	}

	if name == "R" {
		return 0
	}

	currentConditions := conditions[name]
	count := 0

	for _, condition := range currentConditions.conditions {

		intervalsCopy := make([]Interval, len(intervals))
		copy(intervalsCopy, intervals)

		idx := getIntervalIdxBasedOnIdent(condition.identifier)

		if condition.condition == "<" && intervalsCopy[idx].min < condition.number {

			if intervalsCopy[idx].max > condition.number {
				intervalsCopy[idx].max = condition.number - 1
				intervals[idx].min = condition.number
			}
			count += getAllPossibilities(condition.ifTrue, intervalsCopy)
		} else if condition.condition == ">" && intervalsCopy[idx].max > condition.number {
			if intervalsCopy[idx].min < condition.number {
				intervalsCopy[idx].min = condition.number + 1
				intervals[idx].max = condition.number
			}
			count += getAllPossibilities(condition.ifTrue, intervalsCopy)
		}
	}

	count += getAllPossibilities(currentConditions.final, intervals)

	return count

}

func getIntervalIdxBasedOnIdent(identifier string) int {
	switch identifier {
	case "x":
		return 0
	case "m":
		return 1
	case "a":
		return 2
	case "s":
		return 3
	}
	return -1
}

func parseCondition(condition string) Condition {
	var c Condition
	if strings.Contains(condition, "<") {
		items := strings.Split(condition, "<")
		parts := strings.Split(items[1], ":")
		number, _ := strconv.Atoi(parts[0])
		c = Condition{
			identifier: items[0],
			condition:  "<",
			number:     number,
			ifTrue:     parts[1],
		}

	} else if strings.Contains(condition, ">") {
		items := strings.Split(condition, ">")
		parts := strings.Split(items[1], ":")
		number, _ := strconv.Atoi(parts[0])
		c = Condition{
			identifier: items[0],
			condition:  ">",
			number:     number,
			ifTrue:     parts[1],
		}

	}
	return c
}
