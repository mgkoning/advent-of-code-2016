package main

import "fmt"

type floor struct {
	microchips []string
	generators []string
}

type move struct {
	fromFloor  int
	toFloor    int
	microchips []string
	generators []string
}

type sequence struct {
	resultState []floor
	moves       []move
	states      [][]floor
}

func main() {
	startState := getStartStateExample()
	possibleMoves := determinePossibleMoves(startState, 0, 1)
	sequences := make([]sequence, 0)
	for _, possibleMove := range possibleMoves {
		resultState := doMove(possibleMove, startState)
		if isBadState(resultState) {
			continue
		}
		fmt.Println("Adding move ", possibleMove)
		sequences = append(sequences, sequence{resultState: resultState, moves: []move{possibleMove}, states: [][]floor{resultState}})
	}
	var winner sequence
	advancements := 0
	for {
		fmt.Println("Sequences to consider: ", len(sequences))
		seq, found := findWinner(sequences)
		if found {
			winner = seq
			break
		}
		sequences = advance(sequences)
		advancements++
		fmt.Println("Advancements: ", advancements)
		if len(sequences) == 0 {
			panic("No more moves to consider!")
		}
	}
	fmt.Println("Winning sequence: ", winner.moves)
	state := startState
	for i, move := range winner.moves {
		state = doMove(move, state)
		fmt.Println("State after move ", i, ": ", state)
	}
	fmt.Println(winner.resultState)
	fmt.Println("# moves: ", len(winner.moves))
}

func findWinner(sequences []sequence) (sequence, bool) {
	for _, seq := range sequences {
		if done(seq.resultState) {
			return seq, true
		}
	}
	return sequence{}, false
}

func advance(sequences []sequence) []sequence {
	newSituation := make([]sequence, 0)
	for _, seq := range sequences {
		theMove := seq.moves[len(seq.moves)-1]

		currentFloor := theMove.toFloor
		for _, nextFloor := range []int{currentFloor + 1, currentFloor - 1} {
			if nextFloor < 0 || 3 < nextFloor {
				continue
			}
			possibleMoves := determinePossibleMoves(seq.resultState, currentFloor, nextFloor)
			for _, possibleMove := range possibleMoves {
				resultState := doMove(possibleMove, seq.resultState)
				if isBadState(resultState) {
					continue
				}
				if seenBefore(seq.states, resultState) {
					continue
				}
				newMoves := make([]move, len(seq.moves), len(seq.moves)+1)
				copy(newMoves, seq.moves)
				newMoves = append(newMoves, possibleMove)
				newStates := make([][]floor, len(seq.states), len(seq.states)+1)
				copy(newStates, seq.states)
				newStates = append(newStates, resultState)
				newSequence := sequence{resultState: resultState, moves: newMoves, states: newStates}
				newSituation = append(newSituation, newSequence)
			}
		}
	}
	return newSituation
}

func getStartState() []floor {
	return []floor{
		floor{microchips: []string{"strontium", "plutonium"}, generators: []string{"strontium", "plutonium"}},
		floor{microchips: []string{"ruthenium", "curium"}, generators: []string{"thulium", "ruthenium", "curium"}},
		floor{microchips: []string{"thulium"}},
		floor{},
	}
}

func getStartStateEasy() []floor {
	return []floor{
		floor{microchips: []string{"strontium"}, generators: []string{"strontium"}},
		floor{},
		floor{},
		floor{},
	}
}

func getStartStateExample() []floor {
	return []floor{
		floor{microchips: []string{"hydrogen", "lithium"}},
		floor{generators: []string{"hydrogen"}},
		floor{generators: []string{"lithium"}},
		floor{},
	}
}

func (move move) String() string {
	return fmt.Sprint(
		"Move microchips ", move.microchips, " and generators ", move.generators,
		" from floor ", move.fromFloor, " to floor ", move.toFloor,
	)
}

