package main

import "fmt"

var exampleInput = 5
var puzzleInput = 3014603

func main() {
	determineAnswer(exampleInput)
	determineAnswer(puzzleInput)
	determineAnswerPart2(exampleInput)
	determineAnswerPart2(puzzleInput)
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

type elf struct {
	num  int
	prev *elf
	next *elf
}

func determineAnswerPart2(elvesCount int) {
	var startElf = &elf{num: 1}
	var previousElf = startElf
	for i := 2; i <= elvesCount; i++ {
		var newElf = elf{num: i, prev: previousElf}
		previousElf.next = &newElf
		previousElf = &newElf
	}
	startElf.prev = previousElf
	startElf.prev.next = startElf
	remaining := elvesCount
	currentElf := startElf
	for 1 < remaining {
		eliminee := scan(currentElf, remaining/2)
		eliminate(eliminee)
		remaining--
		currentElf = currentElf.next
		if (elvesCount-remaining)%500 == 0 {
			fmt.Println("Remaining:", remaining)
		}
	}
	fmt.Println("Elf", currentElf.num, "stole all the presents.")
}

func scan(from *elf, num int) *elf {
	result := from
	for 0 < num {
		result = result.next
		num--
	}
	return result
}

func eliminate(eliminee *elf) {
	next := eliminee.next
	prev := eliminee.prev
	next.prev = prev
	prev.next = next
}
