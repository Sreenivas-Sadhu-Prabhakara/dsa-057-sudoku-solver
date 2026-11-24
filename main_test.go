package main

import "testing"

func valid(board *[9][9]int) bool {
	var rows, cols, boxes [9]uint16
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			v := board[r][c]
			if v < 1 || v > 9 {
				return false
			}
			bit := uint16(1) << (v - 1)
			b := (r/3)*3 + c/3
			if rows[r]&bit != 0 || cols[c]&bit != 0 || boxes[b]&bit != 0 {
				return false
			}
			rows[r] |= bit
			cols[c] |= bit
			boxes[b] |= bit
		}
	}
	return true
}

func TestSolve(t *testing.T) {
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
	given := board
	if !Solve(&board) {
		t.Fatal("expected solvable")
	}
	if !valid(&board) {
		t.Fatal("solution violates sudoku rules")
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if given[r][c] != 0 && board[r][c] != given[r][c] {
				t.Fatalf("clue at %d,%d changed", r, c)
			}
		}
	}

	// Contradictory clues: two 5s in a row.
	bad := [9][9]int{}
	bad[0][0], bad[0][5] = 5, 5
	if Solve(&bad) {
		t.Error("expected unsolvable for conflicting clues")
	}

	// Empty board is solvable.
	empty := [9][9]int{}
	if !Solve(&empty) || !valid(&empty) {
		t.Error("empty board should solve to a valid grid")
	}
}
