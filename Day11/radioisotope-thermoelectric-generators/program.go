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

type state struct {
	currentFloor  int
	numberOfMoves int
	floors        []floor
}

func (state state) String() string {
	result := fmt.Sprintf("E%v", state.currentFloor)
	for i, floor := range state.floors {
		result += fmt.Sprintf("F%vG%vM%v", i, floor.generators, floor.microchips)
	}
	return result
}

func main() {
	var answer = determineAnswer()
	fmt.Println(answer.floors)
	fmt.Println("# moves: ", answer.numberOfMoves)
}

func determineAnswer() state {
	var seenStates = make(map[string]state)
	var statesToConsider = make([]state, 0)

	startFloors := getStartState()
	totalObjects := countObjects(startFloors)
	initialState := state{currentFloor: 0, numberOfMoves: 0, floors: startFloors}
	seenStates[initialState.String()] = initialState
	fmt.Println(initialState.String())
	statesToConsider = append(statesToConsider, initialState)

	for iterations := 0; ; iterations++ {
		if len(statesToConsider) == 0 {
			break
		}
		curState := statesToConsider[0]
		statesToConsider = statesToConsider[1:]

		for _, nextFloor := range []int{curState.currentFloor - 1, curState.currentFloor + 1} {
			if nextFloor < 0 || 3 < nextFloor {
				continue
			}
			possibleMoves := determinePossibleMoves(curState.floors, curState.currentFloor, nextFloor)
			for _, possibleMove := range possibleMoves {
				resultFloors := doMove(possibleMove, curState.floors)
				objectCount := countObjects(resultFloors)
				if objectCount != totalObjects {
					fmt.Println("offending move", possibleMove)
					fmt.Println("before", curState.floors)
					fmt.Println("after", resultFloors)
					panic(fmt.Sprint("Mismatched count of ", objectCount, " rather than ", totalObjects))
				}
				nextState := state{currentFloor: possibleMove.toFloor, numberOfMoves: curState.numberOfMoves + 1, floors: resultFloors}
				if isBadState(resultFloors) {
					continue
				}
				var stateName = nextState.String()
				var _, seen = seenStates[stateName]
				if seen {
					continue
				}
				seenStates[stateName] = nextState
				// if seenBefore(visitedStates, nextState) {
				// 	continue
				// }
				if done(nextState.floors) {
					fmt.Println("Found winner after #moves: ", nextState.numberOfMoves)
					return nextState
				}
				statesToConsider = append(statesToConsider, nextState)
			}
		}
	}
	panic("no winner found")
}

func countObjects(floors []floor) int {
	count := 0
	for _, floor := range floors {
		count += len(floor.generators) + len(floor.microchips)
	}
	return count
}

func getStartState() []floor {
	return []floor{
		floor{microchips: []string{"plutonium", "strontium"}, generators: []string{"plutonium", "strontium"}},
		floor{microchips: []string{"curium", "ruthenium"}, generators: []string{"curium", "ruthenium", "thulium"}},
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
		for _, microchip2 := range currentFloor.microchips {
			if microchip == microchip2 {
				continue
			}
			candidate := move{fromFloor: fromFloor, toFloor: toFloor, microchips: []string{microchip, microchip2}}
			moves = append(moves, candidate)
		}
		candidate := move{fromFloor: fromFloor, toFloor: toFloor, microchips: []string{microchip}}
		moves = append(moves, candidate)
	}
	for _, generator := range currentFloor.generators {
		for _, generator2 := range currentFloor.generators {
			if generator == generator2 {
				continue
			}
			candidate := move{fromFloor: fromFloor, toFloor: toFloor, generators: []string{generator, generator2}}
			moves = append(moves, candidate)
		}
		candidate := move{fromFloor: fromFloor, toFloor: toFloor, generators: []string{generator}}
		moves = append(moves, candidate)
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
	//sort.Strings(newState[move.toFloor].generators)
	newState[move.toFloor].microchips = append(newState[move.toFloor].microchips, move.microchips...)
	//sort.Strings(newState[move.toFloor].microchips)
	newState[move.fromFloor].generators = remove(newState[move.fromFloor].generators, move.generators)
	newState[move.fromFloor].microchips = remove(newState[move.fromFloor].microchips, move.microchips)
	return newState
}

func copyState(state []floor) []floor {
	resultState := make([]floor, 4)
	copy(resultState, state)
	for i := range resultState {
		resultMicrochips := make([]string, len(state[i].microchips))
		copy(resultMicrochips, state[i].microchips)
		resultState[i].microchips = resultMicrochips
		resultGenerators := make([]string, len(state[i].generators))
		copy(resultGenerators, state[i].generators)
		resultState[i].generators = resultGenerators
	}
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

func seenBefore(previous []state, new state) bool {
	for _, prv := range previous {
		if isSameState(prv, new) {
			return true
		}
	}
	return false
}

func isSameState(previous state, new state) bool {
	if previous.currentFloor != new.currentFloor {
		return false
	}
	for i := range previous.floors {
		if !areSame(previous.floors[i].generators, new.floors[i].generators) {
			return false
		}
		if !areSame(previous.floors[i].microchips, new.floors[i].microchips) {
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
