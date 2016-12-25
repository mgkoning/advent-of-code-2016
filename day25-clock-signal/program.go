package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type state struct {
	registers   map[string]int
	codePointer int
}

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

var showOutput = false

func main() {
	startA := 198

	instructions := getInstructions()
	currentState := state{
		codePointer: 0,
		registers: map[string]int{
			"a": startA, "b": 0, "c": 0, "d": startA + 2532,
		},
	}
	for currentState.codePointer < len(instructions) {
		currentState = execute(instructions[currentState.codePointer], currentState)
	}
	fmt.Println(currentState)

}

func execute(instruction string, state state) state {
	resultState := state
	resultState.registers = duplicate(state.registers)
	parts := strings.Fields(instruction)
	switch parts[0] {
	case "inc":
		state = increment(state, parts[1], 1)
	case "dec":
		state = increment(state, parts[1], -1)
	case "cpy":
		state = copy(state, parts[1], parts[2])
	case "jnz":
		state = jumpNotZero(state, parts[1], parts[2])
	case "out":
		state = output(state, parts[1])
	}
	return state
}

func increment(state state, register string, addition int) state {
	if showOutput {
		fmt.Printf("Adding %v to %v\n", addition, register)
	}
	_, registerExists := state.registers[register]
	if registerExists {
		state.registers[register] += addition
	}
	state.codePointer++
	return state
}

func copy(state state, src string, dst string) state {
	if showOutput {
		fmt.Printf("Copying %v to %v\n", src, dst)
	}
	val := getIntOrRegister(state, src)
	_, registerExists := state.registers[dst]
	if registerExists {
		state.registers[dst] = val
	}
	state.codePointer++
	return state
}

func isNumeric(s string) bool {
	return numRegexp.MatchString(s)
}

func getIntOrRegister(state state, val string) int {
	if isNumeric(val) {
		return mustParseInt(val)
	}
	return state.registers[val]
}

func jumpNotZero(state state, test string, offset string) state {
	val := getIntOrRegister(state, test)
	move := 1
	if val != 0 {
		move = getIntOrRegister(state, offset)
	}
	if showOutput {
		fmt.Printf("Jumping %v because %v = %v\n", move, test, val)
	}
	state.codePointer += move
	return state
}

func output(state state, exp string) state {
	val := getIntOrRegister(state, exp)
	fmt.Printf("Output: %v, registers: a: %5d b: %5d c: %5d d: %5d\n", val, state.registers["a"], state.registers["b"], state.registers["c"], state.registers["d"])
	state.codePointer++
	return state
}

var numRegexp = regexp.MustCompile(`-?\d+`)

func mustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	check(err)
	return int(i)
}

func duplicate(m map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func getInstructions() []string {
	bytes, err := ioutil.ReadFile("code.txt")
	check(err)
	contents := string(bytes)
	return strings.Split(contents, "\n")
}
