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

var showOutput = true

func main() {
	instructions := getInstructions()
	startA := 7
	currentState := state{
		codePointer: 0,
		registers: map[string]int{
			"a": startA, "b": 0, "c": 0, "d": 0,
		},
	}
	for currentState.codePointer < len(instructions) {
		currentState = execute(instructions[currentState.codePointer], currentState, instructions)
		fmt.Println(currentState)
	}
}

func execute(instruction string, state state, instructions []string) state {
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
	case "tgl":
		state = toggle(state, parts[1], instructions)
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

func toggle(state state, offset string, instructions []string) state {
	instructionIndex := state.codePointer + getIntOrRegister(state, offset)
	if showOutput {
		fmt.Printf("Toggling instruction at index '%v': ", instructionIndex)
	}
	if 0 <= instructionIndex && instructionIndex < len(instructions) {
		instructions[instructionIndex] = toggleInstruction(instructions[instructionIndex])
	}
	state.codePointer++
	return state
}

var instructionRegexp = regexp.MustCompile(`^(dec|inc|cpy|jnz|tgl)`)

func toggleInstruction(instruction string) string {
	return instructionRegexp.ReplaceAllStringFunc(instruction, func(op string) string {
		var result string
		switch op {
		case "dec":
		case "tgl":
			result = "inc"
		case "inc":
			result = "dec"
		case "jnz":
			result = "cpy"
		case "cpy":
			result = "jnz"
		default:
			panic(fmt.Sprintf("Don't understand '%v'", op))
		}
		if showOutput {
			fmt.Printf("Changing op '%v' to '%v'\n", op, result)
		}
		return result
	})
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
