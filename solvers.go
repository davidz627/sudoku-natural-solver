package main

func onlyChoice(sudoku Sudoku) []SolvedCell {
	// TODO
	// Check each row
	solutions := make([]SolvedCell, 0)
	for row := 0; row < 9; row++ {
		zeros := make([]Cell, 0)
		rowSum := 0
		for col := 0; col < 9; col++ {
			if sudoku.matrix[row][col] == 0 {
				zeros = append(zeros, Cell{row, col})
			}
			rowSum += sudoku.matrix[row][col]
		}
		if len(zeros) == 1 {
			solutions = append(solutions, SolvedCell{
				Cell: Cell{
					row: zeros[0].row,
					col: zeros[0].col,
				},
				value:  45 - rowSum,
				reason: "only choice",
			})
		}
	}
	for col := 0; col < 9; col++ {
		zeros := make([]Cell, 0)
		colSum := 0
		for row := 0; row < 9; row++ {
			if sudoku.matrix[row][col] == 0 {
				zeros = append(zeros, Cell{row, col})
			}
			colSum += sudoku.matrix[row][col]
		}
		if len(zeros) == 1 {
			solutions = append(solutions, SolvedCell{
				Cell: Cell{
					row: zeros[0].row,
					col: zeros[0].col,
				},
				value:  45 - colSum,
				reason: "only choice",
			})
		}
	}
	return solutions
}
