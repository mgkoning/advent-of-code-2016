package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
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

	// part 2
	fmt.Println("Longest path example 1:", getLongestPath(allPaths(passcodeExample1)))
	// 370
	fmt.Println("Longest path example 2:", getLongestPath(allPaths(passcodeExample2)))
	// 492
	fmt.Println("Longest path example 3:", getLongestPath(allPaths(passcodeExample3)))
	// 830
	fmt.Println("Longest path puzzle:", getLongestPath(allPaths(puzzlePasscode)))
}

/* results:
Shortest path example 1: DDRRRD
Shortest path example 2: DDUDRLRRUDRD
Shortest path example 3: DRURDRUDDLLDLUURRDULRLDUUDDDRR
Shortest path puzzle: RLRDRDUDDR
Time elapsed (ms): 59.0418
Longest path example 1: 370
Time elapsed (ms): 53.5446
Longest path example 2: 492
Time elapsed (ms): 62.0445
Longest path example 3: 830
Time elapsed (ms): 37.0261
Longest path puzzle: 420
*/

// assume sorted
func getLongestPath(paths []string) int {
	return len(paths[len(paths)-1])
}

type position struct {
	x int
	y int
}

type step struct {
	path     string
	position position
}

func timeSince(from time.Time) {
	fmt.Println("Time elapsed (ms):", float64(time.Now().Sub(from).Nanoseconds())/float64(1000*1000))
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

func allPaths(passcode string) []string {
	defer timeSince(time.Now())
	startPosition := position{0, 0}
	nodesToVisit := make([]step, 1, 16)
	nodesToVisit[0] = step{"", startPosition}
	foundPaths := make([]string, 0)
	for len(nodesToVisit) > 0 {
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
				foundPaths = append(foundPaths, newPath)
				continue
			}
			nodesToVisit = append(nodesToVisit, step{newPath, newPosition})
		}
	}
	return foundPaths
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

/* Save some time by not converting to string, but just returning first two bytes. */
func getMd5Prefix(passcode string, path string) []byte {
	var md5 = md5.Sum([]byte(passcode + path))
	return md5[:2]
}

var directions = strings.Split("UDLR", "")

/* Use the nibbles in md5Prefix to determine allowed directions. */
/* Index 0 needs byte 0 rightshifted by 4, Index 1 needs byte 0 bitwise ANDed with 0xf, etc. */
func getAllowedDirections(md5Prefix []byte) []string {
	result := make([]string, 0, 4)
	for index, direction := range directions {
		nibble := md5Prefix[index/2] >> (((uint(index) + 1) % 2) * 4) & 0xf
		if nibble < 0xb {
			continue
		}
		result = append(result, direction)
	}
	return result
}
