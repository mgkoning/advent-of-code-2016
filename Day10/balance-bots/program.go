package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

type value struct {
	value int
	bot   string
}

type bot struct {
	name string
	low  string
	high string
}

func main() {
	specifications := getInput()
	bots, values := parseSpecifications(specifications)

	valuesPerBot := make(map[string][]int)
	outputValues := make(map[string]int)
	for _, value := range values {
		valuesPerBot[value.bot] = append(valuesPerBot[value.bot], value.value)
	}

	runProcess(bots, valuesPerBot, outputValues)

	answerPart1 := answerPart1(valuesPerBot)
	answerPart2 := answerPart2(outputValues)
	fmt.Println(answerPart1, "compares 17 and 61")
	fmt.Println("Outputs 0, 1 and 2 multiplied makes", answerPart2)
}

func parseSpecifications(specifications []string) (bots map[string]bot, values []value) {
	bots = make(map[string]bot)
	values = make([]value, 0)
	for _, line := range specifications {
		if isBot(line) {
			bot := parseBot(line)
			bots[bot.name] = bot
		} else {
			values = append(values, parseValue(line))
		}
	}
	return bots, values
}

func answerPart1(valuesPerBot map[string][]int) string {
	var answerPart1 = ""
	for name, values := range valuesPerBot {
		sort.Ints(values)
		if values[0] == 17 && values[1] == 61 {
			answerPart1 = name
		}
		fmt.Println(name, "compares", values[0], "and", values[1])
	}
	fmt.Println()
	return answerPart1
}

func answerPart2(outputValues map[string]int) int {
	var answerPart2 = 1
	for name, value := range outputValues {
		if name == "output 0" || name == "output 1" || name == "output 2" {
			answerPart2 *= value
		}
		fmt.Println(name, "holds value", value)
	}
	fmt.Println()
	return answerPart2
}

func runProcess(bots map[string]bot, valuesPerBot map[string][]int, outputValues map[string]int) {
	done := make(map[string]bool)
	botName, botValues := getNext(valuesPerBot, done)
	for {
		if botName == "" {
			break
		}
		process(botName, botValues, bots, valuesPerBot, outputValues)
		done[botName] = true
		botName, botValues = getNext(valuesPerBot, done)
	}

}

func getNext(valuesPerBot map[string][]int, done map[string]bool) (string, []int) {
	for k, v := range valuesPerBot {
		if len(v) < 2 {
			continue
		}
		if done[k] {
			continue
		}
		return k, v
	}
	return "", nil
}

func process(botName string, values []int, bots map[string]bot, valuesPerBot map[string][]int, outputValues map[string]int) {
	bot := bots[botName]
	sort.Ints(values)
	doOutput(bot.low, values[0], valuesPerBot, outputValues)
	doOutput(bot.high, values[1], valuesPerBot, outputValues)
}

func doOutput(name string, value int, valuesPerBot map[string][]int, outputValues map[string]int) {
	if isBot(name) {
		valuesPerBot[name] = append(valuesPerBot[name], value)
	} else {
		outputValues[name] = value
	}
}

func isBot(s string) bool {
	return strings.HasPrefix(s, "bot ")
}

func getInput() []string {
	bytes, err := ioutil.ReadFile("input.txt")
	check(err)
	return strings.Split(string(bytes), "\n")
}

var botRegex = regexp.MustCompile(`(bot \d+) gives low to ((bot|output) \d+) and high to ((bot|output) \d+)`)

func parseBot(line string) bot {
	matches := botRegex.FindStringSubmatch(line)
	if matches == nil {
		panic(fmt.Sprintf("Do not understand bot '%v'", line))
	}
	return bot{name: matches[1], low: matches[2], high: matches[4]}
}

var valueRegex = regexp.MustCompile(`value (\d+) goes to (bot \d+)`)

func parseValue(line string) value {
	matches := valueRegex.FindStringSubmatch(line)
	if matches == nil {
		panic(fmt.Sprintf("Do not understand value '%v'", line))
	}
	return value{value: mustParseInt(matches[1]), bot: matches[2]}
}

func mustParseInt(s string) int {
	result, err := strconv.ParseInt(s, 10, 32)
	check(err)
	return int(result)
}
