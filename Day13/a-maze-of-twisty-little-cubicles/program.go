package main

import "fmt"

var designersFavoriteNumber = 1362

type coordinate struct {
	X int
	Y int
}

func main() {
	start := coordinate{1, 1}
	goal := coordinate{31, 39}
	fmt.Println("goal reached after steps:", minDistance(start, goal))
}

func minDistance(start coordinate, goal coordinate) int {
	seen := make(map[coordinate]int)
	toVisit := make([]coordinate, 0)
	toVisit = append(toVisit, start)
	seen[start] = 0
	for 0 < len(toVisit) {
		next := toVisit[0]
		toVisit = toVisit[1:]
		neighborDistance := seen[next] + 1
		for _, neighbor := range next.getNeighbors() {
			if neighbor.X < 0 || neighbor.Y < 0 {
				continue
			}
			if !neighbor.isOpenSpace(designersFavoriteNumber) {
				continue
			}
			_, neighborSeen := seen[neighbor]
			if neighborSeen {
				continue
			}
			if neighbor == goal {
				return neighborDistance
			}
			seen[neighbor] = neighborDistance
			toVisit = append(toVisit, neighbor)
		}
	}
	panic(fmt.Sprintf("Goal %v not found", goal))
}

func (c coordinate) getNeighbors() []coordinate {
	return []coordinate{
		{c.X, c.Y - 1},
		{c.X, c.Y + 1},
		{c.X - 1, c.Y},
		{c.X + 1, c.Y},
	}
}

func (c coordinate) isOpenSpace(seed int) bool {
	return numberOf1Bits(c.X*c.X+3*c.X+2*c.X*c.Y+c.Y+c.Y*c.Y+seed)%2 == 0
}

func numberOf1Bits(number int) int {
	sum := 0
	for 0 < number {
		sum += number & 0x1
		number >>= 1
	}
	return sum
}
