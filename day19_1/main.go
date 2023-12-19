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

type LineContent struct {
	x int
	m int
	a int
	s int
}

var conditions = make(map[string][]Condition)

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

		for _, condition := range strings.Split(conditionStr, ",") {
			conditionData := parseCondition(condition)
			conditionsArr = append(conditionsArr, conditionData)
		}

		conditions[name] = conditionsArr

	}

	for _, line := range strings.Split(parts[1], "\n") {
		line = line[1 : len(line)-1]
		lineContent := LineContent{}
		for idx, value := range strings.Split(line, ",") {
			number, _ := strconv.Atoi(strings.Split(value, "=")[1])
			switch idx {
			case 0:
				lineContent.x = number
			case 1:
				lineContent.m = number
			case 2:
				lineContent.a = number
			case 3:
				lineContent.s = number
			}
		}
		isDone := false
		isDoneInner := false
		currentConditions := conditions["in"]

		for !isDone {
			isDoneInner = false
			for _, condition := range currentConditions {
				if isDoneInner {
					break
				}
				if condition.identifier == "" {
					if condition.ifTrue == "A" {
						sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
						isDone = true
					} else if condition.ifTrue == "R" {
						isDone = true
					} else {
						currentConditions = conditions[condition.ifTrue]
					}
					break
				}

				switch condition.condition {
				case "<":
					switch condition.identifier {
					case "x":
						if lineContent.x < condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					case "m":
						if lineContent.m < condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					case "a":
						if lineContent.a < condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					case "s":
						if lineContent.s < condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					}
				case ">":
					switch condition.identifier {
					case "x":
						if lineContent.x > condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					case "m":
						if lineContent.m > condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					case "a":
						if lineContent.a > condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					case "s":
						if lineContent.s > condition.number {
							if condition.ifTrue == "A" {
								sum += lineContent.a + lineContent.s + lineContent.m + lineContent.x
								isDone = true
								isDoneInner = true
							} else if condition.ifTrue == "R" {
								isDone = true
								isDoneInner = true
							} else {
								currentConditions = conditions[condition.ifTrue]
								isDoneInner = true
							}
						}
					}
				}
			}
		}

	}

	fmt.Println(sum)

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

	} else {
		c = Condition{
			identifier: "",
			condition:  "",
			number:     -1,
			ifTrue:     condition,
		}
	}
	return c
}
