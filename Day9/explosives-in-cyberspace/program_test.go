package main

import "testing"

var decompressTestCases = []decompressedLengthTestCase{
	{input: "ADVENT", output: int64(len("ADVENT"))},
	{input: "A(1x5)BC", output: int64(len("ABBBBBC"))},
	{input: "(3x3)XYZ", output: int64(len("XYZXYZXYZ"))},
	{input: "A(2x2)BCD(2x2)EFG", output: int64(len("ABCBCDEFEFG"))},
	{input: "(6x1)(1x3)A", output: int64(len("(1x3)A"))},
	{input: "X(8x2)(3x3)ABCY", output: int64(len("X(3x3)ABC(3x3)ABCY"))},
}

func TestDecompress(t *testing.T) {
	for _, testCase := range decompressTestCases {
		var actualOutput = decompressedLength(testCase.input, false)
		if testCase.output != actualOutput {
			t.Error("Expected", testCase.output, "got", actualOutput)
		}
	}
}

type decompressedLengthTestCase struct {
	input  string
	output int64
}

var decompressedLengthVersion2TestCases = []decompressedLengthTestCase{
	{input: "ADVENT", output: int64(len("ADVENT"))},
	{input: "(3x3)XYZ", output: int64(len("XYZXYZXYZ"))},
	{input: "X(8x2)(3x3)ABCY", output: int64(len("XABCABCABCABCABCABCY"))},
	{input: "(27x12)(20x12)(13x14)(7x10)(1x12)A", output: 241920},
	{input: "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", output: 445},
}

func TestDecompressedLengthVersion2(t *testing.T) {
	for _, testCase := range decompressedLengthVersion2TestCases {
		var actualOutput = decompressedLength(testCase.input, true)
		if testCase.output != actualOutput {
			t.Error("Expected", testCase.output, "got", actualOutput)
		}
	}
}
