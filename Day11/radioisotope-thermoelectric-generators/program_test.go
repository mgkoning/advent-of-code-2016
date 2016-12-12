package main

import "testing"

func TestRemove(t *testing.T) {
	before := []string{"a", "b", "c", "d"}
	result := remove(before, []string{})
	if len(result) != 4 {
		t.Error("Expected remove with empty array to yield same length")
	}
	toRemove := []string{"a", "c"}
	result = remove(before, toRemove)
	if len(result) != 2 {
		t.Error("Expected result to be length 2, was", len(result))
	}
	if result[0] != "b" || result[1] != "d" {
		t.Error("Expected 'b, d', got", result)
	}
}

func TestDoMove(t *testing.T) {
	var result = doMove(move{
		fromFloor:  0,
		toFloor:    1,
		microchips: []string{"strontium"},
		generators: []string{},
	}, getStartState())
	if !contains(result[1].microchips, "strontium") {
		t.Error("Strontium not on floor 1")
	}
	if contains(result[0].microchips, "strontium") {
		t.Error("Strontium still on floor 0")
	}
}

func TestDone(t *testing.T) {
	floors := []floor{
		floor{}, floor{}, floor{},
		floor{microchips: []string{"a"}, generators: []string{"a"}},
	}
	if !done(floors) {
		t.Error("Floors should be done")
	}
}

func TestIsBadState(t *testing.T) {
	floors := []floor{
		floor{}, floor{}, floor{},
		floor{microchips: []string{"H", "L"}, generators: []string{"H"}},
	}
	if !isBadState(floors) {
		t.Error("bad state not detected")
	}
}
