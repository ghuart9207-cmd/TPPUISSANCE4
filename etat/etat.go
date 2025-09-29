package etat

import (
	"example.com/TPPUISSANCE4/grille"
	"example.com/TPPUISSANCE4/variables"
)

func CheckWin(b *grille.Board, symbol grille.Cell) bool {
	// Check horizontal
	for r := 0; r < variables.Rows; r++ {
		for c := 0; c <= variables.Cols-4; c++ {
			if b[r][c] == symbol && b[r][c+1] == symbol && b[r][c+2] == symbol && b[r][c+3] == symbol {
				return true
			}
		}
	}
	// Check vertical
	for c := 0; c < variables.Cols; c++ {
		for r := 0; r <= variables.Rows-4; r++ {
			if b[r][c] == symbol && b[r+1][c] == symbol && b[r+2][c] == symbol && b[r+3][c] == symbol {
				return true
			}
		}
	}
	// Check diagonal (bottom-left to top-right)
	for r := 3; r < variables.Rows; r++ {
		for c := 0; c <= variables.Cols-4; c++ {
			if b[r][c] == symbol && b[r-1][c+1] == symbol && b[r-2][c+2] == symbol && b[r-3][c+3] == symbol {
				return true
			}
		}
	}
	// Check diagonal (top-left to bottom-right)
	for r := 0; r <= variables.Rows-4; r++ {
		for c := 0; c <= variables.Cols-4; c++ {
			if b[r][c] == symbol && b[r+1][c+1] == symbol && b[r+2][c+2] == symbol && b[r+3][c+3] == symbol {
				return true
			}
		}
	}
	return false
}
func IsDraw(b *grille.Board) bool {
	for c := 0; c < variables.Cols; c++ {
		if b[0][c] == grille.Empty {
			return false
		}
	}
	return true
}
func IsGameOver(b *grille.Board, symbol grille.Cell) (bool, string) {
	if CheckWin(b, symbol) {
		return true, "win"
	}
	if IsDraw(b) {
		return true, "draw"
	}
	return false, ""
}
