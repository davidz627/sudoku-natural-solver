package main

import "testing"

func TestOnlyChoice(t *testing.T) {
	testCases := []struct {
		name              string
		sudoku            string
		expectedSolutions []SolvedCell
	}{
		{
			name: "basic row",
			sudoku: `12345678.
					 ..6....28
					 5........
					 3.2.....1
					 89.2...5.
					 ...5.....
					 ....6.8.5
					 ..1...67.
					 2....7.13`,
			expectedSolutions: []SolvedCell{
				{
					Cell:  Cell{0, 8},
					value: 9,
					// TODO(p0) Reason should be an enum type
					reason: "only choice",
				},
			},
		},
		{
			name: "basic col",
			sudoku: `1234567..
					 2.6....28
					 3........
					 4.2.....1
					 .9.2...5.
					 6..5.....
					 7...6.8.5
					 8.1...67.
					 9....7.13`,
			expectedSolutions: []SolvedCell{
				{
					Cell:   Cell{4, 0},
					value:  5,
					reason: "only choice",
				},
			},
		},
		{
			name: "basic square",
			sudoku: `1234567..
					 45.....28
					 789......
					 4.2.....1
					 .9.2...5.
					 6..5.....
					 7...6.8.5
					 8.1...67.
					 9....7.13`,
			expectedSolutions: []SolvedCell{
				{
					Cell:   Cell{1, 2},
					value:  6,
					reason: "only choice",
				},
			},
		},
		{
			name: "one of each",
			sudoku: `12.456789
					 4..789123
					 789..3456
					 234567891
					 567891234
					 89123..67
					 34567..12
					 678912345
					 912345678`,
			expectedSolutions: []SolvedCell{
				{
					Cell:   Cell{0, 2},
					value:  3,
					reason: "only choice",
				},
				{
					Cell:   Cell{3, 2},
					value:  1,
					reason: "only choice",
				},
				{
					Cell:   Cell{6, 6},
					value:  9,
					reason: "only choice",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Logf("Running test case: %s", tc.name)
		sudoku := parse(tc.sudoku)
		solutions := onlyChoice(sudoku)
		if len(solutions) != len(tc.expectedSolutions) {
			t.Errorf("expected %d solutions, got %d", len(tc.expectedSolutions), len(solutions))
		}
		for i := 0; i < len(solutions); i++ {
			if solutions[i] != tc.expectedSolutions[i] {
				t.Errorf("Expected solution %v, got %v", tc.expectedSolutions[i], solutions[i])
			}
		}
	}
}
