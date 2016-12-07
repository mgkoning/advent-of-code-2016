package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func main() {
	contents := readInput()
	addresses := strings.Split(contents, "\n")
	var supportingTLS []string
	for _, address := range addresses {
		if !supportsTLS(address) {
			continue
		}
		supportingTLS = append(supportingTLS, address)
		fmt.Printf("Supports TLS: %v\n", address)
	}
	fmt.Printf("Addresses supporting TLS: %v\n", len(supportingTLS))
}

func readInput() string {
	data, err := ioutil.ReadFile("day7.txt")
	check(err)
	return string(data)
}

func supportsTLS(address string) bool {
	sequences, hypernetSequences := parse(address)
	for _, sequence := range hypernetSequences {
		if hasABBA(sequence) {
			return false
		}
	}
	for _, sequence := range sequences {
		if hasABBA(sequence) {
			return true
		}
	}
	return false
}

func parse(address string) (sequences [][]string, hypernetSequences [][]string) {
	exploded := strings.Split(address, "")
	var sequence []string
	for _, char := range exploded {
		if char == "[" {
			sequences = joinAndAppendIfNotEmpty(sequences, sequence)
			sequence = make([]string, 0)
			continue
		}
		if char == "]" {
			hypernetSequences = joinAndAppendIfNotEmpty(hypernetSequences, sequence)
			sequence = make([]string, 0)
			continue
		}
		sequence = append(sequence, char)
	}
	// assuming there's no bad input, a non-empty sequence
	// means that the last char wasn't [ or ]
	sequences = joinAndAppendIfNotEmpty(sequences, sequence)
	return sequences, hypernetSequences
}

func joinAndAppendIfNotEmpty(sequences [][]string, sequence []string) [][]string {
	if len(sequence) == 0 {
		return sequences
	}
	return append(sequences, sequence)
}

func hasABBA(sequence []string) bool {
	for index := 0; index < len(sequence)-3; index++ {
		if sequence[index] != sequence[index+1] &&
			sequence[index] == sequence[index+3] &&
			sequence[index+1] == sequence[index+2] {
			return true
		}
	}
	return false
}
