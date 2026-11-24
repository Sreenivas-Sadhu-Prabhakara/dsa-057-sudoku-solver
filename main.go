package main

import "fmt"

// Solve fills a 9×9 Sudoku board in place (0 = empty) and reports whether a
// solution exists. Row, column and box usage are tracked as bitmasks so
// candidate checks are single AND operations.
func Solve(board *[9][9]int) bool {
	var rows, cols, boxes [9]uint16
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if v := board[r][c]; v != 0 {
				bit := uint16(1) << (v - 1)
				b := (r/3)*3 + c/3
				if rows[r]&bit != 0 || cols[c]&bit != 0 || boxes[b]&bit != 0 {
					return false // the given clues already conflict
				}
				rows[r] |= bit
				cols[c] |= bit
				boxes[b] |= bit
			}
		}
	}
	var fill func(idx int) bool
	fill = func(idx int) bool {
		for idx < 81 && board[idx/9][idx%9] != 0 {
			idx++
		}
		if idx == 81 {
			return true
		}
		r, c := idx/9, idx%9
		b := (r/3)*3 + c/3
		for v := 1; v <= 9; v++ {
			bit := uint16(1) << (v - 1)
			if rows[r]&bit != 0 || cols[c]&bit != 0 || boxes[b]&bit != 0 {
				continue
			}
			board[r][c] = v
			rows[r] |= bit
			cols[c] |= bit
			boxes[b] |= bit
			if fill(idx + 1) {
				return true
			}
			board[r][c] = 0
			rows[r] &^= bit
			cols[c] &^= bit
			boxes[b] &^= bit
		}
		return false
	}
	return fill(0)
}

func main() {
	board := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	Solve(&board)
	for _, row := range board {
		fmt.Println(row)
	}
}
