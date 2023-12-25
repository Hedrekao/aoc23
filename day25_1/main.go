package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func main() {
	content, err := utils.LoadTextFile("data/input25.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(content, "\n")
	g := graphviz.New()
	g.SetLayout(graphviz.NEATO)
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	set := make(map[string]*cgraph.Node)

	connectionsMap := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, ": ")

		var node *cgraph.Node

		if _, ok := set[parts[0]]; !ok {
			node, err = graph.CreateNode(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			set[parts[0]] = node
		} else {
			node = set[parts[0]]
		}

		connections := strings.Split(parts[1], " ")

		for _, connection := range connections {
			var secondNode *cgraph.Node
			if _, ok := set[connection]; !ok {
				secondNode, err = graph.CreateNode(connection)
				if err != nil {
					log.Fatal(err)
				}
				set[connection] = secondNode
			} else {
				secondNode = set[connection]
			}
			graph.CreateEdge("", node, secondNode)

			if _, ok := connectionsMap[connection]; !ok {
				connectionsMap[connection] = []string{parts[0]}
			} else {
				connectionsMap[connection] = append(connectionsMap[connection], parts[0])
			}
		}

		if _, ok := connectionsMap[parts[0]]; !ok {
			connectionsMap[parts[0]] = connections
		} else {
			connectionsMap[parts[0]] = append(connectionsMap[parts[0]], connections...)
		}

	}

	if err := g.RenderFilename(graph, graphviz.SVG, "./day25_1/graph.svg"); err != nil {
		log.Fatal(err)
	}

	connectionsMap = removeConnection(connectionsMap, "prk", "gpz")
	connectionsMap = removeConnection(connectionsMap, "zhg", "qdv")
	connectionsMap = removeConnection(connectionsMap, "rfq", "lsk")

	wg := sync.WaitGroup{}
	wg.Add(2)
	result := 1
	go bfs("xvk", connectionsMap, &result, &wg)
	go bfs("vpj", connectionsMap, &result, &wg)

	wg.Wait()

	fmt.Println(result)

}

func bfs(startNode string, connections map[string][]string, count *int, wg *sync.WaitGroup) {
	queue := []string{startNode}
	visited := make(map[string]bool)

	defer wg.Done()

	localCount := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := visited[node]; !ok {
			visited[node] = true
			localCount++
			queue = append(queue, connections[node]...)
		}
	}

	*count *= localCount

}

func removeConnection(connections map[string][]string, node1 string, node2 string) map[string][]string {

	if _, ok := connections[node1]; !ok {
		return connections
	}

	for i, connection := range connections[node1] {
		if connection == node2 {
			connections[node1] = append(connections[node1][:i], connections[node1][i+1:]...)
		}
	}

	for i, connection := range connections[node2] {
		if connection == node1 {
			connections[node2] = append(connections[node2][:i], connections[node2][i+1:]...)
		}
	}

	return connections

}
