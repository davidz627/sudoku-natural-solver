package main

import (
	"fmt"
	"strings"
)

/* Simplest Sudoku
`123456789
456789123
789123456
234567891
567891234
891234567
345678912
678912345
912345678`
*/

// https://qqwing.com/generate.html
const testSudoku = `123456789
					456789123
					789123456
					234567891
					567891234
					891234567
					345678912
					678912345
					912345678`

func main() {
	fmt.Println("Hello!")
	sudoku := parse(testSudoku)
	solutions := onlyChoice(sudoku)
	fmt.Println(solutions)
	// Solve based on human solving methods
	// https://www.sudokudragon.com/sudokustrategy.htm
	// Sudoku generator: https://qqwing.com/generate.html
	// TODO(p0)
	// Only Choice

}

type Cell struct {
	row int
	col int
}

type SolvedCell struct {
	Cell
	value  uint8
	reason string
}

type Sudoku struct {
	matrix [][]uint8
}

func sum(arr []uint8) uint8 {
	sum := uint8(0)
	for _, v := range arr {
		sum += v
	}
	return sum
}

func getRow(sudoku Sudoku, row int) []uint8 {
	return sudoku.matrix[row]
}

func getCol(sudoku Sudoku, col int) []uint8 {
	colArr := make([]uint8, 9)
	for i := 0; i < 9; i++ {
		colArr[i] = sudoku.matrix[i][col]
	}
	return colArr
}

func getBox(sudoku Sudoku, row int, col int) []uint8 {
	// TODO(p0) this needs to be tested...
	boxArr := make([]uint8, 9)
	// determine what box the row/col is in
	rowBox := row / 3
	colBox := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boxArr[i*3+j] = sudoku.matrix[rowBox*3+i][colBox*3+j]
		}
	}
	return boxArr
}

func printSudoku(sudoku Sudoku) {
	for _, row := range sudoku.matrix {
		fmt.Println(row)
	}
}

func parse(sudoku string) Sudoku {
	// parse a sudoku string into a 9x9 matrix where . is not specified
	lines := strings.Split(sudoku, "\n")
	if len(lines) != 9 {
		panic("wrong number of rows")
	}
	matrix := make([][]uint8, 9)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) != 9 {
			panic("wrong number of columns")
		}
		matrix[i] = make([]uint8, 9)
		for j, c := range line {
			if c == '.' {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = uint8(c - '0')
			}
		}
	}
	return Sudoku{matrix}

}
