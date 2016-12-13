package main

import "fmt"

type floor struct {
	microchips []string
	generators []string
}

func main() {
	var answer = determineAnswer(getStartState())
	fmt.Println("Steps part 1: ", answer)
	var answer2 = determineAnswer(getStartStatePart2())
	fmt.Println("Steps part 2: ", answer2)
}

func determineAnswer(puzzle [4]floor) int {
	steps := 0
	sumItems := 0
	for i := 0; i < 3; i++ {
		floor := puzzle[i]
		sumItems += len(floor.microchips) + len(floor.generators)
		if sumItems < 3 {
			steps++
		} else {
			steps += sumItems + sumItems - 3
		}
	}
	return steps
}

func getStartState() [4]floor {
	return [4]floor{
		floor{microchips: []string{"plutonium", "strontium"}, generators: []string{"plutonium", "strontium"}},
		floor{microchips: []string{"curium", "ruthenium"}, generators: []string{"curium", "ruthenium", "thulium"}},
		floor{microchips: []string{"thulium"}},
		floor{},
	}
}

func getStartStatePart2() [4]floor {
	return [4]floor{
		floor{microchips: []string{"plutonium", "strontium", "elerium", "dilithium"}, generators: []string{"plutonium", "strontium", "elerium", "dilithium"}},
		floor{microchips: []string{"curium", "ruthenium"}, generators: []string{"curium", "ruthenium", "thulium"}},
		floor{microchips: []string{"thulium"}},
		floor{},
	}
}
