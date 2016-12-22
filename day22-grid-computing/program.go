package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var numGoRoutines = 4

func main() {
	nodes := parseNodes(getFileLines("df.txt")[2:])
	nodeChannel := make(chan node, len(nodes))
	sumChannel := make(chan int, numGoRoutines*2)
	for i := 0; i < numGoRoutines; i++ {
		go determineViablePairs(nodeChannel, nodes, sumChannel)
	}
	receivedValues := 0
	for _, node := range nodes {
		nodeChannel <- node
	}
	close(nodeChannel)
	sum := 0
	for {
		viablePairs := <-sumChannel
		receivedValues++
		sum += viablePairs
		if receivedValues == len(nodes) {
			break
		}
	}
	fmt.Println("Total number of viable pairs:", sum)
	fmt.Println("Part 2")
	nodeMap := make(map[coordinate]node)
	maxX, maxY := 0, 0
	for _, node := range nodes {
		nodeMap[node.location] = node
		maxX = max(maxX, node.location.x)
		maxY = max(maxY, node.location.y)
	}
	printGrid(nodeMap, maxX, maxY)
	// proceed to analysis on paper after dumping grid.
}

func max(one, other int) int {
	if one < other {
		return other
	}
	return one
}

type coordinate struct {
	x int
	y int
}

type node struct {
	location       coordinate
	size           int
	used           int
	available      int
	usedPercentage int
}

func printGrid(nodeMap map[coordinate]node, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			node := nodeMap[coordinate{x, y}]
			fmt.Printf("%03d/%03d ", node.used, node.size)
		}
		fmt.Println()
	}
}

func determineViablePairs(nodeChannel chan node, nodes []node, sumChannel chan int) {
	for {
		nodeToProcess, ok := <-nodeChannel
		if !ok {
			break
		}
		sumChannel <- determineViablePairsForNode(nodeToProcess, nodes)
	}
}

func determineViablePairsForNode(node node, nodes []node) int {
	viablePairs := 0
	if node.used == 0 {
		return viablePairs
	}
	for _, n := range nodes {
		if node.location == n.location {
			continue
		}
		if node.used <= n.available {
			viablePairs++
		}
	}
	return viablePairs
}

var dfLineRegexp = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)

func parseNodes(nodeLines []string) (result []node) {
	result = make([]node, len(nodeLines))
	for i, line := range nodeLines {
		matches := dfLineRegexp.FindStringSubmatch(line)
		if matches == nil {
			panic(fmt.Sprintf("Line %v not understood", line))
		}
		x := mustParseInt(matches[1])
		y := mustParseInt(matches[2])
		result[i] = node{
			coordinate{x, y},
			mustParseInt(matches[3]),
			mustParseInt(matches[4]),
			mustParseInt(matches[5]),
			mustParseInt(matches[6]),
		}
	}
	return
}

func mustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func getFileLines(fileName string) []string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\r\n")
}
