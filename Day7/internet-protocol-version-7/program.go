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
	var supportingSSL []string
	for _, address := range addresses {
		sequences, hypernetSequences := parse(address)
		if supportsTLS(sequences, hypernetSequences) {
			supportingTLS = append(supportingTLS, address)
		}
		if supportSSL(sequences, hypernetSequences) {
			supportingSSL = append(supportingSSL, address)
		}
	}
	fmt.Printf("Addresses supporting TLS: %v\n", len(supportingTLS))
	fmt.Printf("Addresses supporting SSL: %v\n", len(supportingSSL))
}

func readInput() string {
	data, err := ioutil.ReadFile("day7.txt")
	check(err)
	return string(data)
}

func supportsTLS(sequences [][]string, hypernetSequences [][]string) bool {
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

func supportSSL(sequences [][]string, hypernetSequences [][]string) bool {
	ABAs := getABAs(sequences)
	for _, ABA := range ABAs {
		for _, sequence := range hypernetSequences {
			if hasBAB(sequence, ABA) {
				return true
			}
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

func getABAs(sequences [][]string) [][]string {
	var ABAs [][]string
	for _, sequence := range sequences {
		for index := 0; index < len(sequence)-2; index++ {
			if sequence[index] != sequence[index+1] &&
				sequence[index] == sequence[index+2] {
				foundABA := sequence[index : index+3]
				ABAs = append(ABAs, foundABA)
			}
		}
	}
	return ABAs
}

func hasBAB(hypernetsequence []string, ABA []string) bool {
	BAB := strings.Join([]string{ABA[1], ABA[0], ABA[1]}, "")
	return strings.Contains(strings.Join(hypernetsequence, ""), BAB)
}
