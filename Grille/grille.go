package grille

import (
	"fmt"

	"example.com/TPPUISSANCE4/variables"
)

type Cell rune

const (
	Empty Cell = '.'
	X     Cell = 'X'
	O     Cell = 'O'
)

type Board [variables.Rows][variables.Cols]Cell

func NewBoard() *Board {
	var b Board
	for r := 0; r < variables.Rows; r++ {
		for c := 0; c < variables.Cols; c++ {
			b[r][c] = Empty
		}
	}
	return &b
}

func (b *Board) Drop(col int, player Cell) error {
	if col < 0 || col >= variables.Cols {
		return fmt.Errorf("valeur non acceptée : colonne %d inexistante", col)
	}
	for r := variables.Rows - 1; r >= 0; r-- {
		if (*b)[r][col] == Empty {
			(*b)[r][col] = player
			return nil
		}
	}
	return fmt.Errorf("valeur non acceptée : colonne %d pleine", col)
}

func (b *Board) String() string {
	var s string
	for r := 0; r < variables.Rows; r++ {
		s += "| "
		for c := 0; c < variables.Cols; c++ {
			s += fmt.Sprintf("%c ", (*b)[r][c])
		}
		s += "|\n"
	}
	s += "  "
	for c := 0; c < variables.Cols; c++ {
		s += fmt.Sprintf("%d ", c)
	}
	return s
}
