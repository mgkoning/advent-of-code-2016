package main

import "testing"

var exampleDiscs = []disc{
	{1, 5, 4},
	{2, 2, 1},
}

func TestCanBePassed(t *testing.T) {
	assertCanBePassed(exampleDiscs[0], 0, true, t)
	assertCanBePassed(exampleDiscs[1], 0, false, t)
	assertCanBePassed(exampleDiscs[0], 5, true, t)
	assertCanBePassed(exampleDiscs[1], 5, true, t)
}

func assertCanBePassed(disc disc, startTime int, expected bool, t *testing.T) {
	if canBePassed(disc, startTime) != expected {
		t.Error("At time", startTime, "disc", disc, "expected", expected)
	}
}
