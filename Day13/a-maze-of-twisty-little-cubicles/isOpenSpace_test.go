package main

import "testing"

type isOpenSpaceTestCase struct {
	input          coordinate
	seed           int
	expectedOutput bool
}

var isOpenSpaceTestCases = []isOpenSpaceTestCase{
	{coordinate{0, 0}, 10, true},
	{coordinate{1, 1}, 10, true},
	{coordinate{4, 3}, 10, false},
	{coordinate{9, 6}, 10, false},
	{coordinate{9, 5}, 10, true},
	{coordinate{31, 39}, designersFavoriteNumber, true},
}

func TestIsOpenSpace(t *testing.T) {
	for _, testCase := range isOpenSpaceTestCases {
		var output = testCase.input.isOpenSpace(testCase.seed)
		if output != testCase.expectedOutput {
			t.Error(testCase.input, "expected to be", testCase.expectedOutput, "but was", output)
		}
	}
}
