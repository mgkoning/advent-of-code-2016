package main

import "testing"

type testCase struct {
	input          int
	expectedOutput int
}

var testCases = []testCase{
	{1, 1},
	{2, 1},
	{3, 2},
	{0x11001101001, 6},
	{249732, 10},
}

func TestNumberOf1Bits(t *testing.T) {
	for _, testCase := range testCases {
		var output = numberOf1Bits(testCase.input)
		if output != testCase.expectedOutput {
			t.Error("Expected", testCase.expectedOutput, "but was", output)
		}
	}
}
