package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "LVIII",
			output: 58,
		},
		{
			input:  "MCMXCIV",
			output: 1994,
		},
	}

	for _, testCase := range testCases {
		result := romanToInt(testCase.input)
		if result != testCase.output {
			t.Errorf("input: %s, expected: %d, got: %d", testCase.input, testCase.output, result)
		}
	}
}
