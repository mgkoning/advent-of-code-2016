package main

import "testing"

type decompressTestCase struct {
	input  string
	output string
}

var testCases = []decompressTestCase{
	{input: "ADVENT", output: "ADVENT"},
	{input: "A(1x5)BC", output: "ABBBBBC"},
	{input: "(3x3)XYZ", output: "XYZXYZXYZ"},
	{input: "A(2x2)BCD(2x2)EFG", output: "ABCBCDEFEFG"},
	{input: "(6x1)(1x3)A", output: "(1x3)A"},
	{input: "X(8x2)(3x3)ABCY", output: "X(3x3)ABC(3x3)ABCY"},
}

func TestDecompress(t *testing.T) {
	for _, testCase := range testCases {
		var actualOutput = decompress(testCase.input)
		if testCase.output != actualOutput {
			t.Error("Expected", testCase.output, "got", actualOutput)
		}
	}

}
