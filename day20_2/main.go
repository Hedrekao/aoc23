package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type ModuleType int

type Pulse bool

const (
	EMPTY ModuleType = iota
	BROADCAST
	FLIP
	CONJUCTION

	HIGH Pulse = true
	LOW  Pulse = false
)

type Input struct {
	src   string
	dst   string
	pulse Pulse
}

type Queue []Input

func (q *Queue) Push(n Input) {
	*q = append(*q, n)
}

func (q *Queue) Pop() Input {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type Module struct {
	Type     ModuleType
	On       bool
	Memories map[string]Pulse
	Dest     []string
}

var modules = make(map[string]Module)

func main() {
	content, err := utils.LoadTextFile("data/input20.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		moduleTypeStr := parts[0]
		var label string
		var moduleType ModuleType
		switch moduleTypeStr[0:1] {
		case "%":
			moduleType = FLIP
			label = moduleTypeStr[1:]
		case "&":
			moduleType = CONJUCTION
			label = moduleTypeStr[1:]
		default:
			moduleType = EMPTY
			label = moduleTypeStr
			if label == "broadcaster" {
				moduleType = BROADCAST
			}
		}

		dests := strings.Split(parts[1], ", ")

		module, ok := modules[label]

		if !ok {
			module = Module{Type: moduleType, On: false, Memories: make(map[string]Pulse), Dest: dests}
		} else {
			module.Dest = dests
			module.Type = moduleType
		}

		modules[label] = module

		for _, dest := range dests {
			destModule, ok := modules[dest]
			if !ok {
				memories := make(map[string]Pulse)
				memories[label] = LOW
				destModule = Module{Memories: memories}
			} else {
				destModule.Memories[label] = LOW
			}
			modules[dest] = destModule
		}

	}

	buttonModule := Module{Type: EMPTY, On: false, Memories: make(map[string]Pulse), Dest: []string{"broadcaster"}}
	modules["button"] = buttonModule

	final_module := modules["rx"]

	var penultimateModule Module

	for moduleName := range final_module.Memories {
		penultimateModule = modules[moduleName]
	}

	prepenultimateModules := make(map[string]bool)

	for moduleName := range penultimateModule.Memories {
		prepenultimateModules[moduleName] = true
	}

	buttonPressed := 1
	cycles := make([]int, 0)
	for {
		cycledModule := bfs(prepenultimateModules)
		if cycledModule != nil {
			cycles = append(cycles, buttonPressed)
			delete(prepenultimateModules, *cycledModule)
		}

		if len(prepenultimateModules) == 0 {
			break
		}

		buttonPressed++
	}

	fmt.Println(lcm(cycles))

}

func bfs(cycledModuls map[string]bool) *string {
	queue := make(Queue, 0)
	startInput := Input{"button", "broadcaster", LOW}
	queue.Push(startInput)

	lowPulseCount := 0
	highPulseCount := 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		destModule := modules[current.dst]

		if current.pulse == HIGH {
			highPulseCount++
		} else {
			lowPulseCount++
		}

		switch destModule.Type {
		case EMPTY:
			for _, dest := range destModule.Dest {
				queue.Push(Input{current.dst, dest, current.pulse})
			}
		case BROADCAST:
			for _, dest := range destModule.Dest {
				queue.Push(Input{current.dst, dest, current.pulse})
			}
		case FLIP:
			if current.pulse == LOW {
				var newPulse Pulse
				if destModule.On {
					newPulse = LOW
					destModule.On = false
				} else {
					newPulse = HIGH
					destModule.On = true
				}
				modules[current.dst] = destModule
				for _, dest := range destModule.Dest {
					queue.Push(Input{current.dst, dest, newPulse})
				}
			}
		case CONJUCTION:
			destModule.Memories[current.src] = current.pulse
			modules[current.dst] = destModule
			newPulse := LOW
			for _, pulse := range destModule.Memories {
				if pulse == LOW {
					newPulse = HIGH
					break
				}
			}
			for _, dest := range destModule.Dest {
				queue.Push(Input{current.dst, dest, newPulse})
			}

			if _, ok := cycledModuls[current.dst]; ok && newPulse == HIGH {
				result := current.dst
				return &result
			}

		}

	}

	return nil

}

func lcm(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm2(result, numbers[i])
	}
	return result
}

func lcm2(x, y int) int {
	return x * y / gcd2(x, y)
}

func gcd2(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
