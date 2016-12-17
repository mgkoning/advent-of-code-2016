package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var roomSize = 3
var passcodeExample1 = "ihgpwlah"
var passcodeExample2 = "kglvqrro"
var passcodeExample3 = "ulqzkmiv"
var puzzlePasscode = "rrrbmfta"

func main() {
	fmt.Println("Shortest path example 1:", shortestPath(passcodeExample1))
	// DDRRRD
	fmt.Println("Shortest path example 2:", shortestPath(passcodeExample2))
	// DDUDRLRRUDRD
	fmt.Println("Shortest path example 3:", shortestPath(passcodeExample3))
	// DRURDRUDDLLDLUURRDULRLDUUDDDRR
	fmt.Println("Shortest path puzzle:", shortestPath(puzzlePasscode))
}

type position struct {
	x int
	y int
}

type step struct {
	path     string
	position position
}

func shortestPath(passcode string) string {
	startPosition := position{0, 0}
	nodesToVisit := make([]step, 1, 16)
	nodesToVisit[0] = step{"", startPosition}
	for {
		here := nodesToVisit[0]
		nodesToVisit = nodesToVisit[1:]
		allowedDirections := getAllowedDirections(getMd5Prefix(passcode, here.path))
		for _, direction := range allowedDirections {
			newPosition := determinePosition(here.position, direction)
			if isOutside(newPosition) {
				continue
			}
			newPath := here.path + direction
			if newPosition.x == roomSize && newPosition.y == roomSize {
				return newPath
			}
			nodesToVisit = append(nodesToVisit, step{newPath, newPosition})
		}
	}
}

func isOutside(p position) bool {
	return p.x < 0 || p.x > roomSize || p.y < 0 || p.y > roomSize
}

var directionMap = map[string]func(position) position{
	"U": func(p position) position { return position{p.x, p.y - 1} },
	"D": func(p position) position { return position{p.x, p.y + 1} },
	"L": func(p position) position { return position{p.x - 1, p.y} },
	"R": func(p position) position { return position{p.x + 1, p.y} },
}

func determinePosition(position position, direction string) position {
	return directionMap[direction](position)
}

func getMd5Prefix(passcode string, path string) []byte {
	var md5 = md5.Sum([]byte(passcode + path))
	return md5[:2]
}

var directions = strings.Split("UDLR", "")

func getAllowedDirections(md5Prefix []byte) []string {
	result := make([]string, 0, 4)
	for index, direction := range directions {
		nibble := md5Prefix[index/2] >> (((uint(index) + 1) % 2) * 4) & 15
		if nibble > 0xa {
			result = append(result, direction)
		}
	}
	return result
}
