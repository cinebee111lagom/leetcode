package backtracking

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name   string
		board  [][]byte
		word   string
		exists bool
	}{
		{
			"basic case",
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{
			"not found",
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
		{
			"single char found",
			[][]byte{{'A'}},
			"A",
			true,
		},
		{
			"single char not found",
			[][]byte{{'A'}},
			"B",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.board, tt.word); got != tt.exists {
				t.Errorf("Exist() = %v, want %v", got, tt.exists)
			}
		})
	}
}