package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

var exampleMaze = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

type coordinate struct {
	x int
	y int
}

type maze struct {
	positions map[coordinate]string
	targets   []coordinate
}

type leg struct {
	from int
	to   int
}

type pathResult struct {
	coordinate coordinate
	length     int
}

func main() {
	exampleMaze := readMaze(exampleMaze, 5)
	determineShortestPathThroughAll(exampleMaze)
	maze := readMaze(readAirductsFile(), 8)
	determineShortestPathThroughAll(maze)
}

func determineShortestPathThroughAll(maze maze) {
	/* First, determine distance between all pairs. */
	distances := determineDistances(maze)
	/* Then, find optimal path based on shortest distance between all pairs. */
	nodesVisited := make([]bool, len(maze.targets))
	shortest := determineShortestPath(0, nodesVisited, distances, false)
	fmt.Printf("Shortest path through all is %v steps\n", shortest)
	shortestWithReturn := determineShortestPath(0, nodesVisited, distances, true)
	fmt.Printf("Shortest path through all and returning is %v steps\n", shortestWithReturn)
}

func determineShortestPath(from int, nodesVisited []bool, distances map[leg]int, includeReturn bool) int {
	visited := setVisited(nodesVisited, from)
	if allVisited(visited) {
		if includeReturn {
			return getDistance(distances, 0, from)
		}
		return 0
	}
	shortest := math.MaxInt32
	for i := range visited {
		if visited[i] {
			continue
		}
		pathLength := getDistance(distances, from, i) + determineShortestPath(i, visited, distances, includeReturn)
		if shortest <= pathLength {
			continue
		}
		shortest = pathLength
	}
	return shortest
}

func setVisited(nodes []bool, nodeIndex int) []bool {
	result := make([]bool, len(nodes))
	copy(result, nodes)
	result[nodeIndex] = true
	return result
}

func allVisited(nodesVisited []bool) bool {
	for _, visited := range nodesVisited {
		if !visited {
			return false
		}
	}
	return true
}

func getDistance(distances map[leg]int, from, to int) int {
	if from < to {
		return distances[leg{from, to}]
	}
	return distances[leg{to, from}]
}

func determineDistances(maze maze) map[leg]int {
	distances := make(map[leg]int)
	for i, source := range maze.targets {
		for j := i + 1; j < len(maze.targets); j++ {
			destination := maze.targets[j]
			distances[leg{i, j}] = determineDistance(maze, source, destination)
		}
	}
	return distances
}

/* Breadth first search is fairly inefficient for this maze size... but it will do. */
func determineDistance(maze maze, source, destination coordinate) int {
	fmt.Println("Determining distance between", source, "and", destination)
	visited := make(map[coordinate]bool)
	toVisit := []pathResult{pathResult{source, 0}}
	for {
		next := toVisit[0]
		toVisit = toVisit[1:]
		visited[next.coordinate] = true
		for _, neighbor := range destinations(next.coordinate) {
			if neighbor == destination {
				fmt.Println("Found destination at", 1+next.length)
				return 1 + next.length
			}
			if visited[neighbor] || maze.positions[neighbor] == "#" {
				continue
			}
			toVisit = append(toVisit, pathResult{neighbor, 1 + next.length})
		}
	}
}

func destinations(c coordinate) []coordinate {
	return []coordinate{
		coordinate{c.x, c.y + 1},
		coordinate{c.x, c.y - 1},
		coordinate{c.x + 1, c.y},
		coordinate{c.x - 1, c.y},
	}
}

func readAirductsFile() string {
	bytes, err := ioutil.ReadFile("airducts.txt")
	check(err)
	return string(bytes)
}

func readMaze(layout string, numberOfTargets int) maze {
	positions := make(map[coordinate]string)
	targets := make([]coordinate, numberOfTargets)
	lines := strings.Split(layout, "\n")
	for y, line := range lines {
		exploded := strings.Split(line, "")
		for x, cell := range exploded {
			if cell == "\r" {
				continue
			}
			position := coordinate{x, y}
			positions[position] = cell
			if !isNumeric(cell) {
				continue
			}
			targetID := mustParseInt(cell)
			targets[targetID] = position
		}
	}
	return maze{positions, targets}
}

var numRegexp = regexp.MustCompile(`\d+`)

func isNumeric(s string) bool {
	return numRegexp.MatchString(s)
}

func mustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	check(err)
	return int(i)
}
