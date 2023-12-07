package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

var cardsPower = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func main() {
	content, err := utils.LoadTextFile("data/input7.txt")

	if err != nil {
		panic(err)
	}
	sum := 0
	lines := strings.Split(content, "\r\n")

	hands := make([]Hand, 0)
	for _, line := range lines {

		splitted := strings.Split(line, " ")
		cards := strings.TrimSpace(splitted[0])
		bid, _ := strconv.Atoi(strings.TrimSpace(splitted[1]))
		hands = append(hands, Hand{cards, bid})
	}

	sortedHands := bubbleSort(hands)

	for i, hand := range sortedHands {
		sum += hand.bid * (i + 1)
	}
	fmt.Println(sum)
}

func bubbleSort(arr []Hand) []Hand {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			hand1Power := evalHand(arr[i-1])
			hand2Power := evalHand(arr[i])
			shouldSwap := hand1Power > hand2Power
			if hand1Power == hand2Power {
				shouldSwap = compareHands(arr[i-1], arr[i]) == 1
			}
			if shouldSwap {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func evalHand(hand Hand) int {
	cards := strings.Split(hand.cards, "")
	countMap := make(map[string]int)
	for _, card := range cards {
		countMap[card]++
	}
	maxCount := 0
	maxCard := ""
	numberOfJokers := 0
	for key, count := range countMap {
		if key == "J" {
			numberOfJokers = count
		}
		if count > maxCount {
			maxCount = count
			maxCard = key
		}
	}
	if maxCount == 1 {
		if numberOfJokers == 1 {
			return 1
		}
		return 0
	} else if maxCount == 2 {
		if numberOfJokers == 2 {
			numberOf2s := 0
			for _, count := range countMap {
				if count == 2 {
					numberOf2s++
				}
			}
			if numberOf2s == 2 {
				return 5
			} else {
				return 3
			}
		}
		if numberOfJokers == 1 {
			delete(countMap, maxCard)
			for _, count := range countMap {
				if count == 2 {
					return 4
				}
			}
			return 3
		}
		delete(countMap, maxCard)
		for _, count := range countMap {
			if count == 2 {
				return 2
			}
		}
		return 1
	} else if maxCount == 3 {
		if numberOfJokers == 1 {
			return 5
		}
		if numberOfJokers == 2 {
			return 6
		}
		if numberOfJokers == 3 {
			delete(countMap, maxCard)
			for _, count := range countMap {
				if count == 2 {
					return 6
				}
			}
			return 5
		}
		delete(countMap, maxCard)
		for _, count := range countMap {
			if count == 2 {
				return 4
			}
		}
		return 3
	} else if maxCount == 4 {
		if numberOfJokers > 0 {
			return 6
		}

		return 5
	} else {
		return 6
	}
}

func compareHands(hand1 Hand, hand2 Hand) int {
	cards1 := strings.Split(hand1.cards, "")
	cards2 := strings.Split(hand2.cards, "")
	for i := 0; i < len(cards1); i++ {
		card1 := cards1[i]
		card2 := cards2[i]
		if cardsPower[card1] > cardsPower[card2] {
			return 1
		} else if cardsPower[card1] < cardsPower[card2] {
			return 2
		}
	}
	return 0
}
