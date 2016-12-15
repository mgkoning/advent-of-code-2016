package main

import "fmt"

type disc struct {
	distance          int
	numberOfPositions int
	positionAtStart   int
}

var discs = []disc{
	{1, 13, 11},
	{2, 5, 0},
	{3, 17, 11},
	{4, 3, 0},
	{5, 7, 2},
	{6, 19, 17},
}

func main() {
	fmt.Println("Part 1:")
	findTime(discs)
	fmt.Println("Part 2:")
	findTime(append(discs, disc{7, 11, 0}))
}

func findTime(discs []disc) {
	var time int
	for time = 0; ; time++ {
		if all(discs, func(disc disc) bool { return canBePassed(disc, time) }) {
			break
		}
	}
	fmt.Println("First opportune time is", time)
}

func all(discs []disc, predicate func(disc) bool) bool {
	for _, disc := range discs {
		if predicate(disc) {
			continue
		}
		return false
	}
	return true
}

func canBePassed(disc disc, startTime int) bool {
	return (startTime+disc.distance+disc.positionAtStart)%disc.numberOfPositions == 0
}