func determinePossibleMoves(state []floor, fromFloor int, toFloor int) []move {
	moves := make([]move, 0)
	currentFloor := state[fromFloor]
	if len(currentFloor.microchips) == 0 && len(currentFloor.generators) == 0 {
		return moves
	}
	for _, microchip := range currentFloor.microchips {
		candidate := move{fromFloor: fromFloor, toFloor: toFloor, microchips: []string{microchip}}
		moves = append(moves, candidate)
		for _, microchip2 := range currentFloor.microchips {
			if microchip == microchip2 {
				continue
			}
			candidate := move{fromFloor: fromFloor, toFloor: toFloor, microchips: []string{microchip, microchip2}}
			moves = append(moves, candidate)
		}
	}
	for _, generator := range currentFloor.generators {
		candidate := move{fromFloor: fromFloor, toFloor: toFloor, generators: []string{generator}}
		moves = append(moves, candidate)
		for _, generator2 := range currentFloor.generators {
			if generator == generator2 {
				continue
			}
			candidate := move{fromFloor: fromFloor, toFloor: toFloor, generators: []string{generator, generator2}}
			moves = append(moves, candidate)
		}
	}
	for _, microchip := range currentFloor.microchips {
		for _, generator := range currentFloor.generators {
			microchips := []string{microchip}
			generators := []string{generator}
			if areAnyChipsFried(microchips, generators) {
				continue
			}
			candidate := move{fromFloor: fromFloor, toFloor: toFloor, microchips: microchips, generators: generators}
			moves = append(moves, candidate)
		}
	}
	return moves
}

func doMove(move move, startState []floor) []floor {
	newState := copyState(startState)
	newState[move.toFloor].generators = append(newState[move.toFloor].generators, move.generators...)
	newState[move.toFloor].microchips = append(newState[move.toFloor].microchips, move.microchips...)
	newState[move.fromFloor].generators = remove(newState[move.fromFloor].generators, move.generators)
	newState[move.fromFloor].microchips = remove(newState[move.fromFloor].microchips, move.microchips)
	return newState
}

func copyState(state []floor) []floor {
	resultState := make([]floor, 4)
	copy(resultState, state)
	return resultState
}

func remove(before []string, toRemove []string) []string {
	if len(toRemove) == 0 {
		return before
	}
	result := make([]string, 0)
	for _, s := range before {
		if contains(toRemove, s) {
			continue
		}
		result = append(result, s)
	}
	return result
}

func seenBefore(previous [][]floor, new []floor) bool {
	for _, prv := range previous {
		if isSameState(prv, new) {
			return true
		}
	}
	return false
}

func isSameState(previous []floor, new []floor) bool {
	for i := range previous {
		if !areSame(previous[i].generators, new[i].generators) {
			return false
		}
		if !areSame(previous[i].microchips, new[i].microchips) {
			return false
		}
	}
	return true
}

func done(state []floor) bool {
	if len(state) != 4 {
		panic("state invalid, must have 4 floors")
	}
	return state[0].isEmpty() && state[1].isEmpty() && state[2].isEmpty() &&
		len(state[3].microchips) == len(state[3].generators)
}

func (floor floor) isEmpty() bool {
	return len(floor.microchips) == 0 && len(floor.generators) == 0
}

func isBadState(floors []floor) bool {
	for _, floor := range floors {
		if areAnyChipsFried(floor.microchips, floor.generators) {
			return true
		}
	}
	return false
}

func areAnyChipsFried(microchips []string, generators []string) bool {
	if len(generators) == 0 {
		return false
	}
	var exposedChips []string
	for _, element := range microchips {
		if contains(generators, element) {
			continue
		}
		exposedChips = append(exposedChips, element)
	}
	return 0 < len(exposedChips)
}

func areSame(one []string, other []string) bool {
	if len(one) != len(other) {
		return false
	}
	for _, element := range one {
		if !contains(other, element) {
			return false
		}
	}
	return true
}

func contains(elements []string, element string) bool {
	for _, e := range elements {
		if e == element {
			return true
		}
	}
	return false
}
