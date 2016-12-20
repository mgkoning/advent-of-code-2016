package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rulesChan := make(chan rule)
	go parseInput(rulesChan)
	rules := make([]rule, 0)
	for rule := range rulesChan {
		rules = append(rules, rule)
	}
	sort.Sort(byRange(rules))
	normalizedRules := []rule{rules[0]}
	for _, r := range rules[1:] {
		normalizedRules = normalize(normalizedRules, r)
	}
	fmt.Println("Answer part 1:", normalizedRules[0].to+1)
	allowedIps := normalizedRules[0].from - 0
	for i := 0; i < len(normalizedRules)-1; i++ {
		allowedIps += normalizedRules[i+1].from - 1 - normalizedRules[i].to
	}
	allowedIps += math.MaxUint32 - normalizedRules[len(normalizedRules)-1].to
	fmt.Println("Answer part 2:", allowedIps)
}

func max(a uint, b uint) uint {
	if a < b {
		return b
	}
	return a
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(rules chan rule) {
	b, err := ioutil.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(b), "\r\n")
	for _, line := range lines {
		bounds := strings.Split(line, "-")
		rules <- rule{mustParseUint(bounds[0]), mustParseUint(bounds[1])}
	}
	close(rules)
}

func mustParseUint(s string) uint {
	r, err := strconv.ParseUint(s, 10, 32)
	check(err)
	return uint(r)
}
