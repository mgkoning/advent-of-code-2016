package main

import (
	"testing"
)

func TestParseBot(t *testing.T) {
	bot := parseBot("bot 2 gives low to bot 1 and high to bot 0")
	if bot.name != "bot 2" {
		t.Error("bot name should be bot 2 but was", bot.name)
	}
	if bot.low != "bot 1" {
		t.Error("bot low should be bot 1 but was", bot.low)
	}
	if bot.high != "bot 0" {
		t.Error("bot high should be bot 0 but was", bot.high)
	}
}

func TestParseValue(t *testing.T) {
	value := parseValue("value 5 goes to bot 2")
	if value.value != 5 {
		t.Error("value should be 5 but was", value.value)
	}
	if value.bot != "bot 2" {
		t.Error("bot should be bot 2 but was", value.bot)
	}
}
