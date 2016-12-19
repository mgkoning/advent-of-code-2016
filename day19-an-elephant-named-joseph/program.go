package main

import "fmt"

var exampleInput = 5
var puzzleInput = 3014603

func main() {
	determineAnswer(exampleInput)
	determineAnswer(puzzleInput)
}

func determineAnswer(elvesCount int) {
	elves := make([]bool, elvesCount)
	previous := 0
	next := 0
	eliminateNext := true
	for {
		next = scanToNext(elves, next)
		if next == previous {
			break
		}
		if eliminateNext {
			elves[next] = true
		} else {
			previous = next
		}
		eliminateNext = !eliminateNext
	}
	fmt.Println("Elf", next+1, "has all presents.")
}

func scanToNext(elves []bool, current int) int {
	next := current
	for {
		next = (next + 1) % len(elves)
		if !elves[next] {
			return next
		}
	}
}
