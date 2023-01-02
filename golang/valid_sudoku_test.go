package leetcode

import "testing"

func TestIsValidSudoku(t *testing.T) {
	var tests = []struct {
		board [][]byte
		want  bool
	}{
		{[][]byte{
			[]byte("53..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		}, true},
		{[][]byte{
			[]byte("83..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		}, false},
	}

	for _, test := range tests {
		if got := isValidSudoku(test.board); got != test.want {
			t.Errorf("isValidSudoku(%v) = %v", test.board, got)
		}
	}
}
